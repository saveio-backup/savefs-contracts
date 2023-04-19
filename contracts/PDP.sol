//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract PDP is Initializable, IPDP, IFsEvent {
    using ProofsMapping for ProofsPool;
    using ChallengeMapping for ChallengePool;
    using MerkleNodeMapping for MerkleNodePool;
    using MerklePathMapping for MerklePathPool;

    ProofsPool proofsPool; // proofs record
    mapping(bytes => ChallengePool) challengeMap; // key => map(chKey => Challenge)
    mapping(bytes => MerklePathPool) merklePathMap; // key => map(mpKey => mnKeys)
    MerkleNodePool merkleNodeMap; // mnKeys => MerkleNode

    event VerifyProofWithMerklePathForFileEvent();

    function initialize() public initializer {}

    function GenChallenge(GenChallengeParams memory gParams)
        public
        view
        virtual
        override
        returns (Challenge[] memory)
    {
        bytes memory blockHashArray = gParams.HashValue;
        bytes memory plant = abi.encodePacked(
            gParams.WalletAddr,
            blockHashArray
        );
        bytes32 hash = sha256(plant);

        bytes memory tmpHash = new bytes(36);
        for (uint32 i = 0; i < 32; i++) {
            tmpHash[i] = hash[i];
        }
        for (uint32 i = 0; i < 4; i++) {
            tmpHash[i + 32] = hash[i];
        }

        uint64 blockNumPerPart;
        uint64 blockNumLastPart;
        uint64 blockNumOfPart;

        if (gParams.FileBlockNum <= 3) {
            blockNumPerPart = 1;
            blockNumLastPart = 1;
            blockNumOfPart = 1;
            gParams.ProveNum = gParams.FileBlockNum;
        } else {
            if (
                gParams.FileBlockNum > 3 &&
                gParams.FileBlockNum < gParams.ProveNum
            ) {
                gParams.ProveNum = 3;
            }
            blockNumPerPart = gParams.FileBlockNum / gParams.ProveNum;
            blockNumLastPart =
                blockNumPerPart +
                (gParams.FileBlockNum % gParams.ProveNum);
            blockNumOfPart = blockNumPerPart;
        }

        Challenge[] memory challenges = new Challenge[](gParams.ProveNum);
        bytes32 blockHash = hash;
        uint32 hashIndex = 0;
        for (uint32 i = 1; i <= gParams.ProveNum; i++) {
            if (i == gParams.ProveNum) {
                blockNumOfPart = blockNumLastPart;
            }

            bytes memory tmp = new bytes(4);
            for (uint32 j = 0; j < 4; j++) {
                tmp[j] = tmpHash[hashIndex + j];
            }
            uint32 rd = uint32(bytesToUint(tmp));
            challenges[i - 1].Index = uint32(
                ((rd + 1) % blockNumOfPart) + (i - 1) * blockNumPerPart
            );
            challenges[i - 1].Rand = uint8(blockHash[hashIndex]) + 1;

            hashIndex++;
            hashIndex = hashIndex % 32;
        }

        return challenges;
    }

    function bytesToUint(bytes memory b) public pure returns (uint256) {
        uint256 number;
        for (uint256 i = 0; i < b.length; i++) {
            number = number + uint8(b[i]) * (2**(8 * (b.length - (i + 1))));
        }
        return number;
    }

    function genChallengeKey(GenChallengeParams memory gParams)
        public
        pure
        returns (string memory)
    {
        string memory key = string(
            abi.encodePacked(
                gParams.WalletAddr,
                gParams.HashValue,
                gParams.FileBlockNum,
                gParams.ProveNum
            )
        );
        return key;
    }

    function PrepareForPdpVerification(
        PrepareForPdpVerificationParams memory pParams
    ) public view virtual override returns (PdpVerificationReturns memory) {
        // TODO
        console.log(pParams.SectorInfo_.NodeAddr);
        PdpVerificationReturns memory pReturns;
        pReturns.Success = true;
        return pReturns;
    }

    function VerifyProof(
        ProofRecord memory vParams,
        Challenge[] memory chgs,
        MerklePath[] memory mp
    ) public payable virtual override {
        ProofRecord memory pr;
        pr.Proof.Version = vParams.Proof.Version;
        pr.Proof.Proofs = vParams.Proof.Proofs;
        pr.Proof.FileIds = vParams.Proof.FileIds;
        pr.Proof.Tags = vParams.Proof.Tags;
        pr.Proof.RootHashes = vParams.Proof.RootHashes;
        pr.State = vParams.State;
        pr.LastUpdateHeight = block.number;
        bytes memory key = GetKeyByProofParams(vParams.Proof);
        SaveChallenge(key, chgs);
        SaveMerklePath(key, mp);
        proofsPool.insert(key, pr);
    }

    function SaveChallenge(bytes memory key, Challenge[] memory chgs)
        public
        payable
        virtual
    {
        for (uint32 i = 0; i < chgs.length; i++) {
            bytes memory chKey = GetChallengeKey(chgs[i]);
            challengeMap[key].insert(chKey, chgs[i]);
        }
    }

    function SaveMerklePath(bytes memory key, MerklePath[] memory mp)
        public
        payable
        virtual
    {
        for (uint32 i = 0; i < mp.length; i++) {
            bytes memory mpKey = GetMerklePathKey(mp[i]);
            bytes[] memory keys = new bytes[](mp[i].Path.length);
            for (uint32 j = 0; j < mp[i].Path.length; j++) {
                bytes memory nodeKey = GetMerkleNodeKey(mp[i].Path[j]);
                merkleNodeMap.insert(nodeKey, mp[i].Path[j]);
                keys[j] = nodeKey;
            }
            merklePathMap[key].insert(mpKey, keys);
        }
    }

    function SubmitVerifyProofRequest(
        ProofParams memory vParams,
        Challenge[] memory chgs,
        MerklePath[] memory mp
    ) public payable virtual override {
        bytes memory key = GetKeyByProofParams(vParams);
        SaveChallenge(key, chgs);
        SaveMerklePath(key, mp);
        ProofRecord memory pr;
        pr.State = false;
        proofsPool.insert(key, pr);
        emit PDPVerifyEvent(FsEvent.PROOF_REQUEST);
    }

    function GetUnVerifyProofList()
        public
        view
        virtual
        override
        returns (ProofRecordWithParams[] memory)
    {
        uint256 count = 0;
        for (
            uint256 i = proofsPool.iterate_start();
            proofsPool.iterate_valid(i);
            i = proofsPool.iterate_next(i)
        ) {
            (, ProofRecord memory pr) = proofsPool.iterate_get(i);
            if (!pr.State) {
                count++;
            }
        }
        ProofRecordWithParams[] memory prList = new ProofRecordWithParams[](
            count
        );
        ProofParams memory vParams;
        Challenge[] memory chgs;
        MerklePath[] memory mp;
        for (
            uint256 i = proofsPool.iterate_start();
            proofsPool.iterate_valid(i);
            i = proofsPool.iterate_next(i)
        ) {
            (, ProofRecord memory pr) = proofsPool.iterate_get(i);
            if (pr.State) {
                continue;
            }
            vParams = pr.Proof;
            bytes memory key = GetKeyByProofParams(vParams);
            chgs = GetChallengeList(key);
            mp = GetMerklePathList(key);
            prList[i] = ProofRecordWithParams(vParams, chgs, mp);
        }
        return prList;
    }

    function GetChallengeList(bytes memory key)
        public
        view
        virtual
        returns (Challenge[] memory)
    {
        Challenge[] memory chgs = new Challenge[](
            challengeMap[key].keys.length
        );
        for (
            uint256 i = challengeMap[key].iterate_start();
            challengeMap[key].iterate_valid(i);
            i = challengeMap[key].iterate_next(i)
        ) {
            (, Challenge memory value) = challengeMap[key].iterate_get(i);
            chgs[i] = value;
        }
        return chgs;
    }

    function GetMerklePathList(bytes memory key)
        public
        view
        virtual
        returns (MerklePath[] memory)
    {
        MerklePath[] memory mp = new MerklePath[](
            merklePathMap[key].keys.length
        );
        for (
            uint256 i = merklePathMap[key].iterate_start();
            merklePathMap[key].iterate_valid(i);
            i = merklePathMap[key].iterate_next(i)
        ) {
            (, bytes[] memory value) = merklePathMap[key].iterate_get(i);
            mp[i].Path = new MerkleNode[](value.length);
            for (uint256 j = 0; j < value.length; j++) {
                mp[i].Path[j] = merkleNodeMap.data[value[j]].value;
            }
        }
        return mp;
    }

    function VerifyProofWithMerklePathForFile(ProofParams memory vParams)
        public
        view
        virtual
        override
        returns (bool)
    {
        bytes memory key = GetKeyByProofParams(vParams);
        ProofRecord memory pr = proofsPool.data[key].value;
        return pr.State;
    }

    function VerifyPlotData(VerifyPlotDataParams memory vParams)
        public
        view
        virtual
        override
        returns (bool)
    {
        // TODO
        console.log(vParams.Index);
        return true;
    }

    function GetChallengeKey(Challenge memory chg)
        public
        pure
        returns (bytes memory)
    {
        bytes memory key = abi.encodePacked(chg.Index, chg.Rand);
        return key;
    }

    function GetMerklePathKey(MerklePath memory mp)
        public
        pure
        returns (bytes memory)
    {
        bytes memory key;
        for (uint32 i = 0; i < mp.Path.length; i++) {
            key = abi.encodePacked(mp.Path[i].Hash, mp.Path[i].Index);
        }
        return key;
    }

    function GetMerkleNodeKey(MerkleNode memory mn)
        public
        pure
        returns (bytes memory)
    {
        bytes memory key = abi.encodePacked(mn.Hash, mn.Index);
        return key;
    }

    function GetKeyByProofParams(ProofParams memory vParams)
        public
        pure
        returns (bytes memory)
    {
        bytes memory ids;
        for (uint32 i = 0; i < vParams.FileIds.length; i++) {
            ids = abi.encodePacked(ids, vParams.FileIds[i]);
        }
        bytes memory tags;
        for (uint32 i = 0; i < vParams.Tags.length; i++) {
            tags = abi.encodePacked(tags, vParams.Tags[i]);
        }
        string memory keyStr = string(
            abi.encodePacked(
                vParams.Version,
                vParams.Proofs,
                ids,
                tags,
                vParams.RootHashes
            )
        );
        bytes memory keyBytes = bytes(keyStr);
        bytes32 key32 = keccak256(keyBytes);
        bytes memory key = abi.encodePacked(key32);
        return key;
    }
}

