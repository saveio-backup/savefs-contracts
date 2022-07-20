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
        view
        virtual
        override
        returns (Challenge[] memory)
    {
        // TODO
        console.log(gParams.WalletAddr);
        Challenge[] memory challenges;
        return challenges;
    }

    function PrepareForPdpVerification(
        PrepareForPdpVerificationParams memory pParams
    ) public view returns (PdpVerificationReturns memory) {
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
        returns (bool)
    {
        // TODO
        console.log(vParams.Index);
        return true;
    }
}
