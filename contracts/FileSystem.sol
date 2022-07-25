//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";
import "./IterableMapping.sol";

/**
 * This contract as Main enter but can't use because
 * code size large than 24 KB.
 */
contract FileSystem is
    Initializable,
    IConfig,
    IFile,
    IList,
    INode,
    IPDP,
    IProve,
    ISector,
    ISpace,
    IFsEvent
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

    /************** config *************/

    function GetSetting()
        public
        view
        virtual
        override
        returns (Setting memory)
    {
        return config.GetSetting();
    }

    function GetSettingWithProveLevel(ProveLevel proveLevel)
        public
        view
        virtual
        override
        returns (Setting memory)
    {
        return config.GetSettingWithProveLevel(proveLevel);
    }

    /************** file *************/

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

    function UpdateFileInfo(FileInfo memory f) public payable virtual override {
        file.UpdateFileInfo(f);
    }

    function CleanupForDeleteFile(
        FileInfo memory fileInfo,
        bool rmInfo,
        bool rmList
    ) public payable virtual override {
        file.CleanupForDeleteFile(fileInfo, rmInfo, rmList);
    }

    function AddFileToUnSettleList(address fileOwner, bytes memory fileHash)
        public
        payable
        virtual
        override
    {
        file.AddFileToUnSettleList(fileOwner, fileHash);
    }

    function CalcFee(
        Setting memory setting,
        uint64 proveTime,
        uint64 copyNum,
        uint64 fileSize,
        uint64 duration
    ) public view virtual override returns (StorageFee memory) {
        return file.CalcFee(setting, proveTime, copyNum, fileSize, duration);
    }

    function CalcDepositFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) public view virtual override returns (StorageFee memory) {
        return file.CalcDepositFee(uploadOption, setting, currentHeight);
    }

    function UpdateFilesForRenew(
        bytes[] memory _fileList,
        Setting memory setting,
        uint256 newExpireHeight
    ) public payable virtual override returns (FileInfo[] memory, bool) {
        return file.UpdateFilesForRenew(_fileList, setting, newExpireHeight);
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
        return
            file.DeleteExpiredFilesFromList(_fileList, walletAddr, storageType);
    }

    function UpdateFileList(address walletAddr, bytes[] memory _list)
        public
        payable
        virtual
        override
    {
        file.UpdateFileList(walletAddr, _list);
    }

    /************** list *************/

    function WhiteListOperate(WhiteListParams memory params)
        public
        virtual
        override
    {
        list.WhiteListOperate(params);
    }

    function GetWhiteList(bytes memory fileHash)
        public
        view
        virtual
        override
        returns (WhiteList[] memory)
    {
        return list.GetWhiteList(fileHash);
    }

    /************** node *************/

    function CalculateNodePledge(NodeInfo memory nodeInfo)
        external
        view
        virtual
        override
        returns (uint64)
    {
        return node.CalculateNodePledge(nodeInfo);
    }

    function Cancel(address walletAddr) external virtual override {
        node.Cancel(walletAddr);
    }

    function GetNodeInfoByNodeAddr(address nodeAddr)
        external
        view
        virtual
        override
        returns (NodeInfo memory)
    {
        return node.GetNodeInfoByNodeAddr(nodeAddr);
    }

    function GetNodeInfoByWalletAddr(address walletAddr)
        external
        view
        virtual
        override
        returns (NodeInfo memory)
    {
        return node.GetNodeInfoByWalletAddr(walletAddr);
    }

    function GetNodeList()
        external
        view
        virtual
        override
        returns (NodeInfo[] memory)
    {
        return node.GetNodeList();
    }

    function NodeUpdate(NodeInfo memory nodeInfo)
        external
        payable
        virtual
        override
    {
        node.NodeUpdate(nodeInfo);
    }

    function Register(NodeInfo memory nodeInfo)
        external
        payable
        virtual
        override
    {
        node.Register(nodeInfo);
    }

    function UpdateNodeInfo(NodeInfo memory nodeInfo)
        external
        payable
        virtual
        override
    {
        node.UpdateNodeInfo(nodeInfo);
    }

    function WithDrawProfit(address walletAddr) external virtual override {
        node.WithDrawProfit(walletAddr);
    }

    function IsNodeRegisted(address walletAddr)
        external
        view
        virtual
        override
        returns (bool)
    {
        return node.IsNodeRegisted(walletAddr);
    }

    /************** pdp *************/

    function GenChallenge(GenChallengeParams memory gParams)
        public
        view
        virtual
        override
        returns (Challenge[] memory)
    {
        return pdp.GenChallenge(gParams);
    }

    function VerifyProofWithMerklePathForFile(
        VerifyProofWithMerklePathForFileParams memory vParams
    ) public view virtual override returns (bool) {
        return pdp.VerifyProofWithMerklePathForFile(vParams);
    }

    function PrepareForPdpVerification(
        PrepareForPdpVerificationParams memory pParams
    ) public view virtual override returns (PdpVerificationReturns memory) {
        return pdp.PrepareForPdpVerification(pParams);
    }

    function VerifyPlotData(VerifyPlotDataParams memory vParams)
        public
        view
        virtual
        override
        returns (bool)
    {
        return pdp.VerifyPlotData(vParams);
    }

    /************** prove *************/

    function CheckNodeSectorProvedInTime(SectorRef memory sectorRef)
        public
        payable
        virtual
        override
    {
        return prove.CheckNodeSectorProvedInTime(sectorRef);
    }

    function FileProve(FileProveParams memory fileProve)
        public
        payable
        virtual
        override
    {
        prove.FileProve(fileProve);
    }

    function GetProveDetailList(bytes memory fileHash)
        public
        view
        virtual
        override
        returns (ProveDetail[] memory)
    {
        return prove.GetProveDetailList(fileHash);
    }

    function SectorProve(SectorProveParams memory sectorProve)
        public
        payable
        virtual
        override
    {
        prove.SectorProve(sectorProve);
    }

    function UpdateProveDetailMeta(
        bytes memory fileHash,
        ProveDetailMeta memory details
    ) public payable virtual override {
        prove.UpdateProveDetailMeta(fileHash, details);
    }

    function DeleteProveDetails(bytes memory fileHash)
        public
        payable
        virtual
        override
    {
        prove.DeleteProveDetails(fileHash);
    }

    /************** sector *************/

    function CreateSector(SectorInfo memory sectorInfo)
        public
        virtual
        override
    {
        sector.CreateSector(sectorInfo);
    }

    function DeleteSector(SectorRef memory sectorRef) public virtual override {
        sector.DeleteSector(sectorRef);
    }

    function GetSectorInfo(SectorRef memory sectorRef)
        public
        view
        virtual
        override
        returns (SectorInfo memory)
    {
        return sector.GetSectorInfo(sectorRef);
    }

    function GetSectorsForNode(address nodeAddr)
        public
        view
        virtual
        override
        returns (SectorInfo[] memory)
    {
        return sector.GetSectorsForNode(nodeAddr);
    }

    function DeleteFileFromSector(
        SectorInfo memory sectorInfo,
        FileInfo memory fileInfo
    ) public payable virtual override {
        sector.DeleteFileFromSector(sectorInfo, fileInfo);
    }

    function AddFileToSector(
        SectorInfo memory sectorInfo,
        FileInfo memory fileInfo
    ) public payable virtual override {
        sector.AddFileToSector(sectorInfo, fileInfo);
    }

    function AddSectorRefForFileInfo(SectorInfo memory sectorInfo)
        public
        payable
        virtual
        override
    {
        sector.AddSectorRefForFileInfo(sectorInfo);
    }

    function UpdateSectorInfo(SectorInfo memory _sector)
        public
        payable
        virtual
        override
    {
        sector.UpdateSectorInfo(_sector);
    }

    /************** space *************/

    function DeleteUserSpace(address walletAddr) public virtual override {
        space.DeleteUserSpace(walletAddr);
    }

    function GetUpdateCost(UserSpaceParams memory params)
        public
        payable
        virtual
        override
        returns (TransferState memory)
    {
        return space.GetUpdateCost(params);
    }

    function GetUserSpace(address walletAddr)
        public
        view
        virtual
        override
        returns (UserSpace memory)
    {
        return space.GetUserSpace(walletAddr);
    }

    function ManageUserSpace(UserSpaceParams memory params)
        public
        payable
        virtual
        override
    {
        space.ManageUserSpace(params);
    }

    function UpdateUserSpace(address walletAddr, UserSpace memory _userSpace)
        public
        payable
        virtual
        override
    {
        space.UpdateUserSpace(walletAddr, _userSpace);
    }
}
