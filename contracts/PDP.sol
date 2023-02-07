//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract PDP is Initializable, IPDP, IFsEvent {
    using IterableMapping for ProofsPool;
    ProofsPool proofsPool;

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

    function VerifyProofWithMerklePathForFile(ProofParams memory vParams)
        public
        view
        virtual
        override
        returns (bool)
    {
        // TODO
        console.log(vParams.Version);
        return true;
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
        bytes memory challenges;
        for (uint32 i = 0; i < vParams.Challenges.length; i++) {
            challenges = abi.encodePacked(
                challenges,
                vParams.Challenges[i].Index,
                vParams.Challenges[i].Rand
            );
        }
        bytes memory merklePath;
        for (uint32 i = 0; i < vParams.MerklePath_.length; i++) {
            merklePath = abi.encodePacked(
                merklePath,
                vParams.MerklePath_[i].PathLen
            );
        }
        string memory keyStr = string(
            abi.encodePacked(
                vParams.Version,
                vParams.Proofs,
                ids,
                tags,
                challenges,
                merklePath,
                vParams.RootHashes
            )
        );
        bytes memory keyBytes = bytes(keyStr);
        bytes32 key32 = keccak256(keyBytes);
        bytes memory key = abi.encodePacked(key32);
        return key;
    }
}

// map
struct IndexValue {
    uint256 keyIndex;
    ProofsRecord value;
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

library IterableMapping {
    function insert(
        ProofsPool storage self,
        bytes memory key,
        ProofsRecord memory value
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
        returns (bytes memory key, ProofsRecord memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
