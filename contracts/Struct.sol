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
