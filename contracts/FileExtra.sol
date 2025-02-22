//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

/**
 * Slice contract from file.sol
 */
contract FileExtra is IFsEvent {
    mapping(bytes => FileInfo) fileInfos; // fileHash => FileInfo
    mapping(bytes => SectorRef[]) fileSectorRefs; // fileHash => SectorRef[]

    mapping(address => bytes[]) fileList; // walletAddr => bytes[]
    mapping(address => bytes[]) primaryFileList; // walletAddr => bytes[]
    mapping(address => bytes[]) candidateFileList; // walletAddr => bytes[]
    mapping(address => bytes[]) unSettledFileList; // walletAddr => bytes[]

    function StoreFile(
        FileInfo memory fileInfo,
        IConfig config,
        ISpace space,
        IProve prove,
        uint64 DEFAULT_PROVE_PERIOD
    ) public payable returns (string memory) {
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
        UploadOption memory option;
        option.ExpiredHeight = fileInfo.ExpiredHeight;
        option.ProveInterval = fileInfo.ProveInterval;
        option.CopyNum = fileInfo.CopyNum;
        option.FileSize = fileInfo.FileBlockSize * fileInfo.FileBlockNum;
        StorageFee memory uploadFee = CalcDepositFee(
            option,
            setting,
            block.number
        );
        fileInfo.Deposit =
            uploadFee.TxnFee +
            uploadFee.SpaceFee +
            uploadFee.ValidationFee;
        fileInfo.ProveTimes = CalcProveTimesByUploadInfo(option, block.number);
        if (fileInfo.StorageType_ == StorageType.Normal) {
            UserSpace memory us = space.GetUserSpace(fileInfo.FileOwner);
            if (us.Balance < fileInfo.Deposit) {
                return "userspace Insufficient balance";
            }
            if (us.Remain < fileInfo.FileBlockSize * fileInfo.FileBlockNum) {
                return "userspace Insufficient space";
            }
            if (us.ExpireHeight < fileInfo.ExpiredHeight) {
                return "userspace ExpiredHeight is wrong";
            }
            us.Balance -= fileInfo.Deposit;
            us.Remain -= fileInfo.FileBlockNum * fileInfo.FileBlockSize;
            us.Used += fileInfo.FileBlockNum * fileInfo.FileBlockSize;
        } else {
            if (msg.value < fileInfo.Deposit) {
                return "insufficient deposit";
            }
            fileInfo.StorageType_ = StorageType.Professional;
        }
        fileInfo.ProveBlockNum = setting.MaxProveBlockNum;
        fileInfo.BlockHeight = block.number;
        // store file
        SaveFile(fileInfo);
        ProveDetailMeta memory meta;
        meta.CopyNum = fileInfo.CopyNum;
        meta.ProveDetailNum = 0;
        prove.UpdateProveDetailMeta(fileInfo.FileHash, meta);
        return "";
    }

    function GetFileInfo(bytes memory fileHash)
        public
        view
        returns (FileInfo memory)
    {
        return fileInfos[fileHash];
    }

    function GetFileInfos(bytes[] memory _fileList)
        public
        view
        returns (FileInfo[] memory)
    {
        require(_fileList.length > 0, "fileList is empty");
        FileInfo[] memory _fileInfos = new FileInfo[](_fileList.length);
        for (uint256 i = 0; i < _fileList.length; i++) {
            bytes memory fileHash = _fileList[i];
            _fileInfos[i] = fileInfos[fileHash];
        }
        return _fileInfos;
    }

    function UpdateFileInfo(FileInfo memory f) public payable {
        fileInfos[f.FileHash] = f;
    }

    function DeleteFileInfo(bytes memory fileHash) public {
        delete fileInfos[fileHash];
    }

    function GetFileSectorRefs(bytes memory fileHash)
        public
        view
        returns (SectorRef[] memory)
    {
        return fileSectorRefs[fileHash];
    }

    function AddFileSectorRef(bytes memory fileHash, SectorRef memory sectorRef)
        public
        payable
    {
        fileSectorRefs[fileHash].push(sectorRef);
    }

    function GetFileList(address _walletAddr)
        public
        view
        returns (bytes[] memory)
    {
        return fileList[_walletAddr];
    }

    function UpdateFileList(address _walletAddr, bytes[] memory _fileList)
        public
    {
        fileList[_walletAddr] = _fileList;
    }

    function DelFileFromList(address walletAddr, bytes memory fileHash) public {
        for (uint256 i = 0; i < fileList[walletAddr].length; i++) {
            if (keccak256(fileList[walletAddr][i]) == keccak256(fileHash)) {
                delete fileList[walletAddr][i];
                break;
            }
        }
    }

    function AddFileToFileList(address _walletAddr, bytes memory _fileHash)
        public
    {
        fileList[_walletAddr].push(_fileHash);
    }

    function GetPrimaryFiles(address walletAddr)
        public
        view
        returns (bytes[] memory)
    {
        return primaryFileList[walletAddr];
    }

    function DelFileFromPrimaryList(address walletAddr, bytes memory fileHash)
        public
    {
        for (uint256 i = 0; i < primaryFileList[walletAddr].length; i++) {
            if (
                keccak256(primaryFileList[walletAddr][i]) == keccak256(fileHash)
            ) {
                delete primaryFileList[walletAddr][i];
                break;
            }
        }
    }

    function AddFileToPrimaryList(address walletAddr, bytes memory fileHash)
        public
    {
        primaryFileList[walletAddr].push(fileHash);
    }

    function GetCandidateFiles(address walletAddr)
        public
        view
        returns (bytes[] memory)
    {
        return candidateFileList[walletAddr];
    }

    function DelFileFromCandidateList(address walletAddr, bytes memory fileHash)
        public
    {
        for (uint256 i = 0; i < candidateFileList[walletAddr].length; i++) {
            if (
                keccak256(candidateFileList[walletAddr][i]) ==
                keccak256(fileHash)
            ) {
                delete candidateFileList[walletAddr][i];
                break;
            }
        }
    }

    function AddFileToCandidateList(address walletAddr, bytes memory fileHash)
        public
    {
        candidateFileList[walletAddr].push(fileHash);
    }

    function GetUnSettledFileList(address walletAddr)
        public
        view
        returns (bytes[] memory)
    {
        return unSettledFileList[walletAddr];
    }

    function AddFileToUnSettleList(address fileOwner, bytes memory fileHash)
        public
        payable
    {
        unSettledFileList[fileOwner].push(fileHash);
    }

    function UpdateUnSettleList(address fileOwner, bytes[] memory _fileList)
        public
        payable
    {
        unSettledFileList[fileOwner] = _fileList;
    }

    function DelFileFromUnSettledList(address walletAddr, bytes memory fileHash)
        public
    {
        for (uint256 i = 0; i < unSettledFileList[walletAddr].length; i++) {
            if (
                keccak256(unSettledFileList[walletAddr][i]) ==
                keccak256(fileHash)
            ) {
                delete unSettledFileList[walletAddr][i];
                break;
            }
        }
    }

    function CalcProveTimesByUploadInfo(
        UploadOption memory uploadOption,
        uint256 beginHeight
    ) public pure returns (uint64) {
        return
            uint64(uploadOption.ExpiredHeight - beginHeight) /
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

    function CalcFee(
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

    function CalcDepositFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public pure returns (StorageFee memory) {
        uint64 proveTime = CalcProveTimesByUploadInfo(
            uploadOption,
            currentHeight
        );
        StorageFee memory fee = CalcFee(
            setting,
            proveTime,
            uploadOption.CopyNum,
            uploadOption.FileSize,
            uint64(uploadOption.ExpiredHeight - currentHeight)
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
        if (uploadOption.WhiteList_.length > 0) {
            fee = txGas * 4;
        } else {
            fee = txGas * 3;
        }
        StorageFee memory sf;
        sf.TxnFee = fee;
        if (uploadOption.StorageType_ == StorageType.Normal) {
            return sf;
        }
        StorageFee memory depositFee = CalcDepositFee(
            uploadOption,
            setting,
            currentHeight
        );
        sf.ValidationFee = depositFee.ValidationFee;
        sf.SpaceFee = depositFee.SpaceFee;
        return sf;
    }

    function UpdateFilesForRenew(
        bytes[] memory _fileList,
        Setting memory setting,
        uint256 newExpireHeight
    ) public payable returns (FileInfo[] memory, bool) {
        FileInfo[] memory fileInfo;
        for (uint256 i = 0; i < _fileList.length; i++) {
            FileInfo memory _fileInfo = GetFileInfo(_fileList[i]);
            if (_fileInfo.StorageType_ != StorageType.Normal) {
                continue;
            }
            if (newExpireHeight <= _fileInfo.ExpiredHeight) {
                continue;
            }
            bool res = UpdateFileInfoForRenew(
                setting,
                newExpireHeight,
                _fileInfo
            );
            if (!res) {
                return (fileInfo, false);
            }
        }
        return (fileInfo, true);
    }

    function UpdateFileInfoForRenew(
        Setting memory setting,
        uint256 newExpireHeight,
        FileInfo memory fileInfo
    ) public payable returns (bool) {
        fileInfo.ExpiredHeight = newExpireHeight;
        UploadOption memory uploadOpt;
        uploadOpt.ExpiredHeight = fileInfo.ExpiredHeight;
        uploadOpt.ProveInterval = fileInfo.ProveInterval;
        uploadOpt.CopyNum = fileInfo.CopyNum;
        uploadOpt.FileSize = fileInfo.FileBlockSize * fileInfo.FileBlockNum;
        uint256 beginHeight = fileInfo.BlockHeight;
        StorageFee memory newDeposit = CalcDepositFee(
            uploadOpt,
            setting,
            beginHeight
        );
        uint64 newDepositSum = newDeposit.TxnFee +
            newDeposit.SpaceFee +
            newDeposit.ValidationFee;
        if (newDepositSum <= fileInfo.Deposit) {
            return false;
        }
        fileInfo.Deposit = newDepositSum;
        fileInfo.ProveTimes = CalcProveTimesByUploadInfo(
            uploadOpt,
            beginHeight
        );
        UpdateFileInfo(fileInfo);
        return true;
    }

    function ChangeFileOwner(OwnerChange memory ownerChange) public {
        FileInfo memory fileInfo = GetFileInfo(ownerChange.FileHash);
        require(
            fileInfo.FileOwner == ownerChange.CurOwner,
            "Current owner must be the owner of the file"
        );
        fileInfo.FileOwner = ownerChange.NewOwner;
        UpdateFileInfo(fileInfo);
    }

    function SaveFile(FileInfo memory fileInfo) public {
        UpdateFileInfo(fileInfo);
        AddFileToFileList(fileInfo.FileOwner, fileInfo.FileHash);
        for (uint256 i = 0; i < fileInfo.PrimaryNodes.length; i++) {
            AddFileToPrimaryList(fileInfo.PrimaryNodes[i], fileInfo.FileHash);
        }
        for (uint256 i = 0; i < fileInfo.CandidateNodes.length; i++) {
            AddFileToCandidateList(
                fileInfo.CandidateNodes[i],
                fileInfo.FileHash
            );
        }
    }

    function FileReNew(
        Setting memory setting,
        FileReNewInfo memory fileReNewInfo
    ) public payable returns (string memory) {
        FileInfo memory fileInfo = GetFileInfo(fileReNewInfo.FileHash);
        if (fileInfo.BlockHeight <= 0) {
            return "file not exist";
        }
        if (fileInfo.StorageType_ != StorageType.Professional) {
            return "file type error";
        }
        if (fileInfo.ExpiredHeight <= block.number) {
            return "file expired";
        }
        StorageFee memory totalRenew = CalcFee(
            setting,
            fileReNewInfo.ReNewTimes,
            fileInfo.CopyNum,
            fileInfo.FileBlockNum * fileInfo.FileBlockSize,
            fileReNewInfo.ReNewTimes * fileInfo.ProveInterval
        );
        uint64 reNewFee = totalRenew.ValidationFee + totalRenew.SpaceFee;
        if (msg.value < reNewFee) {
            return "not enough money";
        }
        fileInfo.ProveTimes += fileReNewInfo.ReNewTimes;
        fileInfo.Deposit += reNewFee;
        fileInfo.ExpiredHeight +=
            fileInfo.ProveInterval *
            fileReNewInfo.ReNewTimes;
        UpdateFileInfo(fileInfo);
        return "";
    }

    function ChangeFilePrivilege(PriChange memory priChange) public {
        FileInfo memory fileInfo = GetFileInfo(priChange.fileHash);
        fileInfo.Privilege = priChange.privilege;
        UpdateFileInfo(fileInfo);
    }

    function DeleteFileFromList(FileInfo memory fileInfo) public {
        DelFileFromList(fileInfo.FileOwner, fileInfo.FileHash);
        for (uint256 i = 0; i < fileInfo.PrimaryNodes.length; i++) {
            DelFileFromPrimaryList(fileInfo.PrimaryNodes[i], fileInfo.FileHash);
        }
        for (uint256 i = 0; i < fileInfo.CandidateNodes.length; i++) {
            DelFileFromCandidateList(
                fileInfo.CandidateNodes[i],
                fileInfo.FileHash
            );
        }
    }
}
