//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "./Types.sol";

abstract contract IFileSystem {
    function GetSetting() public pure virtual returns (Setting memory);

    function CalculateNodePledge(NodeInfo memory nodeInfo)
        public
        pure
        virtual
        returns (uint64);

    function NodeRegister(NodeInfo memory nodeInfo) public payable virtual;

    function NodeUpdate(NodeInfo memory nodeInfo) public payable virtual;

    function NodeCancel(address walletAddr) public virtual;

    function GetNodeList() public view virtual returns (NodeInfo[] memory);

    function GetNodeInfoByWalletAddr(address walletAddr)
        public
        view
        virtual
        returns (NodeInfo memory);

    function GetNodeInfoByNodeAddr(address nodeAddr)
        public
        view
        virtual
        returns (NodeInfo memory);

    function NodeWithDrawProfit(address walletAddr) public virtual;

    function GetUploadStorageFee(UploadOption memory uploadOption)
        public
        view
        virtual
        returns (StorageFee memory);

    function GetFileInfo(bytes memory fileHash)
        public
        view
        virtual
        returns (FileInfo memory);

    function GetFileInfos(FileList memory fileList)
        public
        view
        virtual
        returns (FileInfo[] memory);

    function GetFileList(address walletAddr)
        public
        view
        virtual
        returns (FileList memory);

    function GetWhiteList(bytes memory fileHash)
        public
        view
        virtual
        returns (WhiteList memory);

    function GetUserSpace(address walletAddr)
        public
        view
        virtual
        returns (UserSpace memory);
}
