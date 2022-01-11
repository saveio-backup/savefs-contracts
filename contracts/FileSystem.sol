//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/** enum ********************** */
enum FsEvent {
    STORE_FILE,
    DELETE_FILE,
    DELETE_FILES,
    SET_USER_SPACE,
    REG_NODE,
    UN_REG_NODE,
    PROVE_FILE,
    FILE_PDP_SUCCESS,
    CREATE_SECTOR,
    DELETE_SECTOR
}

enum WHileListOp {
    ADD,
    DEL,
    ADD_COV,
    DEL_ALL,
    UPDATE
}

enum ProveLevel {
    HIGH,
    MEDIEUM,
    LOW
}

enum StorageType {
    Normal,
    Professional
}

/** setting ********************** */
struct Setting {
    uint64 GasPrice;
    uint64 GasPerGBPerBlock;
    uint64 GasPerKBPerBlock;
    uint64 GasForChallenge;
    uint64 MaxProveBlockNum;
    uint64 MinVolume;
    uint64 DefaultProvePeriod;
    uint64 DefaultProveLevel;
    uint64 DefaultCopyNum;
}

/** node info ******************** */
struct NodeInfo {
    uint64 Pledge;
    uint64 Profit;
    uint64 Volume;
    uint64 RestVol;
    uint64 ServiceTime;
    address WalletAddr;
    address NodeAddr;
}

struct NodeList {
    uint64 AddrNum;
    address[] AddrList;
}

/** userspace ******************** */
struct UserSpace {
    uint64 Used;
    uint64 Remain;
    uint64 ExpireHeight;
    uint64 Balance;
    uint64 UpdateHeight;
}

/** file ********************** */
struct StorageFee {
    uint64 TxnFee;
    uint64 SpaceFee;
    uint64 ValidationFee;
}

struct Role {
    address Addr;
    uint64 BaseHeight;
    uint64 ExpireHeight;
}

struct WhiteListOp {
    bytes FileHash;
    WHileListOp Op;
    WhiteList List;
}

struct WhiteList {
    uint64 Num;
    Role[] List;
}

struct UploadOption {
    bytes FileDesc;
    uint64 FileSize;
    uint64 ProveInterval;
    uint64 ProveLevel;
    uint64 ExpiredHeight;
    uint64 Privilege;
    uint64 CopyNum;
    bool Encrypt;
    bytes EncryptPassword;
    bool RegisterDNS;
    bool BindDNS;
    bytes DnsURL;
    WhiteList WhiteList_;
    bool Share;
    StorageType StorageType_;
}

struct SectorRef {
    address NodeAddr;
    uint64 SectorId;
}

struct PlotInfo {
    uint64 NumberID;
    uint64 StartNonce;
    uint64 Nonces;
}

struct FileInfo {
    bytes FileHash;
    address FileOwner;
    bytes FileDesc;
    uint64 Privilege;
    uint64 FileBlockNum;
    uint64 FileBlockSize;
    uint64 ProveInterval;
    uint64 ProveTimes;
    uint64 ExpiredHeight;
    uint64 CopyNum;
    uint64 Deposit;
    bytes FileProveParam;
    uint64 ProveBlockNum;
    uint256 BlockHeight;
    bool ValidFlag;
    StorageType StorageType_;
    uint64 RealFileSize;
    NodeList PrimaryNodes;
    NodeList CandidateNodes;
    ProveLevel ProveLevel_;
    // SectorRef[] SectorRefs; // TODO
    bool IsPlotFile;
    PlotInfo PlotInfo_;
}

struct FileList {
    uint64 FileNum;
    bytes[] List;
}

struct SectorInfo {
    address NodeAddr;
    uint64 SectorID;
    uint64 Size;
    uint64 Used;
    ProveLevel ProveLevel_;
    uint64 FirstProveHeight;
    uint64 NextProveHeight;
    uint64 TotalBlockNum;
    uint64 FileNum;
    uint64 GroupNum;
    bool IsPlots;
    FileList FileList_;
}

