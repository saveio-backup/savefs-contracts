//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract PDP is Initializable, IPDP, IFsEvent {
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

    function VerifyProofWithMerklePathForFile(
        VerifyProofWithMerklePathForFileParams memory vParams
    ) public view virtual override returns (bool) {
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
}
