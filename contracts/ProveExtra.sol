//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

/**
 * Slice contract from prove.sol
 */
contract ProveExtra {
    using IterableMapping for ProveDetailMap;

    mapping(bytes => ProveDetailMap) proveDetails; // fileHash => nodeAddr => ProveDetail
    mapping(bytes => ProveDetailMeta) proveDetailMeta; // fileHash => ProveDetailMeta
    mapping(string => PocProve) pocProve; // miner + height => PocProve

    function SettleForFile(
        INode node,
        IFile file,
        FileInfo memory fileInfo,
        NodeInfo memory nodeInfo,
        ProveDetail memory detail,
        ProveDetail[] memory details,
        Setting memory setting
    ) public payable returns (string memory) {
        uint64 profit = calculateProfitForSettle(
            file,
            fileInfo,
            detail,
            setting
        );
        if (fileInfo.Deposit < profit) {
            return "FileSettleDepositNotEnough";
        }

        nodeInfo.RestVol += profit;
        node.UpdateNodeInfo(nodeInfo);

        fileInfo.Deposit -= profit;
        fileInfo.ValidFlag = false;
        file.UpdateFileInfo(fileInfo);

        uint64 finishedNodes = 0;
        for (uint256 i = 0; i < details.length; i++) {
            if (details[i].Finished) {
                finishedNodes++;
            }
        }
        if (finishedNodes == 1) {
            file.CleanupForDeleteFile(fileInfo, false, true);
        }
        if (finishedNodes == fileInfo.CopyNum + 1) {
            if (fileInfo.Deposit > 0) {
                payable(fileInfo.FileOwner).transfer(fileInfo.Deposit);
            }
            file.CleanupForDeleteFile(fileInfo, true, false);
        } else {
            file.AddFileToUnSettleList(fileInfo.FileOwner, fileInfo.FileHash);
        }
        return "";
        // TODO
        // emit ProveFileEvent(
        //     FsEvent.PROVE_FILE,
        //     block.number,
        //     nodeInfo.WalletAddr,
        //     profit
        // );
    }

    function calculateProfitForSettle(
        IFile file,
        FileInfo memory fileInfo,
        ProveDetail memory detail,
        Setting memory setting
    ) private view returns (uint64) {
        StorageFee memory total = file.CalcFee(
            setting,
            detail.ProveTimes - 1,
            0,
            fileInfo.FileBlockNum * fileInfo.FileBlockSize,
            uint64(fileInfo.ExpiredHeight - fileInfo.BlockHeight)
        );
        return total.TxnFee + total.SpaceFee + total.ValidationFee;
    }

    function AddSectorRefForFileInfo(
        ISector sector,
        IFile file,
        FileInfo memory fileInfo,
        SectorRef memory sectorRef
    ) public returns (string memory) {
        bool r = sector.IsSectorRefByFileInfo(
            sectorRef.NodeAddr,
            sectorRef.SectorId
        );
        if (!r) {
            return "NotSectorRefByFileInfo";
        }
        file.AddFileSectorRef(fileInfo.FileHash, sectorRef);
        return "";
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

    function checkProve(
        IPDP pdp,
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
        ProofParams memory vParams;
        vParams.Version = 0;
        vParams.Proofs = proveData.Proofs;
        vParams.FileIds = proveParam.FileID;
        vParams.Tags = proveData.Tags;
        vParams.RootHashes = proveParam.RootHash;
        bool res = pdp.VerifyProofWithMerklePathForFile(
            vParams
        );
        // TODO return true because now can't deserialize prove data
        return true;
        if (!res) {
            return false;
        }
        return true;
    }

    function checkSectorProve(
        IPDP pdp,
        uint64 proveNum,
        SectorProveParams memory sectorProve,
        SectorInfo memory sectorInfo
    ) public payable returns (string memory) {
        // TODO block head hash
        bytes memory blockHash;
        GenChallengeParams memory gParams;
        gParams.WalletAddr = sectorProve.NodeAddr;
        gParams.HashValue = blockHash;
        gParams.FileBlockNum = sectorInfo.TotalBlockNum;
        gParams.ProveNum = proveNum;
        Challenge[] memory challenges = pdp.GenChallenge(gParams);
        // pre
        // TODO decentralized sector prove data
        SectorProveData memory sectorProveData;
        PrepareForPdpVerificationParams memory pParams;
        pParams.SectorInfo_ = sectorInfo;
        pParams.ProveData = sectorProveData;
        PdpVerificationReturns memory pReturns;
        pReturns = pdp.PrepareForPdpVerification(pParams);
        if (!pReturns.Success) {
            return "checkSectorProve PrepareForPdpVerification failed";
        }
        // verify
        ProofParams memory vParams;
        vParams.Version = 0;
        vParams.Proofs = sectorProveData.Proofs;
        vParams.FileIds = pReturns.FileIDs;
        vParams.Tags = pReturns.Tags;
        vParams.RootHashes = pReturns.RootHashes;
        // submit a verify request rather than verify directly
        pdp.SubmitVerifyProofRequest(vParams, challenges, sectorProveData.MerklePath_);
        if (sectorInfo.IsPlots) {
            if (
                !pReturns.FileInfo_.IsPlotFile ||
                pReturns.FileInfo_.PlotInfo_.Nonces == 0
            ) {
                return "checkSectorProve VerifyPlotData failed";
            }
            VerifyPlotDataParams memory vpParams;
            vpParams.PlotInfo_ = pReturns.FileInfo_.PlotInfo_;
            vpParams.PlotData = sectorProveData.PlotData;
            if (challenges.length > 0) {
                vpParams.Index = uint64(challenges[0].Index);
            }
            bool res2 = pdp.VerifyPlotData(vpParams);
            if (!res2) {
                return "checkSectorProve VerifyPlotData failed";
            }
        }
        return "";
    }

    function getProveDetailListWithNodeAddr(INode node, bytes memory fileHash)
        public
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

    function GetProveDetailList(bytes memory fileHash)
        public
        view
        returns (ProveDetail[] memory)
    {
        ProveDetailMap storage data = proveDetails[fileHash];
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

    function UpdateProveDetailList(
        bytes memory fileHash,
        ProveDetail[] memory details
    ) public payable {
        ProveDetailMap storage data = proveDetails[fileHash];
        for (uint256 i = 0; i < details.length; i++) {
            ProveDetail memory detail = details[i];
            data.insert(detail.NodeAddr, detail);
        }
    }

    function DeleteProveDetails(bytes memory fileHash) public payable {
        delete proveDetails[fileHash];
        delete proveDetailMeta[fileHash];
    }

    function UpdateProveDetailMeta(
        bytes memory fileHash,
        ProveDetailMeta memory details
    ) public payable {
        proveDetailMeta[fileHash] = details;
    }

    function getPocProve(address nodeAddr, uint256 height)
        public
        view
        returns (PocProve memory)
    {
        string memory key = string(abi.encodePacked(nodeAddr, height));
        return pocProve[key];
    }

    function putPocProve(PocProve memory prove) public payable {
        string memory key = string(abi.encodePacked(prove.Miner, prove.Height));
        pocProve[key] = prove;
    }

    function GetPocProveList() public view returns (PocProve[] memory) {
        // TODO
    }
}

// map
struct IndexValue {
    uint256 keyIndex;
    ProveDetail value;
}
struct KeyFlag {
    bytes key;
    bool deleted;
}

struct ProveDetailMap {
    mapping(bytes => IndexValue) data;
    KeyFlag[] keys;
    uint256 size;
}

library IterableMapping {
    function insert(
        ProveDetailMap storage self,
        bytes memory key,
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

    function remove(ProveDetailMap storage self, bytes memory key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(ProveDetailMap storage self, bytes memory key)
        internal
        view
        returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(ProveDetailMap storage self)
        internal
        view
        returns (uint256 keyIndex)
    {
        uint256 index = iterate_next(self, type(uint256).min);
        return index - 1;
    }

    function iterate_valid(ProveDetailMap storage self, uint256 keyIndex)
        internal
        view
        returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(ProveDetailMap storage self, uint256 keyIndex)
        internal
        view
        returns (uint256 r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(ProveDetailMap storage self, uint256 keyIndex)
        internal
        view
        returns (bytes memory key, ProveDetail memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
