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

    mapping(address => bytes[]) fileList; // walletAddr => bytes[]
    mapping(address => bytes[]) primaryFileList; // walletAddr => bytes[]
    mapping(address => bytes[]) candidateFileList; // walletAddr => bytes[]
    mapping(address => bytes[]) unSettledFileList; // walletAddr => bytes[]

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
            FileInfo memory fileInfo = fileInfos[fileHash];
            if (fileInfo.FileHash.length == 0) {
                revert FileNotExist(fileHash);
            }
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

    function UpdateUnSettleList(address fileOwner, bytes[] memory fileList)
        public
        payable
    {
        unSettledFileList[fileOwner] = fileList;
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
}
