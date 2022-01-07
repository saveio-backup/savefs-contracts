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

    function CalculateNodePledge(FsNodeInfo memory fsNodeInfo)
        public
        pure
        virtual
        returns (uint64);

    function FsNodeRegister(FsNodeInfo memory fsNodeInfo)
        public
        payable
        virtual;

    function FsNodeUpdate(FsNodeInfo memory fsNodeInfo) public payable virtual;

    function FsNodeCancel(address walletAddr) public virtual;

    function FsGetNodeList() public view virtual returns (FsNodeInfo[] memory);

    function FsGetNodeInfoByWalletAddr(address walletAddr)
        public
        view
        virtual
        returns (FsNodeInfo memory);

    function FsGetNodeInfoByNodeAddr(address nodeAddr)
        public
        view
        virtual
        returns (FsNodeInfo memory);

    function FsNodeWithDrawProfit(address walletAddr) public virtual;
}
