//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./IFileSystem.sol";
import "./Types.sol";

contract FileSystem is Initializable, IFileSystem {
    /************************************************************************
     * Field define ******************************************************
     */
    mapping(address => NodeInfo) nodesInfo; // walletAddr => NodeInfo
    NodeList nodeList; // nodeAddr list
    mapping(address => mapping(uint64 => SectorInfo)) sectorInfos; // nodeAddr => (sectorId => SectorInfo)
    mapping(bytes => FileInfo) fileInfos; // fileHash => FileInfo
    mapping(address => FileList) fileList; // walletAddr => fileHash list
    mapping(bytes => WhiteList) whiteList; // fileHash => whileList
    mapping(address => UserSpace) userSpace; // walletAddr => UserSpace
    mapping(uint32 => PocProve[]) pocProveList; // blockNumber => PocProve

    /**********************************************************************
     * Event define *******************************************************
     */
    event StoreFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        uint64 fileSize,
        address walletAddr,
        uint64 cost,
        bool isPlotFile
    );

    event DeleteFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        address walletAddr
    );

    event DeleteFilesEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        address walletAddr
    );

    event SetUserSpaceEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sizeType,
        uint64 size,
        uint64 countType,
        uint64 count
    );

    event RegisterNodeEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        address nodeAddr,
        uint64 volume,
        uint64 serviceTime
    );

    event UnRegisterNodeEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr
    );

    event ProveFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 profit
    );

    event FilePDPSuccessEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        address walletAddr
    );

    event CreateSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId,
        uint64 provLevel,
        uint64 size,
        bool isPlots
    );

    event DeleteSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId
    );

    /************************************************************************
     * Error define ******************************************************
     */
    error NotEnoughPledge(uint256 got, uint256 want);
    error ZeroProfit();
    error FileNotExist(bytes);
    /************************************************************************
     * Modifier define ******************************************************
     */
    modifier VolumeRequire(NodeInfo memory nodeInfo) {
        require(
            nodeInfo.Volume >= GetSetting().MinVolume,
            "Volume is too small"
        );
        _;
    }

    modifier NodeRegisted(address walletAddr) {
        require(nodesInfo[walletAddr].Volume != 0, "Node not registered");
        _;
    }

    modifier NodeNotRegisted(address walletAddr) {
        require(nodesInfo[walletAddr].Volume == 0, "Node already registered");
        _;
    }

    modifier NotEmptyFileHash(bytes memory fileHash) {
        require(fileHash.length > 0, "fileHash must be empty");
        _;
    }

    /****************************************************************************
     * Constract method *********************************************************
     */
    function initialize() public initializer {
        console.log("initializer");
    }

    /****************************************************************************
     * Setting info mamanagement ************************************************
     */
    function GetSetting() public pure override returns (Setting memory) {
        Setting memory setting;
        setting.GasPrice = 1;
        setting.GasPerGBPerBlock = 1;
        setting.GasPerKBPerBlock = 1;
        setting.GasForChallenge = 200000;
        setting.MaxProveBlocks = 32;
        setting.MinVolume = 1000 * 1000;
        setting.DefaultProvePeriod = (3600 * 24) / 5;
        setting.DefaultProveLevel = 1;
        setting.DefaultCopyNum = 2;
        return setting;
    }

    /****************************************************************************
     * Node info mamanagement ***************************************************
     */
    function CalculateNodePledge(NodeInfo memory nodeInfo)
        public
        pure
        override
        returns (uint64)
    {
        Setting memory setting = GetSetting();
        return setting.GasPrice * setting.GasPerGBPerBlock * nodeInfo.Volume;
    }

    function NodeRegister(NodeInfo memory nodeInfo)
        public
        payable
        override
        VolumeRequire(nodeInfo)
        NodeNotRegisted(nodeInfo.WalletAddr)
    {
        uint64 pledge = CalculateNodePledge(nodeInfo);
        if (msg.value < pledge) {
            revert NotEnoughPledge(msg.value, pledge);
        }
        nodeInfo.Pledge = pledge;
        nodeInfo.Profit = 0;
        nodeInfo.RestVol = nodeInfo.Volume;
        nodesInfo[nodeInfo.WalletAddr] = nodeInfo;
        nodeList.AddrList.push(nodeInfo.WalletAddr);
        emit RegisterNodeEvent(
            FsEvent.REG_NODE,
            block.number,
            nodeInfo.WalletAddr,
            nodeInfo.NodeAddr,
            nodeInfo.Volume,
            nodeInfo.ServiceTime
        );
    }

    function NodeUpdate(NodeInfo memory nodeInfo)
        public
        payable
        override
        VolumeRequire(nodeInfo)
        NodeRegisted(nodeInfo.WalletAddr)
    {
        require(
            nodesInfo[nodeInfo.WalletAddr].WalletAddr == nodeInfo.WalletAddr,
            "Node walletAddr changed"
        );
        NodeInfo memory oldNode = nodesInfo[nodeInfo.WalletAddr];
        uint64 newPledge = CalculateNodePledge(nodeInfo);
        uint64 oldPledge = oldNode.Pledge;
        if (newPledge < oldPledge) {
            payable(nodeInfo.WalletAddr).transfer(oldPledge - newPledge);
        } else {
            uint64 pledge = newPledge - oldPledge;
            if (msg.value < pledge) {
                revert NotEnoughPledge(msg.value, pledge);
            }
        }
        nodeInfo.Pledge = newPledge;
        nodeInfo.Profit = oldNode.Profit;
        nodeInfo.RestVol = oldNode.RestVol + nodeInfo.Volume - oldNode.Volume;
        nodesInfo[nodeInfo.WalletAddr] = nodeInfo;
    }

    function NodeCancel(address walletAddr)
        public
        override
        NodeRegisted(walletAddr)
    {
        NodeInfo memory nodeInfo = nodesInfo[walletAddr];
        if (nodeInfo.Pledge > 0) {
            payable(nodeInfo.WalletAddr).transfer(
                nodeInfo.Pledge + nodeInfo.Profit
            );
        }
        // TODO delete sector
        delete nodesInfo[walletAddr];
        NodeListRemove(walletAddr);
        emit UnRegisterNodeEvent(FsEvent.UN_REG_NODE, block.number, walletAddr);
    }

    function NodeListRemove(address addr) public {
        for (uint256 i = 0; i < nodeList.AddrList.length - 1; i++) {
            if (nodeList.AddrList[i] == addr) {
                delete nodeList.AddrList[i];
                return;
            }
        }
    }

    /**
     * description: Actually I can't understand why do this
     *              mixed use nodeAddr and walletAddr
     *
     * @return nodeAddr => NodeInfo
     */
    function GetNodeList() public view override returns (NodeInfo[] memory) {
        NodeInfo[] memory _nodesInfo = new NodeInfo[](nodeList.AddrList.length);
        for (uint256 i = 0; i < nodeList.AddrList.length; i++) {
            _nodesInfo[i] = nodesInfo[nodeList.AddrList[i]];
        }
        return _nodesInfo;
    }

    function GetNodeInfoByWalletAddr(address walletAddr)
        public
        view
        override
        returns (NodeInfo memory)
    {
        return nodesInfo[walletAddr];
    }

    function GetNodeInfoByNodeAddr(address nodeAddr)
        public
        view
        override
        returns (NodeInfo memory)
    {
        return nodesInfo[nodeAddr];
    }

    function NodeWithDrawProfit(address walletAddr)
        public
        override
        NodeRegisted(walletAddr)
    {
        NodeInfo memory nodeInfo = nodesInfo[walletAddr];
        if (nodeInfo.Profit > 0) {
            payable(nodeInfo.WalletAddr).transfer(nodeInfo.Profit);
            nodeInfo.Profit = 0;
        } else {
            revert ZeroProfit();
        }
        nodesInfo[walletAddr] = nodeInfo;
    }

    /****************************************************************************
     * File info mamanagement ***************************************************
     */
    function GetUploadStorageFee(UploadOption memory uploadOption)
        public
        view
        override
        returns (StorageFee memory)
    {
        require(uploadOption.FileSize > 0, "fileSize must be greater than 0");
        Setting memory setting = GetSetting();
        StorageFee memory storageFee;
        uint64 fee;
        uint64 txGas = 10000000;
        if (uploadOption.WhiteList_.Num > 0) {
            fee = txGas * 4;
        } else {
            fee = txGas * 3;
        }

        uint64 proveTime = (uploadOption.ExpiredHeight - uint64(block.number)) /
            uploadOption.ProveInterval +
            1;
        uint64 validFee = (uploadOption.CopyNum + 1) *
            uint64(
                proveTime +
                    (setting.GasForChallenge * uploadOption.FileSize) /
                    1024000
            );
        uint64 spaceFee = ((uploadOption.CopyNum + 1) *
            setting.GasPerGBPerBlock *
            uploadOption.FileSize) / uint64(1024000);
        storageFee.TxnFee = fee;
        storageFee.ValidationFee = validFee;
        storageFee.SpaceFee = spaceFee;
        return storageFee;
    }

    function GetFileInfo(bytes memory fileHash)
        public
        view
        override
        NotEmptyFileHash(fileHash)
        returns (FileInfo memory)
    {
        FileInfo memory fileInfo = fileInfos[fileHash];
        return fileInfo;
    }

    function GetFileInfos(FileList memory _fileList)
        public
        view
        override
        returns (FileInfo[] memory)
    {
        require(_fileList.List.length > 0, "fileList is empty");
        require(_fileList.FileNum == _fileList.List.length, "fileNum is wrong");
        FileInfo[] memory _fileInfos = new FileInfo[](_fileList.List.length);
        for (uint256 i = 0; i < _fileList.List.length; i++) {
            bytes memory fileHash = _fileList.List[i];
            FileInfo memory fileInfo = fileInfos[fileHash];
            if (fileInfo.FileHash.length == 0) {
                revert FileNotExist(fileHash);
            }
            _fileInfos[i] = fileInfos[fileHash];
        }
        return _fileInfos;
    }

    function GetFileList(address walletAddr)
        public
        view
        override
        returns (FileList memory)
    {
        return fileList[walletAddr];
    }

    /****************************************************************************
     * WhiteList mamanagement ***************************************************
     */
    function GetWhiteList(bytes memory fileHash)
        public
        view
        override
        NotEmptyFileHash(fileHash)
        returns (WhiteList memory)
    {
        require(whiteList[fileHash].List.length > 0, "whiteList is empty");
        return whiteList[fileHash];
    }

    /****************************************************************************
     * Users=Space mamanagement ***************************************************
     */
    function GetUserSpace(address walletAddr)
        public
        view
        override
        returns (UserSpace memory)
    {
        require(userSpace[walletAddr].UpdateHeight > 0, "userSpace is empty");
        return userSpace[walletAddr];
    }

    /****************************************************************************
     * Sector mamanagement ***************************************************
     */
    function GetSectorInfo(SectorRef memory sectorRef)
        public
        view
        override
        returns (SectorInfo memory)
    {
        require(sectorRef.SectorId > 0, "sectorId must be greater than 0");
        return sectorInfos[sectorRef.NodeAddr][sectorRef.SectorId];
    }

    /****************************************************************************
     * Sector mamanagement ***************************************************
     */
    function GetPocProveList(uint32 height)
        public
        view
        override
        returns (PocProve[] memory)
    {
        require(height > 0, "Block number must lager than 0");
        return pocProveList[height];
    }
}
