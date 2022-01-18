//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./Config.sol";
import "./Node.sol";
import "./Space.sol";
import "./Sector.sol";

/**
 * @title FileSystem
 * @dev FileSystem contract
 */
contract FileSystem is Initializable {
    Config config;
    Node node;
    Space space;
    Sector sector;

    uint64 DEFAULT_BLOCK_INTERVAL = 5;
    uint64 DEFAULT_PROVE_PERIOD = (3600 * 24) / DEFAULT_BLOCK_INTERVAL;
    uint64 IN_SECTOR_SIZE = 1000 * 1000;

    mapping(bytes => FileInfo) fileInfos; // fileHash => FileInfo
    mapping(bytes => ProveDetail[]) proveDetail; // fileHash => ProveDetail[]
    mapping(bytes => ProveDetails) proveDetails; // fileHash => ProveDetails
    mapping(bytes => SectorRef[]) fileSectorRefs; // fileHash => SectorRef[]
    mapping(address => bytes[]) fileList; // walletAddr => bytes[]
    mapping(address => bytes[]) primaryFileList; // walletAddr => bytes[]
    mapping(address => bytes[]) candidateFileList; // walletAddr => bytes[]
    mapping(address => bytes[]) unSettledFileList; // walletAddr => bytes[]
    mapping(uint32 => PocProve[]) pocProveList; // blockNumber => PocProve

    event StoreFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        bytes fileHash,
        uint64 fileSize,
        address walletAddr,
        uint64 cost,
        bool isPlotFile
    );

    event DeleteFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        bytes fileHash,
        address walletAddr
    );

    event DeleteFilesEvent(
        FsEvent eventType,
        uint256 blockHeight,
        bytes[] fileHashs,
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
        bytes fileHash,
        address walletAddr
    );

    error FileNotExist(bytes);
    error UserspaceInsufficientBalance(uint256 got, uint256 want);
    error UserspaceInsufficientSpace(uint256 got, uint256 want);
    error UserspaceWrongExpiredHeight(uint256 got, uint256 want);
    error NotEnoughTransfer(uint256 got, uint256 want);
    error DifferenceFileOwner();
    error FileProveNotFileOwner();
    error FileProveFailed();
    error SectorProveFailed();
    error NodeSectorProvedInTimeError();

    modifier NotEmptyFileHash(bytes memory fileHash) {
        require(fileHash.length > 0, "fileHash must be empty");
        _;
    }

    function initialize(
        Config _config,
        Node _node,
        Space _space,
        Sector _sector
    ) public initializer {
        config = _config;
        node = _node;
        space = _space;
        sector = _sector;
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

    function CalcDepositFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public pure returns (StorageFee memory) {
        uint64 proveTime = CalcProveTimesByUploadInfo(
            uploadOption,
            currentHeight
        );
        StorageFee memory fee = calcFee(
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
        StorageFee memory uploadFee = CalcDepositFee(
            uploadOption,
            setting,
            block.number
        );
        fileInfo.Deposit =
            uploadFee.TxnFee +
            uploadFee.SpaceFee +
            uploadFee.ValidationFee;
        fileInfo.ProveTimes = CalcProveTimesByUploadInfo(
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
        fileInfos[fileInfo.FileHash] = fileInfo;
        bytes[] storage f = fileList[fileInfo.FileOwner];
        // TODO should unique file hash
        f.push(fileInfo.FileHash);
        for (uint256 i = 0; i < fileInfo.PrimaryNodes.AddrList.length; i++) {
            bytes[] storage p = fileList[fileInfo.PrimaryNodes.AddrList[i]];
            p.push(fileInfo.FileHash);
        }
        for (uint256 i = 0; i < fileInfo.CandidateNodes.AddrList.length; i++) {
            bytes[] storage p = fileList[fileInfo.CandidateNodes.AddrList[i]];
            p.push(fileInfo.FileHash);
        }
        ProveDetails memory _proveDetails;
        _proveDetails.CopyNum = fileInfo.CopyNum;
        _proveDetails.ProveDetailNum = 0;
        proveDetails[fileInfo.FileHash] = _proveDetails;
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

    function GetFileInfos(bytes[] memory _fileList)
        public
        view
        returns (FileInfo[] memory)
    {
        require(_fileList.length > 0, "fileList is empty");
        FileInfo[] memory _fileInfos = new FileInfo[](_fileList.length);
        for (uint256 i = 0; i < _fileList.length; i++) {
            bytes memory fileHash = _fileList[i];
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
        returns (bytes[] memory)
    {
        return fileList[walletAddr];
    }

    function UpdateFileInfo(FileInfo memory f) public payable {
        fileInfos[f.FileHash] = f;
    }

    function DeleteFileInfo(bytes memory fileHash) public {
        delete fileInfos[fileHash];
    }

    /**
     * Why same as delete file info method?
     */
    function DeleteProveDetails(bytes memory fileHash) public {
        delete fileInfos[fileHash];
    }

    function UpdateFileList(address walletAddr, bytes[] memory list)
        public
        payable
    {
        fileList[walletAddr] = list;
    }

    function UpdateFilesForRenew(
        bytes[] memory _fileList,
        Setting memory setting,
        uint256 newExpireHeight
    ) public view returns (FileInfo[] memory, bool) {
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
    ) public pure returns (bool) {
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
        return true;
    }

    function GetUnSettledFileList(address walletAddr)
        public
        view
        returns (bytes[] memory)
    {
        return unSettledFileList[walletAddr];
    }

    function DelFileFromUnSettledList(address walletAddr, bytes memory fileHash)
        public
        payable
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

    function DelFileFromList(address walletAddr, bytes memory fileHash) public {
        for (uint256 i = 0; i < fileList[walletAddr].length; i++) {
            if (keccak256(fileList[walletAddr][i]) == keccak256(fileHash)) {
                delete fileList[walletAddr][i];
                break;
            }
        }
    }

    function DelFileFromPrimaryList(address walletAddr, bytes memory fileHash)
        public
        payable
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

    function DelFileFromCandidateList(address walletAddr, bytes memory fileHash)
        public
        payable
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

    function cleanupForDeleteFile(
        FileInfo memory fileInfo,
        bool rmInfo,
        bool rmList
    ) public {
        bytes memory fileHash = fileInfo.FileHash;
        if (rmInfo) {
            DeleteFileInfo(fileHash);
            DeleteProveDetails(fileHash);
            DelFileFromUnSettledList(fileInfo.FileOwner, fileHash);
        }
        if (rmList) {
            DelFileFromList(fileInfo.FileOwner, fileHash);
            for (
                uint256 i = 0;
                i < fileInfo.PrimaryNodes.AddrList.length;
                i++
            ) {
                DelFileFromPrimaryList(
                    fileInfo.PrimaryNodes.AddrList[i],
                    fileHash
                );
            }
            for (
                uint256 i = 0;
                i < fileInfo.CandidateNodes.AddrList.length;
                i++
            ) {
                DelFileFromCandidateList(
                    fileInfo.CandidateNodes.AddrList[i],
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
        payable
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
            cleanupForDeleteFile(fileInfo, true, true);
            deletedFiles[n] = _fileList[i];
            n++;
        }
        if (amount > 0) {
            payable(walletAddr).transfer(amount);
        }
        return (deletedFiles, amount, success);
    }

    function DeleteUnsettledFiles(address walletAddr) public payable {
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
            bytes[] memory list = unSettledFileList[walletAddr];
            for (uint256 j = 0; j < list.length; j++) {
                if (keccak256(list[j]) == keccak256(deletedFiles[i])) {
                    delete list[j];
                }
            }
            unSettledFileList[walletAddr] = list;
        }
    }

    function GetUnProvePrimaryFiles(address walletAddr)
        public
        view
        returns (bytes[] memory)
    {
        bytes[] memory list = primaryFileList[walletAddr];
        bytes[] memory unProvePrimaryFiles = new bytes[](list.length);
        uint64 n = 0;
        for (uint256 i = 0; i < list.length; i++) {
            ProveDetail[] memory details = proveDetail[list[i]];
            bool prove = false;
            for (uint256 j = 0; j < details.length; j++) {
                if (details[j].WalletAddr == walletAddr) {
                    prove = true;
                    break;
                }
            }
            if (!prove) {
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
        returns (bytes[] memory)
    {
        bytes[] memory list = candidateFileList[walletAddr];
        bytes[] memory unProveCandidateFiles = new bytes[](list.length);
        uint64 n = 0;
        for (uint256 i = 0; i < list.length; i++) {
            ProveDetail[] memory details = proveDetail[list[i]];
            bool prove = false;
            for (uint256 j = 0; j < details.length; j++) {
                if (details[j].WalletAddr == walletAddr) {
                    prove = true;
                    break;
                }
            }
            if (!prove) {
                continue;
            }
            unProveCandidateFiles[n] = list[i];
            n++;
        }
        return unProveCandidateFiles;
    }

    struct PriChange {
        bytes fileHash;
        uint64 privilege;
    }

    function ChangeFilePivilege(PriChange memory priChange) public payable {
        FileInfo memory fileInfo = GetFileInfo(priChange.fileHash);
        fileInfo.Privilege = priChange.privilege;
        UpdateFileInfo(fileInfo);
    }

    function GetPocProveList(uint32 height)
        public
        view
        returns (PocProve[] memory)
    {
        require(height > 0, "Block number must lager than 0");
        return pocProveList[height];
    }

    function GetFileProveDetails(bytes memory fileHash)
        public
        view
        returns (ProveDetail[] memory)
    {
        ProveDetail[] memory details = proveDetail[fileHash];
        for (uint256 i = 0; i < details.length; i++) {
            NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
                details[i].WalletAddr
            );
            details[i].NodeAddr = nodeInfo.NodeAddr;
        }
        return details;
    }

    struct OwnerChange {
        bytes FileHash;
        address CurOwner;
        address NewOwner;
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

    function deleteFiles(FileInfo[] memory files) public {
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
            SectorRef[] memory sectorRefs = fileSectorRefs[fileInfo.FileHash];
            for (uint256 j = 0; j < sectorRefs.length; j++) {
                SectorInfo memory sectorInfo = sector.GetSectorInfo(
                    sectorRefs[j]
                );
                sector.DeleteFileFromSector(sectorInfo, fileInfo);
            }
            if (fileInfo.Deposit == 0) {
                cleanupForDeleteFile(fileInfo, true, true);
                continue;
            }
            // TODO
        }
    }

    function DeleteFile(bytes memory fileHash) public {
        FileInfo memory fileInfo = GetFileInfo(fileHash);
        FileInfo[] memory files = new FileInfo[](1);
        files[0] = fileInfo;
        deleteFiles(files);
        emit DeleteFileEvent(
            FsEvent.DELETE_FILE,
            block.number,
            fileHash,
            fileInfo.FileOwner
        );
    }

    function DeleteFiles(bytes[] memory fileHashs) public {
        address fileOwner;
        FileInfo[] memory files = new FileInfo[](fileHashs.length);
        for (uint256 i = 0; i < fileHashs.length; i++) {
            FileInfo memory fileInfo = GetFileInfo(fileHashs[i]);
            fileOwner = fileInfo.FileOwner;
            files[i] = fileInfo;
        }
        deleteFiles(files);
        emit DeleteFilesEvent(
            FsEvent.DELETE_FILE,
            block.number,
            fileHashs,
            fileOwner
        );
    }

    function getProveDetailsWithNodeAddr(bytes memory fileHash)
        public
        view
        returns (ProveDetail[] memory)
    {
        ProveDetail[] memory details = proveDetail[fileHash];
        for (uint256 i = 0; i < details.length; i++) {
            NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
                details[i].WalletAddr
            );
            details[i].NodeAddr = nodeInfo.NodeAddr;
        }
        return details;
    }

    struct FileProveParams {
        bytes FileHash;
        bytes ProveData;
        uint256 BlockHeight;
        address NodeWallet;
        uint64 Profit;
        uint64 SectorID;
    }

    function GenChallenge() public pure returns (uint64) {
        // TODO
        return 0;
    }

    function VerifyProofWithMerklePathForFile(uint64)
        public
        pure
        returns (bool)
    {
        // TODO
        return true;
    }

    function checkProve(
        FileProveParams memory fileProve,
        FileInfo memory fileInfo
    ) public view returns (bool) {
        uint256 currBlockHeight = block.number;
        if (
            fileProve.BlockHeight > currBlockHeight + fileInfo.ProveInterval ||
            fileProve.BlockHeight + fileInfo.ProveInterval < currBlockHeight
        ) {
            return false;
        }
        // TODO complete pdp prove
        uint64 challenge = GenChallenge();
        bool result = VerifyProofWithMerklePathForFile(challenge);
        if (!result) {
            return false;
        }
        return true;
    }

    function checkProveExpire(uint256 fileExpiredHeight)
        public
        view
        returns (bool)
    {
        if (block.number > fileExpiredHeight) {
            return true;
        }
        return false;
    }

    function FileProve(FileProveParams memory fileProve) public {
        Setting memory setting = config.GetSetting();
        FileInfo memory fileInfo = GetFileInfo(fileProve.FileHash);
        if (fileInfo.IsPlotFile) {
            if (fileProve.NodeWallet != fileInfo.FileOwner) {
                revert FileProveNotFileOwner();
            }
        } else {
            bool canProve = false;
            NodeList memory list = fileInfo.PrimaryNodes;
            address[] memory primaryNodes = list.AddrList;
            for (uint256 i = 0; i < primaryNodes.length; i++) {
                if (primaryNodes[i] == fileProve.NodeWallet) {
                    canProve = true;
                    break;
                }
            }
            if (!canProve) {
                NodeList memory list2 = fileInfo.CandidateNodes;
                address[] memory candidateNodes = list2.AddrList;
                for (uint256 i = 0; i < candidateNodes.length; i++) {
                    if (candidateNodes[i] == fileProve.NodeWallet) {
                        canProve = true;
                        break;
                    }
                }
            }
            if (!canProve) {
                revert FileProveFailed();
            }
        }
        NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
            fileProve.NodeWallet
        );
        ProveDetail[] memory details = getProveDetailsWithNodeAddr(
            fileInfo.FileHash
        );
        if (fileProve.SectorID != 0 && block.number < fileInfo.ExpiredHeight) {
            for (uint256 i = 0; i < details.length; i++) {
                if (details[i].WalletAddr == fileProve.NodeWallet) {
                    revert FileProveFailed();
                }
            }
        }
        bool success = checkProve(fileProve, fileInfo);
        if (!success) {
            revert FileProveFailed();
        }
        bool found = false;
        bool settleFlag = false;
        uint64 haveProveTimes = 0;
        ProveDetail memory detail;
        uint256 fileExpiredHeight = fileInfo.ExpiredHeight;
        for (uint256 i = 0; i < details.length; i++) {
            if (details[i].WalletAddr == fileProve.NodeWallet) {
                haveProveTimes = detail.ProveTimes;
                if (
                    haveProveTimes == fileInfo.ProveTimes ||
                    block.number > fileExpiredHeight
                ) {
                    detail.Finished = true;
                    settleFlag = true;
                }
                if (haveProveTimes > fileInfo.ProveTimes) {
                    revert FileProveFailed();
                }
                bool r = checkProveExpire(fileInfo.ExpiredHeight);
                if (r) {
                    revert FileProveFailed();
                }
                detail.ProveTimes++;
                found = true;
                break;
            }
        }
        if (!found) {
            if (details.length == fileInfo.CopyNum + 1) {
                revert FileProveFailed();
            }
            if (
                nodeInfo.RestVol <
                fileInfo.FileBlockNum * fileInfo.FileBlockSize
            ) {
                revert FileProveFailed();
            }
            nodeInfo.RestVol -= fileInfo.FileBlockNum * fileInfo.FileBlockSize;
            node.UpdateNodeInfo(nodeInfo);
            detail.NodeAddr = nodeInfo.NodeAddr;
            detail.WalletAddr = fileProve.NodeWallet;
            detail.ProveTimes = 1;
            detail.BlockHeight = block.number;
            detail.Finished = false;
            // TODO
            // details.push(fileProve.FileHash, detail);
        }
        // TODO
        // UpdateProveDetailInfo();
        if (!found) {
            // TODO
        }
        if (settleFlag) {
            // TODO
        }
        emit FilePDPSuccessEvent(
            FsEvent.FILE_PDP_SUCCESS,
            block.number,
            fileInfo.FileHash,
            nodeInfo.WalletAddr
        );
    }

    struct SectorProveParams {
        address NodeAddr;
        uint64 SectorID;
        uint64 ChallengeHeight;
        bytes ProveData;
    }

    function checkSectorProve(
        SectorProveParams memory sectorProve,
        SectorInfo memory sectorInfo
    ) public view returns (bool) {
        // TODO
    }

    function punishForSector(
        SectorInfo memory sectorInfo,
        NodeInfo memory nodeInfo,
        Setting memory setting,
        uint64 times
    ) public {
        // TODO
    }

    function profitSplitForSector(
        SectorInfo memory sectorInfo,
        NodeInfo memory nodeInfo,
        Setting memory setting
    ) public {
        // TODO
    }

    function SectorProve(SectorProveParams memory sectorProve) public payable {
        NodeInfo memory nodeInfo = node.GetNodeInfoByNodeAddr(
            sectorProve.NodeAddr
        );
        SectorInfo memory sectorInfo = sector.GetSectorInfo(
            SectorRef({
                SectorId: sectorProve.SectorID,
                NodeAddr: sectorProve.NodeAddr
            })
        );
        Setting memory setting = config.GetSetting();
        if (block.number < sectorInfo.NextProveHeight) {
            revert SectorProveFailed();
        }
        if (sectorProve.ChallengeHeight != sectorInfo.NextProveHeight) {
            revert SectorProveFailed();
        }
        bool r = checkSectorProve(sectorProve, sectorInfo);
        if (!r) {
            punishForSector(sectorInfo, nodeInfo, setting, 1);
            revert SectorProveFailed();
        }
        profitSplitForSector(sectorInfo, nodeInfo, setting);
        if (sectorInfo.FirstProveHeight == 0) {
            sectorInfo.FirstProveHeight = block.number;
        }
        sectorInfo.NextProveHeight = block.number + setting.DefaultProvePeriod;
        sector.UpdateSectorInfo(sectorInfo);
        if (!sectorInfo.IsPlots) {
            revert SectorProveFailed();
        }
        // TODO poc prove
    }

    function getLastPunishmentHeightForNode(address nodeAddr, uint64 sectorID)
        public
        view
        returns (uint64)
    {
        // TODO
    }

    function calMissingSectorProveTimes(
        SectorInfo memory sectorInfo,
        Setting memory setting,
        uint256 lastHeight,
        uint256 nowHeight
    ) public view returns (uint64) {
        // TODO
    }

    function CheckNodeSectorProvedInTime(SectorRef memory sectorRef)
        public
        payable
    {
        address nodeAddr = sectorRef.NodeAddr;
        uint64 sectorID = sectorRef.SectorId;
        NodeInfo memory nodeInfo = node.GetNodeInfoByNodeAddr(nodeAddr);
        if (nodeInfo.ServiceTime < block.timestamp) {
            revert NodeSectorProvedInTimeError();
        }
        if (sectorID == 0) {
            revert NodeSectorProvedInTimeError();
        }
        SectorInfo memory sectorInfo = sector.GetSectorInfo(sectorRef);
        if (sectorInfo.FileNum == 0) {
            revert NodeSectorProvedInTimeError();
        }
        Setting memory setting = config.GetSettingWithProveLevel(
            sectorInfo.ProveLevel_
        );
        uint256 height = block.number;
        if (sectorInfo.NextProveHeight + setting.DefaultProvePeriod < height) {
            revert NodeSectorProvedInTimeError();
        }
        uint256 lastHeight = getLastPunishmentHeightForNode(nodeAddr, sectorID);
        uint64 times = calMissingSectorProveTimes(
            sectorInfo,
            setting,
            lastHeight,
            height
        );
        if (times == 0) {
            revert NodeSectorProvedInTimeError();
        }
        punishForSector(sectorInfo, nodeInfo, setting, times);
    }
}
