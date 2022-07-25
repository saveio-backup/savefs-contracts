//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";
import "./FileExtra.sol";

/**
 * @title FileSystem
 * @dev FileSystem contract
 */
contract File is Initializable, IFile, IFsEvent {
    IConfig config;
    INode node;
    ISpace space;
    ISector sector;
    IProve prove;
    FileExtra fileExtra;

    uint64 DEFAULT_BLOCK_INTERVAL;
    uint64 DEFAULT_PROVE_PERIOD;
    uint64 IN_SECTOR_SIZE;

    function initialize(
        IConfig _config,
        INode _node,
        ISpace _space,
        ISector _sector,
        IProve _prove,
        FSConfig memory fsConfig,
        FileExtra _fileExtra
    ) public initializer {
        config = _config;
        node = _node;
        space = _space;
        sector = _sector;
        prove = _prove;
        DEFAULT_BLOCK_INTERVAL = fsConfig.DEFAULT_BLOCK_INTERVAL;
        DEFAULT_PROVE_PERIOD =
            fsConfig.DEFAULT_PROVE_PERIOD /
            DEFAULT_BLOCK_INTERVAL;
        IN_SECTOR_SIZE = fsConfig.IN_SECTOR_SIZE;
        fileExtra = _fileExtra;
    }

    function CalcFee(
        Setting memory setting,
        uint64 proveTime,
        uint64 copyNum,
        uint64 fileSize,
        uint64 duration
    ) public view virtual override returns (StorageFee memory) {
        return
            fileExtra.CalcFee(setting, proveTime, copyNum, fileSize, duration);
    }

    function CalcDepositFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public view virtual override returns (StorageFee memory) {
        return fileExtra.CalcDepositFee(uploadOption, setting, currentHeight);
    }

    function calcUploadFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public view returns (StorageFee memory) {
        return fileExtra.calcUploadFee(uploadOption, setting, currentHeight);
    }

    function GetUploadStorageFee(UploadOption memory uploadOption)
        public
        view
        virtual
        override
        returns (StorageFee memory)
    {
        require(uploadOption.FileSize > 0, "fileSize must be greater than 0");
        Setting memory setting = config.GetSetting();
        return calcUploadFee(uploadOption, setting, block.number);
    }

    function StoreFile(FileInfo memory fileInfo)
        public
        payable
        virtual
        override
    {
        require(
            fileExtra.GetFileInfo(fileInfo.FileHash).BlockHeight == 0,
            "file already exist"
        );
        require(fileInfo.ExpiredHeight > block.number, "file expired");
        Setting memory setting = config.GetSetting();
        if (fileInfo.ProveLevel_ == ProveLevel.HIGH) {
            fileInfo.ProveInterval = DEFAULT_PROVE_PERIOD;
        }
        if (fileInfo.ProveLevel_ == ProveLevel.MEDIUM) {
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
        StorageFee memory uploadFee = CalcDepositFee(
            uploadOption,
            setting,
            block.number
        );
        fileInfo.Deposit =
            uploadFee.TxnFee +
            uploadFee.SpaceFee +
            uploadFee.ValidationFee;
        fileInfo.ProveTimes = fileExtra.CalcProveTimesByUploadInfo(
            uploadOption,
            block.number
        );
        if (fileInfo.StorageType_ == StorageType.Normal) {
            UserSpace memory _userSpace = space.GetUserSpace(
                fileInfo.FileOwner
            );
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
        fileExtra.UpdateFileInfo(fileInfo);
        fileExtra.AddFileToFileList(fileInfo.FileOwner, fileInfo.FileHash);
        for (uint256 i = 0; i < fileInfo.PrimaryNodes.length; i++) {
            fileExtra.AddFileToPrimaryList(
                fileInfo.PrimaryNodes[i],
                fileInfo.FileHash
            );
        }
        for (uint256 i = 0; i < fileInfo.CandidateNodes.length; i++) {
            fileExtra.AddFileToCandidateList(
                fileInfo.CandidateNodes[i],
                fileInfo.FileHash
            );
        }
        ProveDetailMeta memory _proveDetailMeta;
        _proveDetailMeta.CopyNum = fileInfo.CopyNum;
        _proveDetailMeta.ProveDetailNum = 0;
        prove.UpdateProveDetailMeta(fileInfo.FileHash, _proveDetailMeta);
        emit StoreFileEvent(
            FsEvent.STORE_FILE,
            block.number,
            fileInfo.FileHash,
            fileInfo.RealFileSize,
            fileInfo.FileOwner,
            fileInfo.Deposit,
            fileInfo.IsPlotFile
        );
    }

    function FileReNew(FileReNewInfo memory fileReNewInfo)
        public
        payable
        virtual
        override
    {
        require(
            fileExtra.GetFileInfo(fileReNewInfo.FileHash).BlockHeight > 0,
            "file not exist"
        );
        require(
            fileExtra.GetFileInfo(fileReNewInfo.FileHash).StorageType_ ==
                StorageType.Professional,
            "file type error"
        );
        require(
            fileExtra.GetFileInfo(fileReNewInfo.FileHash).ExpiredHeight >
                block.number,
            "file expired"
        );
        Setting memory setting = config.GetSetting();
        FileInfo memory fileInfo = fileExtra.GetFileInfo(
            fileReNewInfo.FileHash
        );
        StorageFee memory totalRenew = CalcFee(
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
        fileExtra.UpdateFileInfo(fileInfo);
    }

    function GetFileInfo(bytes memory fileHash)
        public
        view
        virtual
        override
        returns (FileInfo memory)
    {
        return fileExtra.GetFileInfo(fileHash);
    }

    function GetFileInfos(bytes[] memory _fileList)
        public
        view
        virtual
        override
        returns (FileInfo[] memory)
    {
        return fileExtra.GetFileInfos(_fileList);
    }

    function GetFileList(address walletAddr)
        public
        view
        virtual
        override
        returns (bytes[] memory)
    {
        return fileExtra.GetFileList(walletAddr);
    }

    function UpdateFileInfo(FileInfo memory f) public payable virtual override {
        fileExtra.UpdateFileInfo(f);
    }

    function DeleteFileInfo(bytes memory fileHash) public {
        fileExtra.DeleteFileInfo(fileHash);
    }

    function deleteProveDetails(bytes memory fileHash) public {
        fileExtra.DeleteFileInfo(fileHash);
        prove.DeleteProveDetails(fileHash);
    }

    function UpdateFileList(address walletAddr, bytes[] memory list)
        public
        payable
        virtual
        override
    {
        fileExtra.UpdateFileList(walletAddr, list);
    }

    function UpdateFilesForRenew(
        bytes[] memory _fileList,
        Setting memory setting,
        uint256 newExpireHeight
    ) public payable virtual override returns (FileInfo[] memory, bool) {
        return
            fileExtra.UpdateFilesForRenew(_fileList, setting, newExpireHeight);
    }

    function GetUnSettledFileList(address walletAddr)
        public
        view
        virtual
        override
        returns (bytes[] memory)
    {
        return fileExtra.GetUnSettledFileList(walletAddr);
    }

    function AddFileToUnSettleList(address fileOwner, bytes memory fileHash)
        public
        payable
        virtual
        override
    {
        fileExtra.AddFileToUnSettleList(fileOwner, fileHash);
    }

    function CleanupForDeleteFile(
        FileInfo memory fileInfo,
        bool rmInfo,
        bool rmList
    ) public payable virtual override {
        bytes memory fileHash = fileInfo.FileHash;
        if (rmInfo) {
            DeleteFileInfo(fileHash);
            deleteProveDetails(fileHash);
            fileExtra.DelFileFromUnSettledList(fileInfo.FileOwner, fileHash);
        }
        if (rmList) {
            fileExtra.DelFileFromList(fileInfo.FileOwner, fileHash);
            for (uint256 i = 0; i < fileInfo.PrimaryNodes.length; i++) {
                fileExtra.DelFileFromPrimaryList(
                    fileInfo.PrimaryNodes[i],
                    fileHash
                );
            }
            for (uint256 i = 0; i < fileInfo.CandidateNodes.length; i++) {
                fileExtra.DelFileFromCandidateList(
                    fileInfo.CandidateNodes[i],
                    fileHash
                );
            }
        }
    }

    function DeleteExpiredFilesFromList(
        bytes[] memory _fileList,
        address walletAddr,
        StorageType[] memory storageType
    )
        public
        virtual
        override
        returns (
            bytes[] memory,
            uint64,
            bool
        )
    {
        bytes[] memory deletedFiles = new bytes[](_fileList.length);
        uint64 n = 0;
        uint64 amount;
        bool success;
        Setting memory setting = config.GetSetting();
        for (uint256 i = 0; i < _fileList.length; i++) {
            FileInfo memory fileInfo = GetFileInfo(_fileList[i]);
            bool exist = false;
            for (uint256 j = 0; j < storageType.length; j++) {
                if (fileInfo.StorageType_ == storageType[j]) {
                    exist = true;
                    break;
                }
            }
            if (!exist) {
                continue;
            }
            if (
                fileInfo.ExpiredHeight + setting.DefaultProvePeriod >
                block.number
            ) {
                continue;
            }
            amount += fileInfo.Deposit;
            CleanupForDeleteFile(fileInfo, true, true);
            deletedFiles[n] = _fileList[i];
            n++;
        }
        if (amount > 0) {
            payable(walletAddr).transfer(amount);
        }
        return (deletedFiles, amount, success);
    }

    function DeleteUnsettledFiles(address walletAddr) public virtual override {
        bytes[] memory unsettledList = GetUnSettledFileList(walletAddr);
        StorageType[] memory sType = new StorageType[](2);
        sType[0] = StorageType.Normal;
        sType[1] = StorageType.Professional;
        bytes[] memory deletedFiles;
        uint64 amount;
        bool success;
        (deletedFiles, amount, success) = DeleteExpiredFilesFromList(
            unsettledList,
            walletAddr,
            sType
        );
        for (uint256 i = 0; i < deletedFiles.length; i++) {
            bytes[] memory list = GetUnSettledFileList(walletAddr);
            for (uint256 j = 0; j < list.length; j++) {
                if (keccak256(list[j]) == keccak256(deletedFiles[i])) {
                    delete list[j];
                }
            }
            fileExtra.UpdateUnSettleList(walletAddr, list);
        }
    }

    function GetUnProvePrimaryFiles(address walletAddr)
        public
        view
        virtual
        override
        returns (bytes[] memory)
    {
        bytes[] memory list = fileExtra.GetPrimaryFiles(walletAddr);
        bytes[] memory unProvePrimaryFiles = new bytes[](list.length);
        uint64 n = 0;
        for (uint256 i = 0; i < list.length; i++) {
            ProveDetail[] memory details = prove.GetProveDetailList(list[i]);
            bool isProve = false;
            for (uint256 j = 0; j < details.length; j++) {
                if (details[j].WalletAddr == walletAddr) {
                    isProve = true;
                    break;
                }
            }
            if (!isProve) {
                continue;
            }
            unProvePrimaryFiles[n] = list[i];
            n++;
        }
        return unProvePrimaryFiles;
    }

    function GetUnProveCandidateFiles(address walletAddr)
        public
        view
        virtual
        override
        returns (bytes[] memory)
    {
        bytes[] memory list = fileExtra.GetCandidateFiles(walletAddr);
        bytes[] memory unProveCandidateFiles = new bytes[](list.length);
        uint64 n = 0;
        for (uint256 i = 0; i < list.length; i++) {
            ProveDetail[] memory details = prove.GetProveDetailList(list[i]);
            bool isProve = false;
            for (uint256 j = 0; j < details.length; j++) {
                if (details[j].WalletAddr == walletAddr) {
                    isProve = true;
                    break;
                }
            }
            if (!isProve) {
                continue;
            }
            unProveCandidateFiles[n] = list[i];
            n++;
        }
        return unProveCandidateFiles;
    }

    function ChangeFilePrivilege(PriChange memory priChange)
        public
        virtual
        override
    {
        FileInfo memory fileInfo = GetFileInfo(priChange.fileHash);
        fileInfo.Privilege = priChange.privilege;
        UpdateFileInfo(fileInfo);
    }

    function ChangeFileOwner(OwnerChange memory ownerChange)
        public
        virtual
        override
    {
        FileInfo memory fileInfo = GetFileInfo(ownerChange.FileHash);
        require(
            fileInfo.FileOwner == ownerChange.CurOwner,
            "Current owner must be the owner of the file"
        );
        fileInfo.FileOwner = ownerChange.NewOwner;
        UpdateFileInfo(fileInfo);
    }

    function deleteFilesInner(FileInfo[] memory files) public {
        if (files.length == 0) {
            return;
        }
        uint64 refundAmount = 0;
        address fileOwner = files[0].FileOwner;
        Setting memory setting = config.GetSetting();
        for (uint256 i = 0; i < files.length; i++) {
            FileInfo memory fileInfo = files[i];
            if (fileInfo.FileOwner != fileOwner) {
                revert DifferenceFileOwner();
            }
        }
        for (uint256 i = 0; i < files.length; i++) {
            FileInfo memory fileInfo = files[i];
            SectorRef[] memory sectorRefs = fileExtra.GetFileSectorRefs(
                fileInfo.FileHash
            );
            for (uint256 j = 0; j < sectorRefs.length; j++) {
                SectorInfo memory sectorInfo = sector.GetSectorInfo(
                    sectorRefs[j]
                );
                sector.DeleteFileFromSector(sectorInfo, fileInfo);
            }
            if (fileInfo.Deposit == 0) {
                CleanupForDeleteFile(fileInfo, true, true);
                continue;
            }
            ProveDetail[] memory details = prove.GetProveDetailList(
                fileInfo.FileHash
            );
            uint64 restProfit = fileInfo.Deposit;
            uint64 fileSize = fileInfo.FileBlockNum * fileInfo.FileBlockSize;
            uint64 singleProveProfit = fileExtra.calcSingleValidFeeForFile(
                setting,
                fileSize
            );
            for (uint256 j = 0; j < details.length; j++) {
                NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
                    details[j].WalletAddr
                );
                uint64 validProfit = (details[j].ProveTimes - 1) *
                    singleProveProfit;
                uint64 storageProfit = fileExtra.calcStorageFeeForOneNode(
                    setting,
                    fileSize,
                    uint64(block.number - fileInfo.BlockHeight)
                );
                uint64 totalProfit = validProfit + storageProfit;
                if (totalProfit > restProfit) {
                    revert InvalidProfit();
                }
                nodeInfo.Profit += totalProfit;
                node.UpdateNodeInfo(nodeInfo);
                restProfit -= totalProfit;
            }
            if (fileInfo.StorageType_ == StorageType.Normal) {
                refundAmount += restProfit;
            } else if (fileInfo.StorageType_ == StorageType.Professional) {
                UserSpace memory userSpace = space.GetUserSpace(
                    fileInfo.FileOwner
                );
                if (
                    userSpace.Used >=
                    fileInfo.FileBlockNum * fileInfo.FileBlockSize
                ) {
                    userSpace.Balance += restProfit;
                    userSpace.Remain +=
                        fileInfo.FileBlockNum *
                        fileInfo.FileBlockSize;
                    userSpace.Used -=
                        fileInfo.FileBlockNum *
                        fileInfo.FileBlockSize;
                } else {
                    revert OpError(1);
                }
                space.UpdateUserSpace(fileInfo.FileOwner, userSpace);
            }
            CleanupForDeleteFile(fileInfo, true, true);
        }
        if (refundAmount == 0) {
            return;
        }
        payable(fileOwner).transfer(refundAmount);
    }

    function DeleteFile(bytes memory fileHash) public virtual override {
        FileInfo memory fileInfo = GetFileInfo(fileHash);
        FileInfo[] memory files = new FileInfo[](1);
        files[0] = fileInfo;
        deleteFilesInner(files);
        emit DeleteFileEvent(
            FsEvent.DELETE_FILE,
            block.number,
            fileHash,
            fileInfo.FileOwner
        );
    }

    function DeleteFiles(bytes[] memory fileHashs) public virtual override {
        address fileOwner;
        FileInfo[] memory files = new FileInfo[](fileHashs.length);
        for (uint256 i = 0; i < fileHashs.length; i++) {
            FileInfo memory fileInfo = GetFileInfo(fileHashs[i]);
            fileOwner = fileInfo.FileOwner;
            files[i] = fileInfo;
        }
        deleteFilesInner(files);
        emit DeleteFilesEvent(
            FsEvent.DELETE_FILE,
            block.number,
            fileHashs,
            fileOwner
        );
    }
}
