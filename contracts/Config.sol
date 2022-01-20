//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";

contract Config is Initializable {
    function GetSetting() public pure returns (Setting memory) {
        Setting memory setting;
        setting.GasPrice = 1;
        setting.GasPerGBPerBlock = 1;
        setting.GasPerKBPerBlock = 1;
        setting.GasForChallenge = 200000;
        setting.MaxProveBlockNum = 32;
        setting.MinVolume = 1000 * 1000;
        setting.DefaultProvePeriod = (3600 * 24) / 5;
        setting.DefaultProveLevel = 1;
        setting.DefaultCopyNum = 2;
        setting.DefaultBlockInterval = 5;
        setting.MinSectorSize = 1000 * 1000;
        return setting;
    }

    function GetProveIntervalByProveLevel(ProveLevel proveLevel) public pure returns(uint64) {
        Setting memory setting = GetSetting();
        if (proveLevel == ProveLevel.HIGH) {
            return setting.DefaultProvePeriod;
        }
        if (proveLevel == ProveLevel.MEDIEUM) {
            return setting.DefaultProvePeriod * 2;
        }
        if (proveLevel == ProveLevel.LOW) {
            return setting.DefaultProvePeriod * 8;
        }
        return setting.DefaultProvePeriod;
    }

    function GetSettingWithProveLevel(ProveLevel proveLevel)
        public
        pure
        returns (Setting memory)
    {
        Setting memory setting = GetSetting();
        setting.DefaultProvePeriod = GetProveIntervalByProveLevel(proveLevel);
        return setting;
    }
}
