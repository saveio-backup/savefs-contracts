//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";
import "./ProveExtra.sol";

contract Prove is Initializable, IProve, IFsEvent {
    IConfig config;
    IFile fs;
    INode node;
    IPDP pdp;
    ISector sector;
    ProveExtra proveExtra;

    uint64 SECTOR_PROVE_BLOCK_NUM;

    mapping(address => mapping(uint64 => uint256)) punishmentHeightForNode;

    function initialize(
        IConfig _config,
        IFile _fs,
        INode _node,
        IPDP _pdp,
        ISector _sector,
        ProveConfig memory proveConfig,
        ProveExtra _proveExtra
    ) public initializer {
        config = _config;
        fs = _fs;
        node = _node;
        pdp = _pdp;
        sector = _sector;
        SECTOR_PROVE_BLOCK_NUM = proveConfig.SECTOR_PROVE_BLOCK_NUM;
        proveExtra = _proveExtra;
    }

    function FileProve(FileProveParams memory fileProve)
        public
        payable
        virtual
        override
    {
        FileInfo memory fileInfo = fs.GetFileInfo(fileProve.FileHash);
        if (fileInfo.IsPlotFile) {
            if (fileProve.NodeWallet != fileInfo.FileOwner) {
                emit FsError("FileProve", "FileProveNotFileOwner");
                return;
            }
        } else {
            bool canProve = false;
            address[] memory primaryNodes = fileInfo.PrimaryNodes;
            for (uint256 i = 0; i < primaryNodes.length; i++) {
                if (primaryNodes[i] == fileProve.NodeWallet) {
                    canProve = true;
                    break;
                }
            }
            if (!canProve) {
                address[] memory candidateNodes = fileInfo.CandidateNodes;
                for (uint256 i = 0; i < candidateNodes.length; i++) {
                    if (candidateNodes[i] == fileProve.NodeWallet) {
                        canProve = true;
                        break;
                    }
                }
            }
            if (!canProve) {
                emit FsError("FileProve", "FileProveNotNode");
                return;
            }
        }
        NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
            fileProve.NodeWallet
        );
        ProveDetail[] memory details = proveExtra
            .getProveDetailListWithNodeAddr(node, fileInfo.FileHash);
        if (fileProve.SectorID != 0 && block.number < fileInfo.ExpiredHeight) {
            for (uint256 i = 0; i < details.length; i++) {
                if (details[i].WalletAddr == fileProve.NodeWallet) {
                    emit FsError("FileProve", "FileProveAlreadyProved");
                    return;
                }
            }
        }
        bool success = proveExtra.checkProve(pdp, fileProve, fileInfo);
        if (!success) {
            emit FsError("FileProve", "FileProveCheckFailed");
            return;
        }
        bool found = false;
        bool settleFlag = false;
        uint64 haveProveTimes = 0;
        ProveDetail memory detail;
        uint256 fileExpiredHeight = fileInfo.ExpiredHeight;
        for (uint256 i = 0; i < details.length; i++) {
            if (details[i].WalletAddr == fileProve.NodeWallet) {
                haveProveTimes = detail.ProveTimes;
                if (
                    haveProveTimes == fileInfo.ProveTimes ||
                    block.number > fileExpiredHeight
                ) {
                    detail.Finished = true;
                    settleFlag = true;
                }
                if (haveProveTimes > fileInfo.ProveTimes) {
                    emit FsError("FileProve", "FileProveTimesExceed");
                    return;
                }
                bool r = proveExtra.checkProveExpire(fileInfo.ExpiredHeight);
                if (r) {
                    emit FsError("FileProve", "FileProveExpired");
                    return;
                }
                detail.ProveTimes++;
                found = true;
                break;
            }
        }
        if (!found) {
            if (details.length == fileInfo.CopyNum + 1) {
                emit FsError("FileProve", "FileProveCopyNumExceed");
                return;
            }
            if (
                nodeInfo.RestVol <
                fileInfo.FileBlockNum * fileInfo.FileBlockSize
            ) {
                emit FsError("FileProve", "FileProveNodeVolNotEnough");
                return;
            }
            nodeInfo.RestVol -= fileInfo.FileBlockNum * fileInfo.FileBlockSize;
            node.UpdateNodeInfo(nodeInfo);
            detail.NodeAddr = nodeInfo.NodeAddr;
            detail.WalletAddr = fileProve.NodeWallet;
            detail.ProveTimes = 1;
            detail.BlockHeight = block.number;
            detail.Finished = false;
            ProveDetail[] memory detailsTmp = new ProveDetail[](
                details.length + 1
            );
            for (uint256 i = 0; i < details.length; i++) {
                detailsTmp[i] = details[i];
            }
            detailsTmp[detailsTmp.length - 1] = detail;
            details = detailsTmp;
        }
        UpdateProveDetailList(fileInfo.FileHash, details);
        uint64 sectorID = fileProve.SectorID;

        ISector _sector = sector;
        FileInfo memory _fileInfo = fileInfo;
        uint256 proveHeight = fileProve.BlockHeight;
        SectorRef memory sectorRef = SectorRef({
            NodeAddr: nodeInfo.WalletAddr,
            SectorId: sectorID
        });
        if (!found) {
            SectorInfo memory sectorInfo = _sector.GetSectorInfo(sectorRef);
            if (sectorInfo.IsPlots != _fileInfo.IsPlotFile) {
                emit FsError("FileProve", "FileProveSectorTypeNotMatch");
                return;
            }
            _sector.AddFileToSector(sectorInfo, _fileInfo);
            string memory err = proveExtra.AddSectorRefForFileInfo(
                _sector,
                fs,
                _fileInfo,
                sectorRef
            );
            if (bytes(err).length > 0) {
                emit FsError("FileProve", err);
                return;
            }
            if (sectorInfo.NextProveHeight == 0) {
                sectorInfo.NextProveHeight =
                    proveHeight +
                    _fileInfo.ProveInterval;
            }
            _sector.UpdateSectorInfo(sectorInfo);
        }
        NodeInfo memory _nodeInfo = nodeInfo;
        ProveDetail memory _detail = detail;
        ProveDetail[] memory _details = details;
        Setting memory _setting = config.GetSetting();
        if (settleFlag) {
            if (sectorID != 0) {
                SectorInfo memory sectorInfo = sector.GetSectorInfo(sectorRef);
                sector.DeleteFileFromSector(sectorInfo, _fileInfo);
            }
            string memory err = proveExtra.SettleForFile(
                node,
                fs,
                _fileInfo,
                _nodeInfo,
                _detail,
                _details,
                _setting
            );
            if (
                keccak256(abi.encodePacked(err)) !=
                keccak256(abi.encodePacked(""))
            ) {
                emit FsError("FileProve", err);
                return;
            }
        }
        emit FilePDPSuccessEvent(
            FsEvent.FILE_PDP_SUCCESS,
            block.number,
            _fileInfo.FileHash,
            _nodeInfo.WalletAddr
        );
    }

    function SectorProve(SectorProveParams memory sectorProve)
        public
        payable
        virtual
        override
    {
        NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
            sectorProve.NodeAddr
        );
        SectorInfo memory sectorInfo = sector.GetSectorInfo(
            SectorRef({
                SectorId: sectorProve.SectorID,
                NodeAddr: sectorProve.NodeAddr
            })
        );
        Setting memory setting = config.GetSetting();
        if (block.number < sectorInfo.NextProveHeight) {
            revert SectorProveFailed(1);
        }
        if (sectorProve.ChallengeHeight != sectorInfo.NextProveHeight) {
            revert SectorProveFailed(2);
        }
        bool r = checkSectorProve(sectorProve, sectorInfo);
        if (!r) {
            punishForSector(sectorInfo, nodeInfo, setting, 1);
            revert SectorProveFailed(3);
        }
        bool p = profitSplitForSector(sectorInfo, nodeInfo, setting);
        if (!p) {
            revert SectorProveFailed(5);
        }
        if (sectorInfo.FirstProveHeight == 0) {
            sectorInfo.FirstProveHeight = block.number;
        }
        sectorInfo.NextProveHeight = block.number + setting.DefaultProvePeriod;
        sector.UpdateSectorInfo(sectorInfo);
        if (!sectorInfo.IsPlots) {
            revert SectorProveFailed(4);
        }
        // poc prove
        PocProve memory _pocProve = proveExtra.getPocProve(
            sectorInfo.NodeAddr,
            block.number
        );
        _pocProve.Height = block.number;
        _pocProve.Miner = sectorInfo.NodeAddr;
        _pocProve.PlotSize = sectorInfo.Used;
        proveExtra.putPocProve(_pocProve);
    }

    function GetProveDetailList(bytes memory fileHash)
        public
        view
        virtual
        override
        returns (ProveDetail[] memory)
    {
        return proveExtra.GetProveDetailList(fileHash);
    }

    function UpdateProveDetailList(
        bytes memory fileHash,
        ProveDetail[] memory details
    ) public payable {
        proveExtra.UpdateProveDetailList(fileHash, details);
    }

    function DeleteProveDetails(bytes memory fileHash)
        public
        payable
        virtual
        override
    {
        proveExtra.DeleteProveDetails(fileHash);
    }

    function UpdateProveDetailMeta(
        bytes memory fileHash,
        ProveDetailMeta memory details
    ) public payable virtual override {
        proveExtra.UpdateProveDetailMeta(fileHash, details);
    }

    function punishForSector(
        SectorInfo memory sectorInfo,
        NodeInfo memory nodeInfo,
        Setting memory setting,
        uint64 times
    ) public payable {
        uint64 amount = times *
            calPunishmentForOneSectorProve(setting, sectorInfo);
        if (nodeInfo.Pledge >= amount) {
            nodeInfo.Pledge -= amount;
        } else {
            nodeInfo.Pledge = 0;
            amount = nodeInfo.Pledge;
        }
        if (amount > 0) {
            require(msg.value >= amount, "PunishForSector failed");
            node.UpdateNodeInfo(nodeInfo);
        }
        SetLastPunishmentHeightForNode(
            sectorInfo.NodeAddr,
            sectorInfo.SectorID,
            block.number
        );
    }

    function CheckNodeSectorProvedInTime(SectorRef memory sectorRef)
        public
        payable
        virtual
        override
    {
        address nodeAddr = sectorRef.NodeAddr;
        uint64 sectorID = sectorRef.SectorId;
        NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(nodeAddr);
        if (nodeInfo.ServiceTime < block.timestamp) {
            revert NodeSectorProvedInTimeError();
        }
        if (sectorID == 0) {
            revert NodeSectorProvedInTimeError();
        }
        SectorInfo memory sectorInfo = sector.GetSectorInfo(sectorRef);
        if (sectorInfo.FileNum == 0) {
            revert NodeSectorProvedInTimeError();
        }
        Setting memory setting = config.GetSettingWithProveLevel(
            sectorInfo.ProveLevel_
        );
        uint256 height = block.number;
        if (sectorInfo.NextProveHeight + setting.DefaultProvePeriod < height) {
            revert NodeSectorProvedInTimeError();
        }
        uint256 lastHeight = GetLastPunishmentHeightForNode(nodeAddr, sectorID);
        uint64 times = calMissingSectorProveTimes(
            sectorInfo,
            setting,
            lastHeight,
            height
        );
        if (times == 0) {
            revert NodeSectorProvedInTimeError();
        }
        punishForSector(sectorInfo, nodeInfo, setting, times);
    }

    function SetLastPunishmentHeightForNode(
        address nodeAddr,
        uint64 sectorId,
        uint256 height
    ) public payable {
        punishmentHeightForNode[nodeAddr][sectorId] = height;
    }

    function getProveDetailListWithNodeAddr(bytes memory fileHash)
        private
        view
        returns (ProveDetail[] memory)
    {
        return proveExtra.getProveDetailListWithNodeAddr(node, fileHash);
    }

    function checkSectorProve(
        SectorProveParams memory sectorProve,
        SectorInfo memory sectorInfo
    ) public view returns (bool) {
        // TODO block head hash
        bytes memory blockHash;
        GenChallengeParams memory gParams;
        gParams.WalletAddr = sectorProve.NodeAddr;
        gParams.HashValue = blockHash;
        gParams.FileBlockNum = sectorInfo.TotalBlockNum;
        gParams.ProveNum = SECTOR_PROVE_BLOCK_NUM;
        Challenge[] memory challenges = pdp.GenChallenge(gParams);
        // pre
        // TODO decentralized sector prove data
        SectorProveData memory sectorProveData;
        PrepareForPdpVerificationParams memory pParams;
        pParams.SectorInfo_ = sectorInfo;
        // pParams.Challenges = challenges;
        pParams.ProveData = sectorProveData;
        PdpVerificationReturns memory pReturns;
        pReturns = pdp.PrepareForPdpVerification(pParams);
        if (!pReturns.Success) {
            return false;
        }
        // verify
        VerifyProofWithMerklePathForFileParams memory vParams;
        vParams.Version = 0;
        vParams.Proofs = sectorProveData.Proofs;
        vParams.FileIds = pReturns.FileIDs;
        vParams.Tags = pReturns.Tags;
        vParams.Challenges = pReturns.UpdatedChal;
        vParams.MerklePath_ = pReturns.Path;
        vParams.RootHashes = pReturns.RootHashes;
        bool res = pdp.VerifyProofWithMerklePathForFile(vParams);
        if (!res) {
            return false;
        }
        if (sectorInfo.IsPlots) {
            if (
                !pReturns.FileInfo_.IsPlotFile ||
                pReturns.FileInfo_.PlotInfo_.Nonces == 0
            ) {
                // TODO
                // return false;
            }
            VerifyPlotDataParams memory vpParams;
            vpParams.PlotInfo_ = pReturns.FileInfo_.PlotInfo_;
            vpParams.PlotData = sectorProveData.PlotData;
            if (challenges.length > 0) {
                vpParams.Index = uint64(challenges[0].Index);
            }
            bool res2 = pdp.VerifyPlotData(vpParams);
            if (!res2) {
                return false;
            }
        }
        return true;
    }

    function calcSingleValidFeeForFile(Setting memory setting, uint64 fileSize)
        private
        pure
        returns (uint64)
    {
        return uint64(setting.GasForChallenge * fileSize) / 1024000;
    }

    function calPunishmentForOneSectorProve(
        Setting memory setting,
        SectorInfo memory sectorInfo
    ) private pure returns (uint64) {
        uint64 punishFactor = 2;
        uint64 res = punishFactor *
            calcSingleValidFeeForFile(setting, sectorInfo.Used);
        return res;
    }

    function profitSplitForSector(
        SectorInfo memory sectorInfo,
        NodeInfo memory nodeInfo,
        Setting memory setting
    ) public payable returns (bool) {
        for (uint256 i = 0; i < sectorInfo.FileNum; i++) {
            bytes memory fileHash = sectorInfo.FileList[i];
            FileInfo memory fileInfo = fs.GetFileInfo(fileHash);
            ProveDetail[] memory details = GetProveDetailList(
                fileInfo.FileHash
            );
            bool settleFlag = false;
            uint256 fileExpiredHeight = fileInfo.ExpiredHeight;
            bool found = false;
            ProveDetail memory detail;
            for (uint256 j = 0; j < details.length; j++) {
                if (details[i].WalletAddr == sectorInfo.NodeAddr) {
                    found = true;
                    detail = details[i];
                    uint64 haveProveTimes = detail.ProveTimes;
                    if (
                        haveProveTimes == fileInfo.ProveTimes ||
                        block.number > fileExpiredHeight
                    ) {
                        detail.Finished = true;
                        settleFlag = true;
                    }
                    detail.ProveTimes++;
                    break;
                }
            }
            if (!found) {
                return false;
            }
            SectorInfo memory _sectorInfo = sectorInfo;
            FileInfo memory _fileInfo = fileInfo;
            NodeInfo memory _nodeInfo = nodeInfo;
            Setting memory _setting = setting;
            UpdateProveDetailList(_fileInfo.FileHash, details);
            if (settleFlag) {
                proveExtra.SettleForFile(
                    node,
                    fs,
                    _fileInfo,
                    _nodeInfo,
                    detail,
                    details,
                    _setting
                );
                sector.DeleteFileFromSector(_sectorInfo, _fileInfo);
                emit DeleteFileEvent(
                    FsEvent.DELETE_FILE,
                    block.number,
                    _fileInfo.FileHash,
                    _nodeInfo.WalletAddr
                );
            }
        }
        return true;
    }

    function calMissingSectorProveTimes(
        SectorInfo memory sectorInfo,
        Setting memory setting,
        uint256 lastPunishHeight,
        uint256 currHeight
    ) private pure returns (uint64) {
        uint64 interval = setting.DefaultProvePeriod;
        uint256 nextProveHeight = sectorInfo.NextProveHeight;
        if (nextProveHeight + interval >= currHeight) {
            return 0;
        }
        uint64 totalTimes = uint64(currHeight - nextProveHeight) / interval;
        uint64 punishedTimes;
        if (lastPunishHeight != 0) {
            if (lastPunishHeight > nextProveHeight + interval) {
                punishedTimes =
                    uint64(lastPunishHeight - nextProveHeight) /
                    interval;
            } else {
                punishedTimes = 0;
            }
        }
        if (totalTimes < punishedTimes) {
            return 0;
        }
        return totalTimes - punishedTimes;
    }

    function GetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorID)
        private
        view
        returns (uint256)
    {
        return punishmentHeightForNode[nodeAddr][sectorID];
    }
}
