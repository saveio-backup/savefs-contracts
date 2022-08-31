//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract Dns is Initializable, IFsEvent {
    address admin;
    mapping (bytes=>HeaderInfo) headers;
    function DnsInit(address _admin) public {
        admin = _admin;
        HeaderInfo memory info;
        info.Header = "dsp";
        info.HeaderOwner = _admin;
        info.Desc = "reserved dsp protocol";
        info.BlockHeight = 0;
        info.TTL = 0;
        headers[info.Header] = info;
    }

    function RegisterName() public {

    }

    function GetName(ReqInfo memory req) public view returns (bytes memory) {

    }
}
