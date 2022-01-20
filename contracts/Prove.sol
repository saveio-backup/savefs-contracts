//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./Config.sol";
import "./FileSystem.sol";
import "./Node.sol";
import "./PDP.sol";

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

    mapping(bytes => ProveDetail[]) proveDetail; // fileHash => ProveDetail[]
    mapping(bytes => ProveDetails) proveDetails; // fileHash => ProveDetails

    event FilePDPSuccessEvent(
        FsEvent eventType,
        uint256 blockHeight,
        bytes fileHash,
        address walletAddr
    );

    error FileProveNotFileOwner();
    error FileProveFailed();

    function initialize(
        Config _config,
        FileSystem _fs,
        Node _node,
        PDP _pdp
    ) public initializer {
        config = _config;
        fs = _fs;
        node = _node;
        pdp = _pdp;
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
                revert FileProveFailed();
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
                    revert FileProveFailed();
                }
            }
        }
        bool success = checkProve(fileProve, fileInfo);
        if (!success) {
            revert FileProveFailed();
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
                    revert FileProveFailed();
                }
                bool r = checkProveExpire(fileInfo.ExpiredHeight);
                if (r) {
                    revert FileProveFailed();
                }
                detail.ProveTimes++;
                found = true;
                break;
            }
        }
        if (!found) {
            if (details.length == fileInfo.CopyNum + 1) {
                revert FileProveFailed();
            }
            if (
                nodeInfo.RestVol <
                fileInfo.FileBlockNum * fileInfo.FileBlockSize
            ) {
                revert FileProveFailed();
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
            // TODO
        }
        if (settleFlag) {
            // TODO
        }
        emit FilePDPSuccessEvent(
            FsEvent.FILE_PDP_SUCCESS,
            block.number,
            fileInfo.FileHash,
            nodeInfo.WalletAddr
        );
    }
}
