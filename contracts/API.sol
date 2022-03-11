//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;
import "./Type.sol";

interface IConfig {
    function GetSetting() external pure returns (Setting memory);

    function GetSettingWithProveLevel(ProveLevel proveLevel)
        external
        pure
        returns (Setting memory);
}

interface IFileSystem {
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

    error FileNotExist(bytes);
    error UserspaceInsufficientBalance(uint256 got, uint256 want);
    error UserspaceInsufficientSpace(uint256 got, uint256 want);
    error UserspaceWrongExpiredHeight(uint256 got, uint256 want);
    error NotEnoughTransfer(uint256 got, uint256 want);
    error DifferenceFileOwner();
    error InvalidProfit();
    error OpError(uint64);

    function StoreFile(FileInfo memory fileInfo) external payable;

    function ChangeFileOwner(OwnerChange memory ownerChange) external;

    function ChangeFilePrivilege(PriChange memory priChange) external;

    function DeleteFile(bytes memory fileHash) external;

    function DeleteFiles(bytes[] memory fileHashs) external;
}
