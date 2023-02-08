//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;
import "./type.sol";

interface IFsEvent {
    event FsError(string method, string msg);

    // file
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

    // node
    event RegisterNodeEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        bytes nodeAddr,
        uint64 volume,
        uint64 serviceTime
    );

    event UnRegisterNodeEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr
    );

    // prove
    event FilePDPSuccessEvent(
        FsEvent eventType,
        uint256 blockHeight,
        bytes fileHash,
        address walletAddr
    );

    event ProveFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 profit
    );

    // sector

    event CreateSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId,
        ProveLevel proveLevel,
        uint64 size,
        bool isPlots
    );

    event DeleteSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId
    );

    // space

    event SetUserSpaceEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        UserSpaceType sizeType,
        uint64 size,
        UserSpaceType countType,
        uint64 count
    );

    event GetUpdateCostEvent(TransferState state);

    // dns
    event DnsError(string method, string msg);

    event NotifyNameInfoAdd(address owner, bytes url, NameInfo newer);

    event NotifyHeaderAdd(address owner, bytes header);

    event NotifyNameInfoTransfer(address from, address to, bytes url);

    event NotifyHeaderTransfer(address from, address to, bytes header);

    event NotifyNameInfoChange(address owner, bytes url);

    event DNSNodeRegister(
        bytes ip,
        bytes port,
        address walletAddr,
        uint64 deposit
    );

    event DNSNodeUnReg(address walletAddr);
}

interface IConfig {
    function GetSetting() external view returns (Setting memory);

    function GetSettingWithProveLevel(ProveLevel proveLevel)
        external
        view
        returns (Setting memory);
}

interface IFile {
    function StoreFile(FileInfo memory fileInfo) external payable;

    function FileReNew(FileReNewInfo memory fileReNewInfo) external payable;

    function ChangeFileOwner(OwnerChange memory ownerChange) external;

    function ChangeFilePrivilege(PriChange memory priChange) external;

    function DeleteFile(bytes memory fileHash) external payable;

    function DeleteFiles(bytes[] memory fileHashs) external payable;

    function DeleteUnsettledFiles(address walletAddr) external payable;

    function GetFileInfo(bytes memory fileHash)
        external
        view
        returns (FileInfo memory);

    function GetFileInfos(bytes[] memory _fileList)
        external
        view
        returns (FileInfo[] memory);

    function GetFileList(address walletAddr)
        external
        view
        returns (bytes[] memory);

    function GetUnProvePrimaryFiles(address walletAddr)
        external
        view
        returns (bytes[] memory);

    function GetUnProveCandidateFiles(address walletAddr)
        external
        view
        returns (bytes[] memory);

    function GetUnSettledFileList(address walletAddr)
        external
        view
        returns (bytes[] memory);

    function GetUploadStorageFee(UploadOption memory uploadOption)
        external
        view
        returns (StorageFee memory);

    function UpdateFileInfo(FileInfo memory f) external payable;

    function CleanupForDeleteFile(
        FileInfo memory fileInfo,
        bool rmInfo,
        bool rmList
    ) external payable;

    function AddFileToUnSettleList(address fileOwner, bytes memory fileHash)
        external
        payable;

    function CalcFee(
        Setting memory setting,
        uint64 proveTime,
        uint64 copyNum,
        uint64 fileSize,
        uint64 duration
    ) external view returns (StorageFee memory);

    function CalcDepositFee(
        UploadOption memory uploadOption,
        Setting memory setting,
        uint256 currentHeight
    ) external view returns (StorageFee memory);

    function UpdateFilesForRenew(
        bytes[] memory _fileList,
        Setting memory setting,
        uint256 newExpireHeight
    ) external payable returns (FileInfo[] memory, bool);

    function DeleteExpiredFilesFromList(
        bytes[] memory _fileList,
        address walletAddr,
        StorageType[] memory storageType
    )
        external
        returns (
            bytes[] memory,
            uint64,
            bool
        );

    function UpdateFileList(address walletAddr, bytes[] memory list)
        external
        payable;

    function AddFileSectorRef(bytes memory fileHash, SectorRef memory ref)
        external
        payable;
}

interface IList {
    function WhiteListOperate(WhiteListParams memory params) external;

    function GetWhiteList(bytes memory fileHash)
        external
        view
        returns (WhiteList[] memory);
}

interface INode {
    function CalculateNodePledge(NodeInfo memory nodeInfo)
        external
        view
        returns (uint64);

    function Cancel(address walletAddr) external;

    function GetNodeInfoByNodeAddr(bytes memory nodeAddr)
        external
        view
        returns (NodeInfo memory);

    function GetNodeInfoByWalletAddr(address walletAddr)
        external
        view
        returns (NodeInfo memory);

    function GetNodeList() external view returns (NodeInfo[] memory);

    function NodeUpdate(NodeInfo memory nodeInfo) external payable;

    function Register(NodeInfo memory nodeInfo) external payable;

    function UpdateNodeInfo(NodeInfo memory nodeInfo) external payable;

    function WithDrawProfit(address walletAddr) external;

    function IsNodeRegisted(address walletAddr) external view returns (bool);
}

interface IPDP {
    function GenChallenge(GenChallengeParams memory gParams)
        external
        view
        returns (Challenge[] memory);

    function VerifyProof(ProofRecord memory vParams) external payable;

    function VerifyProofWithMerklePathForFile(ProofParams memory vParams)
        external
        view
        returns (bool);

    function PrepareForPdpVerification(
        PrepareForPdpVerificationParams memory pParams
    ) external view returns (PdpVerificationReturns memory);

    function VerifyPlotData(VerifyPlotDataParams memory vParams)
        external
        view
        returns (bool);
}

interface IProve {
    function CheckNodeSectorProvedInTime(SectorRef memory sectorRef)
        external
        payable;

    function FileProve(FileProveParams memory fileProve) external payable;

    function GetProveDetailList(bytes memory fileHash)
        external
        view
        returns (ProveDetail[] memory);

    function SectorProve(SectorProveParams memory sectorProve) external payable;

    function UpdateProveDetailMeta(
        bytes memory fileHash,
        ProveDetailMeta memory details
    ) external payable;

    function DeleteProveDetails(bytes memory fileHash) external payable;
}

interface ISector {
    function CreateSector(SectorInfo memory sectorInfo) external;

    function DeleteSector(SectorRef memory sectorRef) external;

    function GetSectorInfo(SectorRef memory sectorRef)
        external
        view
        returns (SectorInfo memory);

    function GetSectorsForNode(address nodeAddr)
        external
        view
        returns (SectorInfo[] memory);

    function DeleteFileFromSector(
        SectorInfo memory sectorInfo,
        FileInfo memory fileInfo
    ) external payable;

    function AddFileToSector(
        SectorInfo memory sectorInfo,
        FileInfo memory fileInfo
    ) external payable;

    function UpdateSectorInfo(SectorInfo memory sector) external payable;

    function IsSectorRefByFileInfo(address nodeAddr, uint64 sectorID)
        external
        view
        returns (bool);
}

interface ISpace {
    function DeleteUserSpace(address walletAddr) external;

    function GetUpdateCost(UserSpaceParams memory params)
        external
        view
        returns (TransferState memory);

    function GetUserSpace(address walletAddr)
        external
        view
        returns (UserSpace memory);

    function ManageUserSpace(UserSpaceParams memory params) external payable;

    function UpdateUserSpace(address walletAddr, UserSpace memory _userSpace)
        external
        payable;
}
