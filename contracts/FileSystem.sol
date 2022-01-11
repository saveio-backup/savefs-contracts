//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Library.sol";
import "./Type.sol";
import "./Node.sol";

/**
 * @title FileSystem
 * @dev FileSystem contract
 */
contract FileSystem is Initializable {
    Config config;
    Node node;

    uint64 DEFAULT_BLOCK_INTERVAL = 5;
    uint64 DEFAULT_PROVE_PERIOD = (3600 * 24) / DEFAULT_BLOCK_INTERVAL;
    uint64 IN_SECTOR_SIZE = 1000 * 1000;

    mapping(bytes => FileInfo) fileInfos; // fileHash => FileInfo
    mapping(address => FileList) fileList; // walletAddr => filelist
    mapping(address => FileList) primaryFileList; // walletAddr => filelist
    mapping(address => FileList) candidateFileList; // walletAddr => filelist
    mapping(address => mapping(uint64 => SectorInfo)) sectorInfos; // nodeAddr => (sectorId => SectorInfo)
    mapping(bytes => ProveDetails) proveDetails; // fileHash => ProveDetails
    mapping(bytes => WhiteList) whiteList; // fileHash => whileList
    mapping(address => UserSpace) userSpace; // walletAddr => UserSpace
    mapping(uint32 => PocProve[]) pocProveList; // blockNumber => PocProve

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

    error FileNotExist(bytes);
    error UserspaceInsufficientBalance(uint256 got, uint256 want);
    error UserspaceInsufficientSpace(uint256 got, uint256 want);
    error UserspaceWrongExpiredHeight(uint256 got, uint256 want);
    error NotEnoughTransfer(uint256 got, uint256 want);

    modifier NotEmptyFileHash(bytes memory fileHash) {
        require(fileHash.length > 0, "fileHash must be empty");
        _;
    }

    function initialize() public initializer {
        console.log("initializer");
    }

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
        returns (StorageFee memory)
    {
        require(uploadOption.FileSize > 0, "fileSize must be greater than 0");
        Setting memory setting = config.GetSetting();
        return calcUploadFee(uploadOption, setting, block.number);
    }

    function StoreFile(FileInfo memory fileInfo) public payable {
        require(
            fileInfos[fileInfo.FileHash].BlockHeight == 0,
            "file already exist"
        );
        require(fileInfo.ExpiredHeight > block.number, "file expired");
        Setting memory setting = config.GetSetting();
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

    function FileReNew(FileReNewInfo memory fileReNewInfo) public payable {
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
        Setting memory setting = config.GetSetting();
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
        NotEmptyFileHash(fileHash)
        returns (FileInfo memory)
    {
        FileInfo memory fileInfo = fileInfos[fileHash];
        return fileInfo;
    }

    function GetFileInfos(FileList memory _fileList)
        public
        view
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
        returns (UserSpace memory)
    {
        require(userSpace[walletAddr].UpdateHeight > 0, "userSpace is empty");
        return userSpace[walletAddr];
    }

    /****************************************************************************
     * Sector mamanagement ***************************************************
     */
    function getSectorTotalSizeForNode(address nodeAddr)
        public
        pure
        returns (uint64)
    {
        // uint64 totalSize = 0;
        // mapping(uint64 => SectorInfo) memory nodeSectorInfos = sectorInfos[nodeAddr];
    }

    function CreateSector(SectorInfo memory sectorInfo) public view {
        require(sectorInfo.SectorID > 0, "sectorId is wrong");
        require(sectorInfo.Size > 0, "sector size is wrong");
        require(
            node.IsNodeRegisted(sectorInfo.NodeAddr),
            "node not registered"
        );
    }

    function GetSectorInfo(SectorRef memory sectorRef)
        public
        view
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
        returns (PocProve[] memory)
    {
        require(height > 0, "Block number must lager than 0");
        return pocProveList[height];
    }
}
