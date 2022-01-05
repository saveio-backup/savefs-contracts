//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./IFileSystem.sol";
import "./Struct.sol";
import "./Error.sol";
import "./Event.sol";
import "./Enum.sol";

contract FileSystem is Initializable, IFileSystem {
    FsNodeInfo[] nodeInfo;
    NodeList nodeList;

    event RegisterNodeEvent(FsEvent eventType);

    function initialize() public initializer {
        console.log("initializer");
    }

    function FsGetSettings() public pure override returns (FsSetting memory) {
        FsSetting memory fsSetting;
        fsSetting.FsGasPrice = 1;
        fsSetting.GasPerGBPerBlock = 1;
        fsSetting.GasPerKBPerBlock = 1;
        fsSetting.GasForChallenge = 200000;
        fsSetting.MaxProveBlocks = 32;
        fsSetting.MinVolume = 1000 * 1000;
        fsSetting.DefaultProvePeriod = (3600 * 24) / 5;
        fsSetting.DefaultProveLevel = 1;
        fsSetting.DefaultCopyNum = 2;
        return fsSetting;
    }

    function FsGetUploadStorageFee(UploadOption memory uploadOption)
        public
        view
        override
        returns (StorageFee memory)
    {
        require(uploadOption.FileSize > 0, "fileSize must be greater than 0");
        FsSetting memory fsSetting = FsGetSettings();
        StorageFee memory storageFee;
        uint64 fee;
        uint64 txGas = 10000000;
        if (uploadOption.WhiteList.Num > 0) {
            fee = txGas * 4;
        } else {
            fee = txGas * 3;
        }

        uint64 proveTime = (uploadOption.ExpiredHeight - uint64(block.number)) /
            uploadOption.ProveInterval +
            1;
        uint64 validFee = (uploadOption.CopyNum + 1) *
            uint64(
                proveTime +
                    (fsSetting.GasForChallenge * uploadOption.FileSize) /
                    1024000
            );
        uint64 spaceFee = ((uploadOption.CopyNum + 1) *
            fsSetting.GasPerGBPerBlock *
            uploadOption.FileSize) / uint64(1024000);
        storageFee.TxnFee = fee;
        storageFee.ValidationFee = validFee;
        storageFee.SpaceFee = spaceFee;
        return storageFee;
    }

    function FsNodeRegister(FsNodeInfo memory fsNodeInfo)
        public
        payable
        override
    {
        require(
            fsNodeInfo.Volume >= FsGetSettings().MinVolume,
            "Volume is too small"
        );
        FsSetting memory fsSetting = FsGetSettings();
        fsNodeInfo.Pledge =
            fsSetting.FsGasPrice *
            fsSetting.GasPerGBPerBlock *
            fsNodeInfo.Volume;
        fsNodeInfo.Profit = 0;
        fsNodeInfo.RestVol = fsNodeInfo.Volume;
        nodeInfo.push(fsNodeInfo);
        nodeList.AddrList.push(fsNodeInfo.WalletAddr);
        emit RegisterNodeEvent(FsEvent.EVENT_FS_REG_NODE);
    }
}
