//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract FileSystem is Initializable {

  function initialize() public initializer {
    console.log("initialize");
  }

  struct FsSetting {
    int FsGasPrice;
    int GasPerGBPerBlock;
    int GasPerKBPerBlock;
    int GasForChallenge;
    int MaxProveBlocks;
    int MinVolume;
    int DefaultProvePeriod;
    int DefaultProveLevel;
    int DefaultCopyNum;
  }

  function GetSettings() public pure returns (FsSetting memory) {
    FsSetting memory fsSetting;
    fsSetting.FsGasPrice = 1;
    fsSetting.GasPerGBPerBlock = 1;
    fsSetting.GasPerKBPerBlock = 1;
    fsSetting.GasForChallenge = 200000;
    fsSetting.MaxProveBlocks = 32;
    fsSetting.MinVolume = 1000 * 1000;
    fsSetting.DefaultProvePeriod = 3600 * 24 / 5;
    fsSetting.DefaultProveLevel = 1;
    fsSetting.DefaultCopyNum = 2;
    return fsSetting;
  }
}