// --------------------------library--------------------------
// map
struct IndexValue {
    uint256 keyIndex;
    ProofRecord value;
}
struct KeyFlag {
    bytes key;
    bool deleted;
}
struct ProofsPool {
    mapping(bytes => IndexValue) data;
    KeyFlag[] keys;
    uint256 size;
}

library ProofsMapping {
    function insert(
        ProofsPool storage self,
        bytes memory key,
        ProofRecord memory value
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

    function remove(ProofsPool storage self, bytes memory key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(ProofsPool storage self, bytes memory key)
        internal
        view
        returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(ProofsPool storage self)
        internal
        view
        returns (uint256 keyIndex)
    {
        uint256 index = iterate_next(self, type(uint256).min);
        return index - 1;
    }

    function iterate_valid(ProofsPool storage self, uint256 keyIndex)
        internal
        view
        returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(ProofsPool storage self, uint256 keyIndex)
        internal
        view
        returns (uint256 r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(ProofsPool storage self, uint256 keyIndex)
        internal
        view
        returns (bytes memory key, ProofRecord memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}

// challenge map
struct IndexValueCH {
    uint256 keyIndex;
    Challenge value;
}
struct KeyFlagCh {
    bytes key;
    bool deleted;
}
struct ChallengePool {
    mapping(bytes => IndexValueCH) data;
    KeyFlagCh[] keys;
    uint256 size;
}

library ChallengeMapping {
    function insert(
        ChallengePool storage self,
        bytes memory key,
        Challenge memory value
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

    function remove(ChallengePool storage self, bytes memory key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(ChallengePool storage self, bytes memory key)
        internal
        view
        returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(ChallengePool storage self)
        internal
        view
        returns (uint256 keyIndex)
    {
        uint256 index = iterate_next(self, type(uint256).min);
        return index - 1;
    }

    function iterate_valid(ChallengePool storage self, uint256 keyIndex)
        internal
        view
        returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(ChallengePool storage self, uint256 keyIndex)
        internal
        view
        returns (uint256 r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(ChallengePool storage self, uint256 keyIndex)
        internal
        view
        returns (bytes memory key, Challenge memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}

// merkle node map
struct IndexValueMN {
    uint256 keyIndex;
    MerkleNode value;
}
struct KeyFlagMN {
    bytes key;
    bool deleted;
}
struct MerkleNodePool {
    mapping(bytes => IndexValueMN) data;
    KeyFlagMN[] keys;
    uint256 size;
}

library MerkleNodeMapping {
    function insert(
        MerkleNodePool storage self,
        bytes memory key,
        MerkleNode memory value
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

    function remove(MerkleNodePool storage self, bytes memory key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(MerkleNodePool storage self, bytes memory key)
        internal
        view
        returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(MerkleNodePool storage self)
        internal
        view
        returns (uint256 keyIndex)
    {
        uint256 index = iterate_next(self, type(uint256).min);
        return index - 1;
    }

    function iterate_valid(MerkleNodePool storage self, uint256 keyIndex)
        internal
        view
        returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(MerkleNodePool storage self, uint256 keyIndex)
        internal
        view
        returns (uint256 r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(MerkleNodePool storage self, uint256 keyIndex)
        internal
        view
        returns (bytes memory key, MerkleNode memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}

// merkle path map
struct IndexValueMP {
    uint256 keyIndex;
    bytes[] value;
}
struct KeyFlagMP {
    bytes key;
    bool deleted;
}
struct MerklePathPool {
    mapping(bytes => IndexValueMP) data;
    KeyFlagMP[] keys;
    uint256 size;
}

library MerklePathMapping {
    function insert(
        MerklePathPool storage self,
        bytes memory key,
        bytes[] memory value
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

    function remove(MerklePathPool storage self, bytes memory key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(MerklePathPool storage self, bytes memory key)
        internal
        view
        returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(MerklePathPool storage self)
        internal
        view
        returns (uint256 keyIndex)
    {
        uint256 index = iterate_next(self, type(uint256).min);
        return index - 1;
    }

    function iterate_valid(MerklePathPool storage self, uint256 keyIndex)
        internal
        view
        returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(MerklePathPool storage self, uint256 keyIndex)
        internal
        view
        returns (uint256 r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(MerklePathPool storage self, uint256 keyIndex)
        internal
        view
        returns (bytes memory key, bytes[] memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
