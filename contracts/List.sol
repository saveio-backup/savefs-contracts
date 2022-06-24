//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract List is Initializable, IList {
    mapping(bytes => WhiteList[]) whiteList; // fileHash => whileList

    function initialize() public initializer {}

    function WhiteListOperate(WhiteListParams memory params)
        public
        virtual
        override
    {
        if (params.Op == WhiteListOpType.ADD) {
            WhiteList[] storage list = whiteList[params.FileHash];
            for (uint256 i = 0; i < params.List.length; i++) {
                list.push(params.List[i]);
            }
        }
        if (params.Op == WhiteListOpType.DEL) {
            WhiteList[] storage list = whiteList[params.FileHash];
            for (uint256 i = 0; i < params.List.length; i++) {
                for (uint256 j = 0; j < list.length; j++) {
                    if (list[j].Addr == params.List[i].Addr) {
                        delete list[j];
                        break;
                    }
                }
            }
        }
        if (params.Op == WhiteListOpType.ADD_COV) {
            WhiteList[] storage list = whiteList[params.FileHash];
            for (uint256 i = 0; i < list.length; i++) {
                if (list[i].ExpireHeight <= list[i].BaseHeight) {
                    delete list[i];
                }
            }
            whiteList[params.FileHash] = list;
        }
        if (params.Op == WhiteListOpType.DEL_ALL) {
            delete whiteList[params.FileHash];
        }
    }

    function GetWhiteList(bytes memory fileHash)
        public
        view
        virtual
        override
        returns (WhiteList[] memory)
    {
        require(fileHash.length > 0, "fileHash must be empty");
        require(whiteList[fileHash].length > 0, "whiteList is empty");
        return whiteList[fileHash];
    }
}
