//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";

/**
 * Slice contract from prove.sol
 */
contract ProvePocProve {
    mapping(string => PocProve) pocProve; // miner + height => PocProve

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
