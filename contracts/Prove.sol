//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./Config.sol";
import "./FileSystem.sol";
import "./Node.sol";
import "./PDP.sol";
import "./Sector.sol";

contract Prove is Initializable {
    struct FileProveParams {
        bytes FileHash;
        bytes ProveData;
        uint256 BlockHeight;
        address NodeWallet;
        uint64 Profit;
        uint64 SectorID;
    }

    Config config;
    FileSystem fs;
    Node node;
    PDP pdp;
    Sector sector;

    mapping(bytes => ProveDetail[]) proveDetail; // fileHash => ProveDetail[]
    mapping(bytes => ProveDetails) proveDetails; // fileHash => ProveDetails

    event FilePDPSuccessEvent(
        FsEvent eventType,
        uint256 blockHeight,
        bytes fileHash,
        address walletAddr
    );

    error FileProveNotFileOwner();
    error FileProveFailed(uint64);
    error SectorProveFailed();
    error NodeSectorProvedInTimeError();

    function initialize(
        Config _config,
        FileSystem _fs,
        Node _node,
        PDP _pdp,
        Sector _sector
    ) public initializer {
        config = _config;
        fs = _fs;
        node = _node;
        pdp = _pdp;
        sector = _sector;
    }

    function SetProveDetails(bytes memory fileHash, ProveDetails memory details)
        public
    {
        proveDetails[fileHash] = details;
    }

    function GetProveDetailList(bytes memory fileHash)
        public
        view
        returns (ProveDetail[] memory)
    {
        return proveDetail[fileHash];
    }

    function getProveDetailsWithNodeAddr(bytes memory fileHash)
        public
        view
        returns (ProveDetail[] memory)
    {
        ProveDetail[] memory details = proveDetail[fileHash];
        for (uint256 i = 0; i < details.length; i++) {
            NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
                details[i].WalletAddr
            );
            details[i].NodeAddr = nodeInfo.NodeAddr;
        }
        return details;
    }

    function GetFileProveDetails(bytes memory fileHash)
        public
        view
        returns (ProveDetail[] memory)
    {
        ProveDetail[] memory details = proveDetail[fileHash];
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
    ) public view returns (bool) {
        uint256 currBlockHeight = block.number;
        if (
            fileProve.BlockHeight > currBlockHeight + fileInfo.ProveInterval ||
            fileProve.BlockHeight + fileInfo.ProveInterval < currBlockHeight
        ) {
            return false;
        }
        // TODO complete pdp prove
        uint64 challenge = pdp.GenChallenge();
        bool result = pdp.VerifyProofWithMerklePathForFile(challenge);
        if (!result) {
            return false;
        }
        return true;
    }

    function checkProveExpire(uint256 fileExpiredHeight)
        public
        view
        returns (bool)
    {
        if (block.number > fileExpiredHeight) {
            return true;
        }
        return false;
    }

    function FileProve(FileProveParams memory fileProve) public {
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
        ProveDetail[] memory details = getProveDetailsWithNodeAddr(
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
            // TODO
            // details.push(fileProve.FileHash, detail);
        }
        // TODO
        // UpdateProveDetailInfo();
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
            // TODO
            // sector.AddFileToSector(sectorInfo, fileInfo);
            // TODO
            // sector.AddSectorRefForFileInfo(fileINfo, sectorInfo);
            if (sectorInfo.NextProveHeight == 0) {
                sectorInfo.NextProveHeight = fileProve.BlockHeight + fileInfo.ProveInterval;
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
                // TODO
                // sector.deleteFileFromSector(sectorInfo, fileInfo);
            }
            // TODO
            // settleForFile(fileInfo, nodeInfo, proveDetail, proveDetails, setting);
        }
        emit FilePDPSuccessEvent(
            FsEvent.FILE_PDP_SUCCESS,
            block.number,
            fileInfo.FileHash,
            nodeInfo.WalletAddr
        );
    }

    struct SectorProveParams {
        address NodeAddr;
        uint64 SectorID;
        uint64 ChallengeHeight;
        bytes ProveData;
    }

    function checkSectorProve(
        SectorProveParams memory sectorProve,
        SectorInfo memory sectorInfo
    ) public view returns (bool) {
        // TODO
    }

    function punishForSector(
        SectorInfo memory sectorInfo,
        NodeInfo memory nodeInfo,
        Setting memory setting,
        uint64 times
    ) public {
        // TODO
    }

    function profitSplitForSector(
        SectorInfo memory sectorInfo,
        NodeInfo memory nodeInfo,
        Setting memory setting
    ) public {
        // TODO
    }

    function SectorProve(SectorProveParams memory sectorProve) public payable {
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
            revert SectorProveFailed();
        }
        if (sectorProve.ChallengeHeight != sectorInfo.NextProveHeight) {
            revert SectorProveFailed();
        }
        bool r = checkSectorProve(sectorProve, sectorInfo);
        if (!r) {
            punishForSector(sectorInfo, nodeInfo, setting, 1);
            revert SectorProveFailed();
        }
        profitSplitForSector(sectorInfo, nodeInfo, setting);
        if (sectorInfo.FirstProveHeight == 0) {
            sectorInfo.FirstProveHeight = block.number;
        }
        sectorInfo.NextProveHeight = block.number + setting.DefaultProvePeriod;
        sector.UpdateSectorInfo(sectorInfo);
        if (!sectorInfo.IsPlots) {
            revert SectorProveFailed();
        }
        // TODO poc prove
    }

    function getLastPunishmentHeightForNode(address nodeAddr, uint64 sectorID)
        public
        view
        returns (uint64)
    {
        // TODO
    }

    function calMissingSectorProveTimes(
        SectorInfo memory sectorInfo,
        Setting memory setting,
        uint256 lastHeight,
        uint256 nowHeight
    ) public view returns (uint64) {
        // TODO
    }

    function CheckNodeSectorProvedInTime(SectorRef memory sectorRef)
        public
        payable
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
        uint256 lastHeight = getLastPunishmentHeightForNode(nodeAddr, sectorID);
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
}
