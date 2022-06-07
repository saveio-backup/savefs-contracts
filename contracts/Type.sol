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

enum UserSpaceType {
    None,
    Add,
    Revoke
}

enum WhiteListOpType {
    ADD,
    DEL,
    ADD_COV,
    DEL_ALL,
    UPDATE
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
    bytes BlocksRoot;
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

struct Challenge {
    uint32 Index;
    uint32 Rand;
}

struct MerkleNode {
    uint64 Layer;
    uint64 Index;
    bytes Hash;
}

struct MerklePath {
    uint64 PathLen;
    MerkleNode[] Path;
}

struct SectorProveData {
    uint64 ProveFileNum;
    uint64 BlockNum;
    bytes Proofs;
    bytes Tags;
    MerklePath[] MerklePath_;
    bytes PlotData;
}

struct PrepareForPdpVerificationParams {
    SectorInfo SectorInfo_;
    Challenge[] Challenges;
    SectorProveData ProveData;
}

struct PdpVerificationReturns {
    bytes[] FileIDs;
    bytes[] Tags;
    Challenge[] UpdatedChal;
    MerklePath[] Path;
    bytes RootHashes;
    FileInfo FileInfo_;
    bool Success;
}

struct GenChallengeParams {
    address WalletAddr;
    bytes HashValue;
    uint64 FileBlockNum;
    uint64 ProveNum;
}

struct VerifyProofWithMerklePathForFileParams {
    uint64 Version;
    bytes Proofs;
    bytes[] FileIds;
    bytes[] Tags;
    Challenge[] Challenges;
    MerklePath[] MerklePath_;
    bytes RootHashes;
}

struct VerifyPlotDataParams {
    PlotInfo PlotInfo_;
    bytes PlotData;
    uint64 Index;
}

struct ProveParam {
    bytes RootHash;
    bytes[] FileID;
}

struct ProveData {
    bytes Proofs;
    uint64 BlockNum;
    bytes[] Tags;
    MerklePath[] MerklePath_;
}

struct OwnerChange {
    bytes FileHash;
    address CurOwner;
    address NewOwner;
}

struct PriChange {
    bytes fileHash;
    uint64 privilege;
}

struct WhiteListParams {
    bytes FileHash;
    WhiteListOpType Op;
    WhiteList[] List;
}

struct FileProveParams {
    bytes FileHash;
    bytes ProveData;
    uint256 BlockHeight;
    address NodeWallet;
    uint64 Profit;
    uint64 SectorID;
}

struct SectorProveParams {
    address NodeAddr;
    uint64 SectorID;
    uint64 ChallengeHeight;
    bytes ProveData;
}

struct UserSpaceOperation {
    UserSpaceType Type;
    uint64 Value;
}

struct UserSpaceParams {
    address WalletAddr;
    address Owner;
    UserSpaceOperation Size;
    UserSpaceOperation BlockCount;
}

struct FSConfig {
    uint64 DEFAULT_BLOCK_INTERVAL;
    uint64 DEFAULT_PROVE_PERIOD;
    uint64 IN_SECTOR_SIZE;
}

struct ProveConfig {
    uint64 SECTOR_PROVE_BLOCK_NUM;
}
