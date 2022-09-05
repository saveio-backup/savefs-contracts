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
    address key;
    bool deleted;
}

struct ProveDetailMap {
    mapping(address => IndexValue) data;
    KeyFlag[] keys;
    uint256 size;
}

library IterableMapping {
    function insert(
        ProveDetailMap storage self,
        address key,
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

    function remove(ProveDetailMap storage self, address key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(ProveDetailMap storage self, address key)
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
        returns (address key, ProveDetail memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
