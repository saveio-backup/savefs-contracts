//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

struct FsSetting {
    uint64 FsGasPrice;
    uint64 GasPerGBPerBlock;
    uint64 GasPerKBPerBlock;
    uint64 GasForChallenge;
    uint64 MaxProveBlocks;
    uint64 MinVolume;
    uint64 DefaultProvePeriod;
    uint64 DefaultProveLevel;
    uint64 DefaultCopyNum;
}

struct StorageFee {
    uint64 TxnFee;
    uint64 SpaceFee;
    uint64 ValidationFee;
}

struct FsNodeInfo {
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

struct Role {
    address Addr;
    uint64 BaseHeight;
    uint64 ExpireHeight;
}

struct WhiteListT {
    uint64 Num;
    Role[] List;
}

struct UploadOption {
    bytes[] FileDesc;
    uint64 FileSize;
    uint64 ProveInterval;
    uint64 ProveLevel;
    uint64 ExpiredHeight;
    uint64 Privilege;
    uint64 CopyNum;
    bool Encrypt;
    bytes[] EncryptPassword;
    bool RegisterDNS;
    bool BindDNS;
    bytes[] DnsURL;
    WhiteListT WhiteList;
    bool Share;
    uint64 StorageType;
}
