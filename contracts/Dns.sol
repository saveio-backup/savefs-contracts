//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract Dns is Initializable, IFsEvent {
    uint64 SYSTEM = 0x00;
    uint64 CUSTOM_HEADER = 0x01;
    uint64 CUSTOM_URL = 0x02;
    uint64 CUSTOM_HEADER_URL = 0x04;
    uint64 UPDATE = 0x08;

    bytes DSP_HEADER = "dsp";
    bytes DSP_PLUGIN_HEADER = "dsp-plugin";

    address admin;
    mapping(bytes => HeaderInfo) headers; // header => HeaderInfo
    mapping(bytes => NameInfo) nameInfos; // header + url => NameInfo
    mapping(bytes => bool) pluginListKey; // header + url => any

    function DnsInit(address _admin) public {
        admin = _admin;
        HeaderInfo memory info;
        info.Header = DSP_HEADER;
        info.HeaderOwner = _admin;
        info.Desc = "reserved dsp protocol";
        info.BlockHeight = 0;
        info.TTL = 0;
        headers[info.Header] = info;
    }

    function RegisterName(RequestName memory req) public payable {
        require(req.Name.length > 4);
        NameInfo memory info;
        if (req.Type == SYSTEM) {
            info.Type = uint64(NameType.Normal);
            info.Header = DSP_HEADER;
            info.URL = CreateDefaultUrl(req.Name);
            info.Name = req.Name;
            info.NameOwner = req.NameOwner;
            info.Desc = req.Desc;
            info.BlockHeight = block.number + 1;
            info.TTL = req.DesireTTL;
        }
        if (req.Type == CUSTOM_HEADER) {
            info.Type = uint64(NameType.Normal);
            info.Header = req.Header;
            info.URL = CreateDefaultUrl(req.Name);
            info.Name = req.Name;
            info.NameOwner = req.NameOwner;
            info.Desc = req.Desc;
            info.BlockHeight = block.number + 1;
            info.TTL = req.DesireTTL;
        }
        if (req.Type == CUSTOM_URL) {
            info.Type = uint64(NameType.Normal);
            info.Header = DSP_HEADER;
            info.URL = req.URL;
            info.Name = req.Name;
            info.NameOwner = req.NameOwner;
            info.Desc = req.Desc;
            info.BlockHeight = block.number + 1;
            info.TTL = req.DesireTTL;
        }
        if (req.Type == CUSTOM_HEADER_URL) {
            info.Type = uint64(NameType.Normal);
            info.Header = req.Header;
            info.URL = req.URL;
            info.Name = req.Name;
            info.NameOwner = req.NameOwner;
            info.Desc = req.Desc;
            info.BlockHeight = block.number + 1;
            info.TTL = req.DesireTTL;
        }
        bytes memory key = concat(info.Header, info.URL);
        nameInfos[key] = info;
        if (
            req.Type == CUSTOM_HEADER_URL &&
            keccak256(req.Header) == keccak256(DSP_PLUGIN_HEADER)
        ) {
            pluginListKey[key] = true;
        }
        emit NotifyNameInfoAdd(
            req.NameOwner,
            GetUrl(info.Header, info.URL),
            info
        );
    }

    function GetUrl(bytes memory header, bytes memory url)
        public
        pure
        returns (bytes memory)
    {
        bytes memory a = concat(header, "://");
        bytes memory b = concat(a, url);
        return b;
    }

    function CreateDefaultUrl(bytes memory name)
        public
        pure
        returns (bytes memory)
    {
        return toBytes(sha256(name));
    }

    function concat(bytes memory b1, bytes memory b2)
        public
        pure
        returns (bytes memory)
    {
        uint256 b1l = b1.length;
        uint256 b2l = b2.length;
        bytes memory result = new bytes(b1.length + b2.length);
        assembly {
            mstore(add(result, b1l), b1)
            mstore(add(result, b2l), b2)
        }
        return result;
    }

    function toBytes(bytes32 _data) public pure returns (bytes memory) {
        return abi.encodePacked(_data);
    }

    function RegisterHeader() public {}

    function TransferName() public {}

    function TransferHeader() public {}

    function GetName(ReqInfo memory req) public view returns (bytes memory) {}
}