struct SectorInfos {
    uint64 SectorCount;
    uint64[] SectorIds;
}

struct FileReNewInfo {
    bytes FileHash;
    address FromAddr;
    uint64 ReNewTimes;
}

/** prove ********************** */
struct PocProve {
    address Miner;
    uint32 Height;
    uint64 PlotSize;
}

struct ProveDetail {
    address NodeAddr;
    address WalletAddr;
    uint64 ProveTimes;
    uint64 BlockHeight;
    bool Finished;
}

struct ProveDetails {
    uint64 CopyNum;
    uint64 ProveDetailNum;
    // ProveDetail[] ProveDetails; // TODO
}

/**
 * @title IFileSystem
 * @dev IFileSystem contract
 */
abstract contract IFileSystem {
    function GetSetting() public pure virtual returns (Setting memory);

    /** node ****************************************************** */
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

    /** file ****************************************************** */
    function GetUploadStorageFee(UploadOption memory uploadOption)
        public
        view
        virtual
        returns (StorageFee memory);

    function StoreFile(FileInfo memory fileInfo) public payable virtual;

    function FileReNew(FileReNewInfo memory fileReNewInfo)
        public
        payable
        virtual;

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

    /** white list ****************************************************** */
    function GetWhiteList(bytes memory fileHash)
        public
        view
        virtual
        returns (WhiteList memory);

    /** userspace ****************************************************** */
    function GetUserSpace(address walletAddr)
        public
        view
        virtual
        returns (UserSpace memory);

    /** sector ****************************************************** */
    function GetSectorInfo(SectorRef memory sectorRef)
        public
        view
        virtual
        returns (SectorInfo memory);

    /** prove ****************************************************** */
    function GetPocProveList(uint32 height)
        public
        view
        virtual
        returns (PocProve[] memory);
}

/**
 * @title FileSystem
 * @dev FileSystem contract
 */
