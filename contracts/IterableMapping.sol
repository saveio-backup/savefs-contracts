//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "./type.sol";

struct IndexValue {
    uint256 keyIndex;
    ProveDetail value;
}
struct KeyFlag {
    address key;
    bool deleted;
}

struct ItMap {
    mapping(address => IndexValue) data;
    KeyFlag[] keys;
    uint256 size;
}

library IterableMapping {
    function insert(
        ItMap storage self,
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

    function remove(ItMap storage self, address key)
        internal
        returns (bool success)
    {
        uint256 keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0) return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size--;
    }

    function contains(ItMap storage self, address key)
        internal
        view
        returns (bool)
    {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(ItMap storage self)
        internal
        view
        returns (uint256 keyIndex)
    {
        uint256 index = iterate_next(self, type(uint256).min);
        return index - 1;
    }

    function iterate_valid(ItMap storage self, uint256 keyIndex)
        internal
        view
        returns (bool)
    {
        return keyIndex < self.keys.length;
    }

    function iterate_next(ItMap storage self, uint256 keyIndex)
        internal
        view
        returns (uint256 r_keyIndex)
    {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(ItMap storage self, uint256 keyIndex)
        internal
        view
        returns (address key, ProveDetail memory value)
    {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
