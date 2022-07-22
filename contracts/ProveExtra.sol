//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";
import "./IterableMapping.sol";

/**
 * Slice contract from prove.sol
 */
contract ProveExtra {
    using IterableMapping for ItMap;

    mapping(bytes => ItMap) proveDetails; // fileHash => nodeAddr => ProveDetail
    mapping(bytes => ProveDetailMeta) proveDetailMeta; // fileHash => ProveDetailMeta
    mapping(string => PocProve) pocProve; // miner + height => PocProve

    function GetProveDetailList(bytes memory fileHash)
        public
        view
        returns (ProveDetail[] memory)
    {
        ItMap storage data = proveDetails[fileHash];
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
        ItMap storage data = proveDetails[fileHash];
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