contract FileSystem is Initializable, IFileSystem {
    /************************************************************************
     * Constant define ******************************************************
     */
    uint64 DEFAULT_BLOCK_INTERVAL = 5;
    uint64 DEFAULT_PROVE_PERIOD = (3600 * 24) / DEFAULT_BLOCK_INTERVAL;

    /************************************************************************
     * Field define ******************************************************
     */
    mapping(address => NodeInfo) nodesInfo; // walletAddr => NodeInfo
    NodeList nodeList; // nodeAddr list
    mapping(address => mapping(uint64 => SectorInfo)) sectorInfos; // nodeAddr => (sectorId => SectorInfo)
    mapping(bytes => FileInfo) fileInfos; // fileHash => FileInfo
    mapping(bytes => ProveDetails) proveDetails; // fileHash => ProveDetails
    mapping(address => FileList) fileList; // walletAddr => filelist
    mapping(address => FileList) primaryFileList; // walletAddr => filelist
    mapping(address => FileList) candidateFileList; // walletAddr => filelist
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
    error UserspaceInsufficientBalance(uint256 got, uint256 want);
    error UserspaceInsufficientSpace(uint256 got, uint256 want);
    error UserspaceWrongExpiredHeight(uint256 got, uint256 want);
    error NotEnoughTransfer(uint256 got, uint256 want);
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
        setting.MaxProveBlockNum = 32;
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
    function calcProveTimesByUploadInfo(
        UploadOption memory uploadOption,
        uint256 beginHeight
    ) public pure returns (uint64) {
        return
            (uploadOption.ExpiredHeight - uint64(beginHeight)) /
            uploadOption.ProveInterval +
            1;
    }

    function calcSingleValidFeeForFile(Setting memory setting, uint64 fileSize)
        public
        pure
        returns (uint64)
    {
        return (setting.GasForChallenge * fileSize) / 1024000;
    }

    function calcValidFeeForOneNode(
        Setting memory setting,
        uint64 proveTime,
        uint64 fileSize
    ) public pure returns (uint64) {
        return proveTime * calcSingleValidFeeForFile(setting, fileSize);
    }

    function calcValidFee(
        Setting memory setting,
        uint64 proveTime,
        uint64 copyNum,
        uint64 fileSize
    ) public pure returns (uint64) {
        return
            (copyNum + 1) *
            calcValidFeeForOneNode(setting, proveTime, fileSize);
    }

    function calcStorageFeeForOneNode(
        Setting memory setting,
        uint64 fileSize,
        uint64 duration
    ) public pure returns (uint64) {
        return (setting.GasPerGBPerBlock * fileSize * duration) / 1024000;
    }

    function calcStorageFee(
        Setting memory setting,
        uint64 copyNum,
        uint64 fileSize,
        uint64 duration
    ) public pure returns (uint64) {
        return
            (copyNum + 1) *
            calcStorageFeeForOneNode(setting, fileSize, duration);
    }

    function calcFee(
        Setting memory setting,
        uint64 proveTime,
        uint64 copyNum,
        uint64 fileSize,
        uint64 duration
    ) public pure returns (StorageFee memory) {
        StorageFee memory fee;
        uint64 validFee = calcValidFee(setting, proveTime, copyNum, fileSize);
        uint64 storageFee = calcStorageFee(
            setting,
            copyNum,
            fileSize,
            duration
        );
        fee.ValidationFee = validFee;
        fee.SpaceFee = storageFee;
        return fee;
    }

    function calcDepositFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public pure returns (StorageFee memory) {
        uint64 proveTime = calcProveTimesByUploadInfo(
            uploadOption,
            currentHeight
        );
        StorageFee memory fee = calcFee(
            setting,
            proveTime,
            uploadOption.CopyNum,
            uploadOption.FileSize,
            uploadOption.ExpiredHeight - uint64(currentHeight)
        );
        return fee;
    }

    function calcUploadFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public pure returns (StorageFee memory) {
        uint64 fee;
        uint64 txGas = 10000000;
        if (uploadOption.WhiteList_.Num > 0) {
            fee = txGas * 4;
        } else {
            fee = txGas * 3;
        }
        StorageFee memory sf;
        sf.TxnFee = fee;
        if (uploadOption.StorageType_ == StorageType.Normal) {
            return sf;
        }
        StorageFee memory depositFee = calcDepositFee(
            uploadOption,
            setting,
            currentHeight
        );
        sf.ValidationFee = depositFee.ValidationFee;
        sf.SpaceFee = depositFee.SpaceFee;
        return sf;
    }

    function GetUploadStorageFee(UploadOption memory uploadOption)
        public
        view
        override
        returns (StorageFee memory)
    {
        require(uploadOption.FileSize > 0, "fileSize must be greater than 0");
        Setting memory setting = GetSetting();
        return calcUploadFee(uploadOption, setting, block.number);
    }

    function StoreFile(FileInfo memory fileInfo) public payable override {
        require(
            fileInfos[fileInfo.FileHash].BlockHeight == 0,
            "file already exist"
        );
        require(fileInfo.ExpiredHeight > block.number, "file expired");
        Setting memory setting = GetSetting();
        if (fileInfo.ProveLevel_ == ProveLevel.HIGH) {
            fileInfo.ProveInterval = DEFAULT_PROVE_PERIOD;
        }
        if (fileInfo.ProveLevel_ == ProveLevel.MEDIEUM) {
            fileInfo.ProveInterval = DEFAULT_PROVE_PERIOD;
        }
        if (fileInfo.ProveLevel_ == ProveLevel.LOW) {
            fileInfo.ProveInterval = DEFAULT_PROVE_PERIOD;
        }
        fileInfo.ValidFlag = true;
        UploadOption memory uploadOption;
        uploadOption.ExpiredHeight = fileInfo.ExpiredHeight;
        uploadOption.ProveInterval = fileInfo.ProveInterval;
        uploadOption.CopyNum = fileInfo.CopyNum;
        uploadOption.FileSize = fileInfo.FileBlockSize * fileInfo.FileBlockNum;
        StorageFee memory uploadFee = calcDepositFee(
            uploadOption,
            setting,
            block.number
        );
        fileInfo.Deposit =
            uploadFee.TxnFee +
            uploadFee.SpaceFee +
            uploadFee.ValidationFee;
        fileInfo.ProveTimes = calcProveTimesByUploadInfo(
            uploadOption,
            block.number
        );
        if (fileInfo.StorageType_ == StorageType.Normal) {
            UserSpace memory _userSpace = userSpace[fileInfo.FileOwner];
            if (_userSpace.Balance < fileInfo.Deposit) {
                revert UserspaceInsufficientBalance(
                    _userSpace.Balance,
                    fileInfo.Deposit
                );
            }
            if (
                _userSpace.Remain <
                fileInfo.FileBlockSize * fileInfo.FileBlockNum
            ) {
                revert UserspaceInsufficientSpace(
                    _userSpace.Remain,
                    fileInfo.FileBlockSize
                );
            }
            if (_userSpace.ExpireHeight < fileInfo.ExpiredHeight) {
                revert UserspaceWrongExpiredHeight(
                    _userSpace.ExpireHeight,
                    fileInfo.ExpiredHeight
                );
            }
            _userSpace.Balance -= fileInfo.Deposit;
            _userSpace.Remain -= fileInfo.FileBlockNum * fileInfo.FileBlockSize;
            _userSpace.Used += fileInfo.FileBlockNum * fileInfo.FileBlockSize;
        } else {
            require(msg.value >= fileInfo.Deposit, "insufficient deposit");
            fileInfo.StorageType_ = StorageType.Professional;
        }
        fileInfo.ProveBlockNum = setting.MaxProveBlockNum;
        fileInfo.BlockHeight = block.number;
        // store file
        fileInfos[fileInfo.FileHash] = fileInfo;
        FileList storage f = fileList[fileInfo.FileOwner];
        f.FileNum++;
        // TODO should unique file hash
        f.List.push(fileInfo.FileHash);
        for (uint256 i = 0; i < fileInfo.PrimaryNodes.AddrList.length; i++) {
            FileList storage p = fileList[fileInfo.PrimaryNodes.AddrList[i]];
            p.List.push(fileInfo.FileHash);
        }
        for (uint256 i = 0; i < fileInfo.CandidateNodes.AddrList.length; i++) {
            FileList storage p = fileList[fileInfo.CandidateNodes.AddrList[i]];
            p.List.push(fileInfo.FileHash);
        }
        ProveDetails memory _proveDetails;
        _proveDetails.CopyNum = fileInfo.CopyNum;
        _proveDetails.ProveDetailNum = 0;
        proveDetails[fileInfo.FileHash] = _proveDetails;
    }

    function FileReNew(FileReNewInfo memory fileReNewInfo)
        public
        payable
        override
    {
        require(
            fileInfos[fileReNewInfo.FileHash].BlockHeight > 0,
            "file not exist"
        );
        require(
            fileInfos[fileReNewInfo.FileHash].StorageType_ ==
                StorageType.Professional,
            "file type error"
        );
        require(
            fileInfos[fileReNewInfo.FileHash].ExpiredHeight > block.number,
            "file expired"
        );
        Setting memory setting = GetSetting();
        FileInfo memory fileInfo = fileInfos[fileReNewInfo.FileHash];
        StorageFee memory totalRenew = calcFee(
            setting,
            fileReNewInfo.ReNewTimes,
            fileInfo.CopyNum,
            fileInfo.FileBlockNum * fileInfo.FileBlockSize,
            fileReNewInfo.ReNewTimes * fileInfo.ProveInterval
        );
        uint64 reNewFee = totalRenew.ValidationFee + totalRenew.SpaceFee;
        if (msg.value < reNewFee) {
            revert NotEnoughTransfer(msg.value, reNewFee);
        }
        fileInfo.ProveTimes += fileReNewInfo.ReNewTimes;
        fileInfo.Deposit += reNewFee;
        fileInfo.ExpiredHeight +=
            fileInfo.ProveInterval *
            fileReNewInfo.ReNewTimes;
        fileInfos[fileReNewInfo.FileHash] = fileInfo;
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
