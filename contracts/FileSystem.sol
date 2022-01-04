//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Struct.sol";
import "./Error.sol";
import "./Util.sol";

contract FileSystem is Initializable {
    FsNodeInfo[] nodeInfo;
    NodeList nodeList;

    function initialize() public initializer {
        console.log("initializer");
    }

    function FsGetSettings() public pure returns (FsSetting memory) {
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

    function FsNodeRegister(FsNodeInfo memory fsNodeInfo) public payable {
        require(
            fsNodeInfo.Volume >= FsGetSettings().MinVolume,
            "Volume is too small"
        );
        FsSetting memory fsSetting = FsGetSettings();
        uint64 pledge = calculateNodePledge(fsNodeInfo, fsSetting);
        fsNodeInfo.Pledge = pledge;
        fsNodeInfo.Profit = 0;
        fsNodeInfo.RestVol = fsNodeInfo.Volume;
        nodeInfo.push(fsNodeInfo);
        nodeList.AddrList.push(fsNodeInfo.WalletAddr);
    }

    function calculateNodePledge(
        FsNodeInfo memory fsNodeInfo,
        FsSetting memory fsSetting
    ) public pure returns (uint64) {
        return
            fsSetting.FsGasPrice *
            fsSetting.GasPerGBPerBlock *
            fsNodeInfo.Volume;
    }
}
