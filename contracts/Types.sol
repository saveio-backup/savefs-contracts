//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

/** enum ********************** */
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

enum WHileListOp {
    ADD,
    DEL,
    ADD_COV,
    DEL_ALL,
    UPDATE
}

/** setting ********************** */
struct Setting {
    uint64 GasPrice;
    uint64 GasPerGBPerBlock;
    uint64 GasPerKBPerBlock;
    uint64 GasForChallenge;
    uint64 MaxProveBlocks;
    uint64 MinVolume;
    uint64 DefaultProvePeriod;
    uint64 DefaultProveLevel;
    uint64 DefaultCopyNum;
}

/** node info ******************** */
struct NodeInfo {
    uint64 Pledge;
    uint64 Profit;
    uint64 Volume;
    uint64 RestVol;
    uint64 ServiceTime;
    address WalletAddr;
    address NodeAddr;
}

struct NodeList {
    uint64 AddrNum;
    address[] AddrList;
}

/** userspace ******************** */
struct UserSpace {
    uint64 Used;
    uint64 Remain;
    uint64 ExpireHeight;
    uint64 Balance;
    uint64 UpdateHeight;
}

/** file ********************** */
struct StorageFee {
    uint64 TxnFee;
    uint64 SpaceFee;
    uint64 ValidationFee;
}

struct Role {
    address Addr;
    uint64 BaseHeight;
    uint64 ExpireHeight;
}

struct WhiteListOp {
    bytes FileHash;
    WHileListOp Op;
    WhiteList List;
}

struct WhiteList {
    uint64 Num;
    Role[] List;
}

struct UploadOption {
    bytes FileDesc;
    uint64 FileSize;
    uint64 ProveInterval;
    uint64 ProveLevel;
    uint64 ExpiredHeight;
    uint64 Privilege;
    uint64 CopyNum;
    bool Encrypt;
    bytes EncryptPassword;
    bool RegisterDNS;
    bool BindDNS;
    bytes DnsURL;
    WhiteList WhiteList_;
    bool Share;
    uint64 StorageType;
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
    uint64 FileBlockSize;
    uint64 ProveInterval;
    uint64 ProveTimes;
    uint64 ExpiredHeight;
    uint64 CopyNum;
    uint64 Deposit;
    bytes FileProveParam;
    uint64 BlockHeight;
    bool ValidFlag;
    uint64 StorageType;
    uint64 RealFileSize;
    NodeList PrimaryNodes;
    NodeList CandidateNodes;
    bytes ProveLevel;
    SectorRef[] SectorRes;
    bool IsPlotFile;
    PlotInfo PlotInfo_;
}

struct FileList {
    uint64 FileNum;
    bytes[] List; // fileHash = bytes
}

struct SectorInfo {
    address NodeAddr;
    uint64 SectorID;
    uint64 Size;
    uint64 Used;
    uint64 ProveLevel;
    uint64 FirstProveHeight;
    uint64 NextProveHeight;
    uint64 TotalBlockNum;
    uint64 FileNum;
    uint64 GroupNum;
    bool IsPlots;
    FileList FileList_;
}

struct SectorInfos {
    uint64 SectorCount;
    uint64[] SectorIds;
}

struct PocProve {
    address Miner;
    uint32 Height;
    uint64 PlotSize;
}
