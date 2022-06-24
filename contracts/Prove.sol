//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./Config.sol";
import "./FileSystem.sol";
import "./Node.sol";
import "./PDP.sol";
import "./Sector.sol";
import "./interface.sol";

contract Prove is Initializable, IProve {
    Config config;
    FileSystem fs;
    Node node;
    PDP pdp;
    Sector sector;

    uint64 SECTOR_PROVE_BLOCK_NUM;
    using IterableMapping for ItMap;

    mapping(bytes => ItMap) proveDetails; // fileHash => nodeAddr => ProveDetail
    mapping(bytes => ProveDetailMeta) proveDetailMeta; // fileHash => ProveDetailMeta
    mapping(address => mapping(uint64 => uint256)) punishmentHeightForNode;
    mapping(string => PocProve) pocProve; // miner + height => PocProve

    function initialize(
        Config _config,
        FileSystem _fs,
        Node _node,
        PDP _pdp,
        Sector _sector,
        ProveConfig memory proveConfig
    ) public initializer {
        config = _config;
        fs = _fs;
        node = _node;
        pdp = _pdp;
        sector = _sector;
        SECTOR_PROVE_BLOCK_NUM = proveConfig.SECTOR_PROVE_BLOCK_NUM;
    }

    function FileProve(FileProveParams memory fileProve)
        public
        virtual
        override
    {
        Setting memory setting = config.GetSetting();
        FileInfo memory fileInfo = fs.GetFileInfo(fileProve.FileHash);
        if (fileInfo.IsPlotFile) {
            if (fileProve.NodeWallet != fileInfo.FileOwner) {
                revert FileProveNotFileOwner();
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
                revert FileProveFailed(1);
            }
        }
        NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
            fileProve.NodeWallet
        );
        ProveDetail[] memory details = getProveDetailListWithNodeAddr(
            fileInfo.FileHash
        );
        if (fileProve.SectorID != 0 && block.number < fileInfo.ExpiredHeight) {
            for (uint256 i = 0; i < details.length; i++) {
                if (details[i].WalletAddr == fileProve.NodeWallet) {
                    revert FileProveFailed(2);
                }
            }
        }
        bool success = checkProve(fileProve, fileInfo);
        if (!success) {
            revert FileProveFailed(3);
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
                    revert FileProveFailed(4);
                }
                bool r = checkProveExpire(fileInfo.ExpiredHeight);
                if (r) {
                    revert FileProveFailed(5);
                }
                detail.ProveTimes++;
                found = true;
                break;
            }
        }
        if (!found) {
            if (details.length == fileInfo.CopyNum + 1) {
                revert FileProveFailed(6);
            }
            if (
                nodeInfo.RestVol <
                fileInfo.FileBlockNum * fileInfo.FileBlockSize
            ) {
                revert FileProveFailed(7);
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
        if (!found) {
            SectorInfo memory sectorInfo = sector.GetSectorInfo(
                SectorRef({
                    NodeAddr: nodeInfo.NodeAddr,
                    SectorId: fileProve.SectorID
                })
            );
            if (sectorInfo.IsPlots != fileInfo.IsPlotFile) {
                revert FileProveFailed(8);
            }
            sector.AddFileToSector(sectorInfo, fileInfo);
            sector.AddSectorRefForFileInfo(sectorInfo);
            if (sectorInfo.NextProveHeight == 0) {
                sectorInfo.NextProveHeight =
                    fileProve.BlockHeight +
                    fileInfo.ProveInterval;
            }
            sector.UpdateSectorInfo(sectorInfo);
        }
        if (settleFlag) {
            if (fileProve.SectorID != 0) {
                SectorInfo memory sectorInfo = sector.GetSectorInfo(
                    SectorRef({
                        NodeAddr: nodeInfo.NodeAddr,
                        SectorId: fileProve.SectorID
                    })
                );
                sector.DeleteFileFromSector(sectorInfo, fileInfo);
            }
            SettleForFile(fileInfo, nodeInfo, detail, details, setting);
        }
        emit FilePDPSuccessEvent(
            FsEvent.FILE_PDP_SUCCESS,
            block.number,
            fileInfo.FileHash,
            nodeInfo.WalletAddr
        );
    }

    function SectorProve(SectorProveParams memory sectorProve)
        public
        virtual
        override
    {
        NodeInfo memory nodeInfo = node.GetNodeInfoByNodeAddr(
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
            console.log(
                "ChallengeHeight error:",
                sectorProve.ChallengeHeight,
                sectorInfo.NextProveHeight
            );
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
        PocProve memory _pocProve = getPocProve(
            sectorInfo.NodeAddr,
            block.number
        );
        _pocProve.Height = block.number;
        _pocProve.Miner = sectorInfo.NodeAddr;
        _pocProve.PlotSize = sectorInfo.Used;
        putPocProve(_pocProve);
    }

    function GetProveDetailList(bytes memory fileHash)
        public
        view
        virtual
        override
        returns (ProveDetail[] memory)
    {
        ItMap storage data = proveDetails[fileHash];
        ProveDetail[] memory result = new ProveDetail[](data.size);
        if (data.size == 0) {
            return result;
        }
        for (
            uint256 i = data.iterate_start();
            data.iterate_valid(i);
            i = data.iterate_next(i)
        ) {
            (, ProveDetail memory value) = data.iterate_get(i);
            result[i] = value;
        }
        return result;
    }

    function UpdateProveDetailMeta(
        bytes memory fileHash,
        ProveDetailMeta memory details
    ) public {
        proveDetailMeta[fileHash] = details;
    }

    function UpdateProveDetailList(
        bytes memory fileHash,
        ProveDetail[] memory details
    ) public {
        ItMap storage data = proveDetails[fileHash];
        for (uint256 i = 0; i < details.length; i++) {
            ProveDetail memory detail = details[i];
            data.insert(detail.NodeAddr, detail);
        }
    }

    function DeleteProveDetails(bytes memory fileHash) public {
        delete proveDetails[fileHash];
        delete proveDetailMeta[fileHash];
    }

    function SettleForFile(
        FileInfo memory fileInfo,
        NodeInfo memory nodeInfo,
        ProveDetail memory detail,
        ProveDetail[] memory details,
        Setting memory setting
    ) public {
        uint64 profit = calculateProfitForSettle(fileInfo, detail, setting);
        if (fileInfo.Deposit < profit) {
            revert FileProveFailed(9);
        }

        nodeInfo.RestVol += profit;
        node.UpdateNodeInfo(nodeInfo);

        fileInfo.Deposit -= profit;
        fileInfo.ValidFlag = false;
        fs.UpdateFileInfo(fileInfo);

        uint64 finishedNodes = 0;
        for (uint256 i = 0; i < details.length; i++) {
            if (details[i].Finished) {
                finishedNodes++;
            }
        }
        if (finishedNodes == 1) {
            fs.CleanupForDeleteFile(fileInfo, false, true);
        }
        if (finishedNodes == fileInfo.CopyNum + 1) {
            if (fileInfo.Deposit > 0) {
                payable(fileInfo.FileOwner).transfer(fileInfo.Deposit);
            }
            fs.CleanupForDeleteFile(fileInfo, true, false);
        } else {
            fs.AddFileToUnSettleList(fileInfo.FileOwner, fileInfo.FileHash);
        }
        emit ProveFileEvent(
            FsEvent.PROVE_FILE,
            block.number,
            nodeInfo.WalletAddr,
            profit
        );
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

    function getPocProve(address nodeAddr, uint256 height)
        public
        view
        returns (PocProve memory)
    {
        string memory key = string(abi.encodePacked(nodeAddr, height));
        return pocProve[key];
    }

    function putPocProve(PocProve memory prove) public {
        string memory key = string(abi.encodePacked(prove.Miner, prove.Height));
        pocProve[key] = prove;
    }

    function GetPocProveList() public view returns (PocProve[] memory) {
        // TODO
    }

    function CheckNodeSectorProvedInTime(SectorRef memory sectorRef)
        public
        virtual
        override
    {
        address nodeAddr = sectorRef.NodeAddr;
        uint64 sectorID = sectorRef.SectorId;
        NodeInfo memory nodeInfo = node.GetNodeInfoByNodeAddr(nodeAddr);
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
    ) public {
        punishmentHeightForNode[nodeAddr][sectorId] = height;
    }

    function getProveDetailListWithNodeAddr(bytes memory fileHash)
        private
        view
        returns (ProveDetail[] memory)
    {
        ProveDetail[] memory details = GetProveDetailList(fileHash);
        for (uint256 i = 0; i < details.length; i++) {
            NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
                details[i].WalletAddr
            );
            details[i].NodeAddr = nodeInfo.NodeAddr;
        }
        return details;
    }

    function checkProve(
        FileProveParams memory fileProve,
        FileInfo memory fileInfo
    ) private view returns (bool) {
        uint256 currBlockHeight = block.number;
        if (
            fileProve.BlockHeight > currBlockHeight + fileInfo.ProveInterval ||
            fileProve.BlockHeight + fileInfo.ProveInterval < currBlockHeight
        ) {
            return false;
        }
        // TODO deserialize = fileInfo.FileProveParam
        ProveParam memory proveParam;
        // TODO params to deserialize  = fileProve.ProveData
        ProveData memory proveData;
        // challenge
        // TODO block head hash
        bytes memory blockHash;
        GenChallengeParams memory gParams;
        gParams.WalletAddr = fileProve.NodeWallet;
        gParams.HashValue = blockHash;
        gParams.FileBlockNum = fileInfo.FileBlockNum;
        gParams.ProveNum = fileInfo.ProveBlockNum;
        Challenge[] memory challenges = pdp.GenChallenge(gParams);
        // verify
        VerifyProofWithMerklePathForFileParams memory vParams;
        vParams.Version = 0;
        vParams.Proofs = proveData.Proofs;
        vParams.FileIds = proveParam.FileID;
        vParams.Tags = proveData.Tags;
        vParams.Challenges = challenges;
        vParams.MerklePath_ = proveData.MerklePath_;
        vParams.RootHashes = proveParam.RootHash;
        bool res = pdp.VerifyProofWithMerklePathForFile(vParams);
        if (!res) {
            return false;
        }
        return true;
    }

    function checkProveExpire(uint256 fileExpiredHeight)
        private
        view
        returns (bool)
    {
        if (block.number > fileExpiredHeight) {
            return true;
        }
        return false;
    }

    function calculateProfitForSettle(
        FileInfo memory fileInfo,
        ProveDetail memory detail,
        Setting memory setting
    ) private view returns (uint64) {
        StorageFee memory total = fs.CalcFee(
            setting,
            detail.ProveTimes - 1,
            0,
            fileInfo.FileBlockNum * fileInfo.FileBlockSize,
            uint64(fileInfo.ExpiredHeight - fileInfo.BlockHeight)
        );
        return total.TxnFee + total.SpaceFee + total.ValidationFee;
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
    ) private returns (bool) {
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
            UpdateProveDetailList(fileInfo.FileHash, details);
            if (settleFlag) {
                SettleForFile(fileInfo, nodeInfo, detail, details, setting);
                sector.DeleteFileFromSector(sectorInfo, fileInfo);
                emit DeleteFileEvent(
                    FsEvent.DELETE_FILE,
                    block.number,
                    fileInfo.FileHash,
                    nodeInfo.WalletAddr
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

struct IndexValue {
    uint256 keyIndex;
    ProveDetail value;
}
struct KeyFlag {
    address key;
    bool deleted;
}

struct ItMap {
    mapping(address => IndexValue) data;
    KeyFlag[] keys;
    uint256 size;
}

library IterableMapping {
    function insert(
        ItMap storage self,
        address key,
        ProveDetail memory value
    ) internal returns (bool replaced) {
        uint256 keyIndex = self.data[key].keyIndex;
        self.data[key].value = value;
        if (keyIndex > 0) return true;
        else {
            keyIndex = self.keys.length;
            self.keys.push();
            self.data[key].keyIndex = keyIndex + 1;
            self.keys[keyIndex].key = key;
            self.size++;
            return false;
        }
    }

    function remove(ItMap storage self, address key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(ItMap storage self, address key)
        internal
        view
        returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(ItMap storage self)
        internal
        view
        returns (uint256 keyIndex)
    {
        uint256 index = iterate_next(self, type(uint256).min);
        return index - 1;
    }

    function iterate_valid(ItMap storage self, uint256 keyIndex)
        internal
        view
        returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(ItMap storage self, uint256 keyIndex)
        internal
        view
        returns (uint256 r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(ItMap storage self, uint256 keyIndex)
        internal
        view
        returns (address key, ProveDetail memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
