//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract PDP is Initializable, IPDP {
    function initialize() public initializer {}

    function GenChallenge(GenChallengeParams memory gParams)
        public
        pure
        virtual
        override
        returns (Challenge[] memory)
    {
        // TODO
        Challenge[] memory challenges;
        return challenges;
    }

    function PrepareForPdpVerification(
        PrepareForPdpVerificationParams memory pParams
    ) public pure returns (PdpVerificationReturns memory) {
        PdpVerificationReturns memory pReturns;

        // TODO
        pReturns.Success = true;
        return pReturns;
    }

    function VerifyProofWithMerklePathForFile(
        VerifyProofWithMerklePathForFileParams memory vParams
    ) public pure virtual override returns (bool) {
        // TODO
        return true;
    }

    function VerifyPlotData(VerifyPlotDataParams memory vParams)
        public
        pure
        returns (bool)
    {
        // TODO
        return true;
    }
}
