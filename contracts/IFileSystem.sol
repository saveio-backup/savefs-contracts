//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "./Struct.sol";

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

    function FsNodeQuery(address walletAddr)
        public
        view
        virtual
        returns (FsNodeInfo memory);

    function FsNodeUpdate(FsNodeInfo memory fsNodeInfo) public payable virtual;
}
