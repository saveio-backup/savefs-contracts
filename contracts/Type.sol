//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

enum FsEvent {
    STORE_FILE,
    DELETE_FILE,
    DELETE_FILES,
    SET_USER_SPACE,
    REG_NODE,
    UN_REG_NODE,
    PROVE_FILE,
    FILE_PDP_SUCCESS,
    CREATE_SECTOR,
    DELETE_SECTOR
}

enum ProveLevel {
    HIGH,
    MEDIUM,
    LOW
}

enum StorageType {
    Normal,
    Professional
}

struct Setting {
    uint64 GasPrice;
    uint64 GasPerGBPerBlock;
    uint64 GasPerKBPerBlock;
    uint64 GasForChallenge;
    uint64 MaxProveBlockNum;
    uint64 MinVolume;
    uint64 DefaultProvePeriod;
    uint64 DefaultProveLevel;
    uint64 DefaultCopyNum;
    uint64 DefaultBlockInterval;
    uint64 MinSectorSize;
}

struct NodeInfo {
    uint64 Pledge;
    uint64 Profit;
    uint64 Volume;
    uint64 RestVol;
    uint64 ServiceTime;
    address WalletAddr;
    address NodeAddr;
}

struct UserSpace {
    uint64 Used;
    uint64 Remain;
    uint64 Balance;
    uint256 ExpireHeight;
    uint256 UpdateHeight;
}

enum UserSpaceType {
    None,
    Add,
    Revoke
}

struct StorageFee {
    uint64 TxnFee;
    uint64 SpaceFee;
    uint64 ValidationFee;
}

struct WhiteList {
    address Addr;
    uint64 BaseHeight;
    uint64 ExpireHeight;
}

struct UploadOption {
    bytes FileDesc;
    uint64 FileSize;
    uint64 ProveInterval;
    uint64 ProveLevel;
    uint256 ExpiredHeight;
    uint64 Privilege;
    uint64 CopyNum;
    bool Encrypt;
    bytes EncryptPassword;
    bool RegisterDNS;
    bool BindDNS;
    bytes DnsURL;
    WhiteList[] WhiteList_;
    bool Share;
    StorageType StorageType_;
}

struct SectorRef {
    address NodeAddr;
    uint64 SectorId;
}

struct PlotInfo {
    uint64 NumberID;
    uint64 StartNonce;
    uint64 Nonces;
}

struct FileInfo {
    bytes FileHash;
    address FileOwner;
    bytes FileDesc;
    uint64 Privilege;
    uint64 FileBlockNum;
    uint64 FileBlockSize;
    uint64 ProveInterval;
    uint64 ProveTimes;
    uint256 ExpiredHeight;
    uint64 CopyNum;
    uint64 Deposit;
    bytes FileProveParam;
    uint64 ProveBlockNum;
    uint256 BlockHeight;
    bool ValidFlag;
    StorageType StorageType_;
    uint64 RealFileSize;
    address[] PrimaryNodes;
    address[] CandidateNodes;
    ProveLevel ProveLevel_;
    bool IsPlotFile;
    PlotInfo PlotInfo_;
}

struct SectorInfo {
    address NodeAddr;
    uint64 SectorID;
    uint64 Size;
    uint64 Used;
    ProveLevel ProveLevel_;
    uint256 FirstProveHeight;
    uint256 NextProveHeight;
    uint64 TotalBlockNum;
    uint64 FileNum;
    uint64 GroupNum;
    bool IsPlots;
    bytes[] FileList;
}

struct SectorInfos {
    uint64 SectorCount;
    uint64[] SectorIds;
}

struct FileReNewInfo {
    bytes FileHash;
    address FromAddr;
    uint64 ReNewTimes;
}

struct PocProve {
    address Miner;
    uint256 Height;
    uint64 PlotSize;
}

struct ProveDetail {
    address NodeAddr;
    address WalletAddr;
    uint64 ProveTimes;
    uint256 BlockHeight;
    bool Finished;
}

struct ProveDetailMeta {
    uint64 CopyNum;
    uint64 ProveDetailNum;
}

struct TransferState {
    address From;
    address To;
    uint64 Value;
}
