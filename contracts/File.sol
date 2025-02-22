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
        FileInfo memory tmpFile = GetFileInfo(fileInfo.FileHash);
        if (tmpFile.BlockHeight > 0) {
            emit FsError("StoreFile", "file already exist");
            return;
        }
        if (fileInfo.ExpiredHeight < block.number) {
            emit FsError(
                "StoreFile",
                "expiredHeight must be greater than current height"
            );
            return;
        }
        string memory err = fileExtra.StoreFile{value: msg.value}(
            fileInfo,
            config,
            space,
            prove,
            DEFAULT_PROVE_PERIOD
        );
        if (bytes(err).length > 0) {
            emit FsError("StoreFile", err);
            return;
        }
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
        Setting memory setting = config.GetSetting();
        string memory err = fileExtra.FileReNew{value: msg.value}(
            setting,
            fileReNewInfo
        );
        if (bytes(err).length > 0) {
            emit FsError("FileReNew", err);
            return;
        }
    }

    function GetFileInfo(bytes memory fileHash)
        public
        view
        virtual
        override
        returns (FileInfo memory)
    {
        FileInfo memory fileInfo = fileExtra.GetFileInfo(fileHash);
        return fileInfo;
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

    function DeleteProveDetails(bytes memory fileHash) public {
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
        if (rmInfo) {
            DeleteFileInfo(fileInfo.FileHash);
            DeleteProveDetails(fileInfo.FileHash);
            fileExtra.DelFileFromUnSettledList(
                fileInfo.FileOwner,
                fileInfo.FileHash
            );
        }
        if (rmList) {
            fileExtra.DeleteFileFromList(fileInfo);
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

    function DeleteUnsettledFiles(address walletAddr)
        public
        payable
        virtual
        override
    {
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
        return GetUnProveFiles(list, walletAddr);
    }

    function GetUnProveCandidateFiles(address walletAddr)
        public
        view
        virtual
        override
        returns (bytes[] memory)
    {
        bytes[] memory list = fileExtra.GetCandidateFiles(walletAddr);
        return GetUnProveFiles(list, walletAddr);
    }

    function GetUnProveFiles(bytes[] memory list, address walletAddr)
        private
        view
        returns (bytes[] memory)
    {
        bytes[] memory files = new bytes[](list.length);
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
            files[n] = list[i];
            n++;
        }
        return files;
    }

    function ChangeFilePrivilege(PriChange memory priChange)
        public
        virtual
        override
    {
        fileExtra.ChangeFilePrivilege(priChange);
    }

    function ChangeFileOwner(OwnerChange memory ownerChange)
        public
        virtual
        override
    {
        fileExtra.ChangeFileOwner(ownerChange);
    }

    function deleteFilesInner(FileInfo[] memory files)
        public
        payable
        returns (string memory error)
    {
        if (files.length == 0) {
            return "no files";
        }
        uint64 refundAmount = 0;
        address fileOwner = files[0].FileOwner;
        Setting memory setting = config.GetSetting();
        for (uint256 i = 0; i < files.length; i++) {
            FileInfo memory info = files[i];
            if (info.FileHash.length == 0) {
                continue;
            }
            if (info.FileOwner != fileOwner) {
                return "file owner is different";
            }
            SectorRef[] memory sectorRefs = fileExtra.GetFileSectorRefs(
                info.FileHash
            );
            for (uint256 j = 0; j < sectorRefs.length; j++) {
                SectorInfo memory sectorInfo = sector.GetSectorInfo(
                    sectorRefs[j]
                );
                sector.DeleteFileFromSector(sectorInfo, info);
            }
            if (info.Deposit == 0) {
                CleanupForDeleteFile(info, true, true);
                continue;
            }
            ProveDetail[] memory details = prove.GetProveDetailList(
                info.FileHash
            );
            uint64 profit = info.Deposit;
            uint64 size = info.FileBlockNum * info.FileBlockSize;
            uint64 fee = fileExtra.calcSingleValidFeeForFile(setting, size);
            for (uint256 j = 0; j < details.length; j++) {
                NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
                    details[j].WalletAddr
                );
                uint64 validProfit = (details[j].ProveTimes - 1) * fee;
                uint64 storageProfit = fileExtra.calcStorageFeeForOneNode(
                    setting,
                    size,
                    uint64(block.number - info.BlockHeight)
                );
                uint64 totalProfit = validProfit + storageProfit;
                if (totalProfit > profit) {
                    return "profit is invalid";
                }
                nodeInfo.Profit += totalProfit;
                node.UpdateNodeInfo(nodeInfo);
                profit -= totalProfit;
            }
            if (info.StorageType_ == StorageType.Normal) {
                UserSpace memory us = space.GetUserSpace(info.FileOwner);
                if (us.Used >= info.FileBlockNum * info.FileBlockSize) {
                    us.Balance += profit;
                    us.Remain += info.FileBlockNum * info.FileBlockSize;
                    us.Used -= info.FileBlockNum * info.FileBlockSize;
                } else {
                    return "user space is invalid";
                }
                space.UpdateUserSpace(info.FileOwner, us);
            }
            if (info.StorageType_ == StorageType.Professional) {
                refundAmount += profit;
            }
            CleanupForDeleteFile(info, true, true);
        }
        if (refundAmount > 0) {
            bool res = payable(fileOwner).send(refundAmount);
            if (!res) {
                return "refund failed";
            }
        }
        return "";
    }

    function DeleteFile(bytes memory fileHash) public payable virtual override {
        FileInfo memory fileInfo = GetFileInfo(fileHash);
        FileInfo[] memory files = new FileInfo[](1);
        files[0] = fileInfo;
        string memory error = deleteFilesInner(files);
        if (bytes(error).length > 0) {
            emit FsError("DeleteFile", error);
            return;
        } else {
            emit DeleteFileEvent(
                FsEvent.DELETE_FILE,
                block.number,
                fileHash,
                fileInfo.FileOwner
            );
        }
    }

    function DeleteFiles(bytes[] memory fileHashs)
        public
        payable
        virtual
        override
    {
        address fileOwner;
        FileInfo[] memory files = new FileInfo[](fileHashs.length);
        for (uint256 i = 0; i < fileHashs.length; i++) {
            FileInfo memory fileInfo = GetFileInfo(fileHashs[i]);
            fileOwner = fileInfo.FileOwner;
            files[i] = fileInfo;
        }
        string memory error = deleteFilesInner(files);
        if (bytes(error).length > 0) {
            emit FsError("DeleteFiles", error);
            return;
        } else {
            emit DeleteFilesEvent(
                FsEvent.DELETE_FILE,
                block.number,
                fileHashs,
                fileOwner
            );
        }
    }

    function AddFileSectorRef(bytes memory fileHash, SectorRef memory ref)
        public
        payable
        virtual
        override
    {
        fileExtra.AddFileSectorRef(fileHash, ref);
    }
}
