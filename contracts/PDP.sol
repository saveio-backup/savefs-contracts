//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";

contract PDP is Initializable {
    function initialize() public initializer {}

    function GenChallenge() public pure returns (uint64) {
        // TODO
        return 0;
    }

    function VerifyProofWithMerklePathForFile(uint64)
        public
        pure
        returns (bool)
    {
        // TODO
        return true;
    }
}
