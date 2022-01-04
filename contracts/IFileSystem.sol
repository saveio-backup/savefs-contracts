//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "./Struct.sol";
import "./Error.sol";

abstract contract IFileSystem {
    function FsGetSettings() public pure virtual returns (FsSetting memory);

    function FsGetUploadStorageFee(UploadOption memory uploadOption)
        public
        view
        virtual
        returns (StorageFee memory);

    function FsNodeRegister(FsNodeInfo memory fsNodeInfo)
        public
        payable
        virtual;
}
