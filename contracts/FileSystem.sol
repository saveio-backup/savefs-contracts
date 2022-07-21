//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";
import "./IterableMapping.sol";

contract FileSystem is
Initializable,
IConfig,
IFile,
IList,
INode,
IPDP,
IProve,
ISector,
ISpace
{
    IConfig config;
    IFile file;
    IList list;
    INode node;
    IPDP pdp;
    IProve prove;
    ISector sector;
    ISpace space;

    function initialize(
        IConfig _config,
        IFile _file,
        IList _list,
        INode _node,
        IPDP _pdp,
        ISpace _space,
        ISector _sector,
        IProve _prove
    ) public initializer {
        config = _config;
        file = _file;
        list = _list;
        node = _node;
        pdp = _pdp;
        prove = _prove;
        sector = _sector;
        space = _space;
    }

    function StoreFile(FileInfo memory fileInfo)
    public
    payable
    virtual
    override
    {
        file.StoreFile(fileInfo);
    }

    function FileReNew(FileReNewInfo memory fileReNewInfo)
    public
    payable
    virtual
    override
    {
        file.FileReNew(fileReNewInfo);
    }

    function ChangeFileOwner(OwnerChange memory ownerChange)
    public
    virtual
    override
    {
        file.ChangeFileOwner(ownerChange);
    }

    function ChangeFilePrivilege(PriChange memory priChange)
    public
    virtual
    override
    {
        file.ChangeFilePrivilege(priChange);
    }

    function DeleteFile(bytes memory fileHash) public virtual override {
        file.DeleteFile(fileHash);
    }

    function DeleteFiles(bytes[] memory fileHashs) public virtual override {
        file.DeleteFiles(fileHashs);
    }

    function DeleteUnsettledFiles(address walletAddr) public virtual override {
        file.DeleteUnsettledFiles(walletAddr);
    }

    function GetFileInfo(bytes memory fileHash)
    public
    view
    virtual
    override
    returns (FileInfo memory)
    {
        return file.GetFileInfo(fileHash);
    }

    function GetFileInfos(bytes[] memory _fileList)
    public
    view
    virtual
    override
    returns (FileInfo[] memory)
    {
        return file.GetFileInfos(_fileList);
    }

    function GetFileList(address walletAddr)
    public
    view
    virtual
    override
    returns (bytes[] memory)
    {
        return file.GetFileList(walletAddr);
    }

    function GetUnProvePrimaryFiles(address walletAddr)
    public
    view
    virtual
    override
    returns (bytes[] memory)
    {
        return file.GetUnProvePrimaryFiles(walletAddr);
    }

    function GetUnProveCandidateFiles(address walletAddr)
    public
    view
    virtual
    override
    returns (bytes[] memory)
    {
        return file.GetUnProveCandidateFiles(walletAddr);
    }

    function GetUnSettledFileList(address walletAddr)
    public
    view
    virtual
    override
    returns (bytes[] memory)
    {
        return file.GetUnSettledFileList(walletAddr);
    }

    function GetUploadStorageFee(UploadOption memory uploadOption)
    public
    view
    virtual
    override
    returns (StorageFee memory)
    {
        return file.GetUploadStorageFee(uploadOption);
    }

    function UpdateFileInfo(FileInfo memory f) public virtual override payable {
        file.UpdateFileInfo(f);
    }

    function CleanupForDeleteFile(
        FileInfo memory fileInfo,
        bool rmInfo,
        bool rmList
    ) public virtual override payable {
        file.CleanupForDeleteFile(fileInfo, rmInfo, rmList);
    }

    function AddFileToUnSettleList(address fileOwner, bytes memory fileHash)
    public
    virtual
    override
    payable
    {
        file.AddFileToUnSettleList(fileOwner, fileHash);
    }

    function CalcFee(
        Setting memory setting,
        uint64 proveTime,
        uint64 copyNum,
        uint64 fileSize,
        uint64 duration
    ) public pure virtual
    override returns (StorageFee memory) {
        return file.CalcFee(setting, proveTime, copyNum, fileSize, duration);
    }

    function CalcDepositFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public pure virtual
    override returns (StorageFee memory) {
        return file.CalcDepositFee(uploadOption, setting, currentHeight);
    }

    function UpdateFilesForRenew(
        bytes[] memory _fileList,
        Setting memory setting,
        uint256 newExpireHeight
    ) public view virtual
    override returns (FileInfo[] memory, bool) {
        return file.UpdateFilesForRenew(_fileList, setting, newExpireHeight);
    }

    function DeleteExpiredFilesFromList(
        bytes[] memory _fileList,
        address walletAddr,
        StorageType[] memory storageType
    )
    public virtual
    override
    returns (
        bytes[] memory,
        uint64,
        bool
    ) {
        return file.DeleteExpiredFilesFromList(_fileList, walletAddr, storageType);
    }

    function UpdateFileList(address walletAddr, bytes[] memory _list)
    public virtual
    override
    payable {
        file.UpdateFileList(walletAddr, _list);
    }


}
