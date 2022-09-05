//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract Dns is Initializable, IFsEvent {
    uint64 SYSTEM;
    uint64 CUSTOM_HEADER;
    uint64 CUSTOM_URL;
    uint64 CUSTOM_HEADER_URL;
    uint64 UPDATE;

    bytes DSP_HEADER;
    bytes DSP_PLUGIN_HEADER;

    address admin;
    mapping(bytes => HeaderInfo) headerInfos; // header => HeaderInfo
    mapping(bytes => NameInfo) nameInfos; // header + url => NameInfo
    mapping(bytes => bool) pluginListKey; // header + url => any
    mapping(string => PeerPoolItem) peerPool; // peerPubKey => PeerPoolItem
    mapping(address => DNSNodeInfo) dnsNodeInfos; // walletAddr => DNSNodeInfo

    function initialize() public initializer {
        SYSTEM = 0x00;
        CUSTOM_HEADER = 0x01;
        CUSTOM_URL = 0x02;
        CUSTOM_HEADER_URL = 0x04;
        UPDATE = 0x08;

        DSP_HEADER = "dsp";
        DSP_PLUGIN_HEADER = "dsp-plugin";
    }

    // dns
    function DnsInit(address _admin) public {
        admin = _admin;
        HeaderInfo memory info;
        info.Header = DSP_HEADER;
        info.HeaderOwner = _admin;
        info.Desc = "reserved dsp protocol";
        info.BlockHeight = 0;
        info.TTL = 0;
        headerInfos[info.Header] = info;
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

    function RegisterHeader(RequestHeader memory req) public payable {
        HeaderInfo memory info;
        info.Header = req.Header;
        info.HeaderOwner = req.NameOwner;
        info.Desc = req.Desc;
        info.BlockHeight = block.number + 1;
        info.TTL = 0;
        headerInfos[req.Header] = info;
        emit NotifyHeaderAdd(req.NameOwner, req.Header);
    }

    function TransferName(TransferInfo memory info) public {
        bytes memory key = concat(info.Header, info.URL);
        NameInfo memory nameInfo = nameInfos[key];
        if (nameInfo.NameOwner != info.From) {
            revert("not owner");
        }
        nameInfos[key].NameOwner = info.To;
        nameInfos[key].BlockHeight = block.number + 1;
        nameInfos[key].TTL = uint64(
            nameInfo.TTL + nameInfo.BlockHeight - block.number
        );
        emit NotifyNameInfoTransfer(
            info.From,
            info.To,
            GetUrl(info.Header, info.URL)
        );
    }

    function TransferHeader(TransferInfo memory info) public {
        HeaderInfo memory headerInfo = headerInfos[info.Header];
        if (headerInfo.HeaderOwner != info.From) {
            revert("not owner");
        }
        uint64 ttl;
        if (headerInfo.TTL + headerInfo.BlockHeight <= block.number) {
            ttl = 0;
        } else {
            ttl = uint64(
                headerInfo.TTL + headerInfo.BlockHeight - block.number
            );
        }
        headerInfos[info.Header].Header = info.Header;
        headerInfos[info.Header].HeaderOwner = info.To;
        headerInfos[info.Header].BlockHeight = block.number + 1;
        headerInfos[info.Header].TTL = ttl;
        emit NotifyHeaderTransfer(info.From, info.To, info.Header);
    }

    function UpdateName(RequestName memory req) public payable {
        bytes memory key = concat(req.Header, req.URL);
        NameInfo memory nameInfo = nameInfos[key];
        if (nameInfo.NameOwner != msg.sender) {
            revert("not owner");
        }
        nameInfos[key].Type = req.Type;
        nameInfos[key].Name = req.Name;
        nameInfos[key].Desc = req.Desc;
        nameInfos[key].TTL = req.DesireTTL;
        nameInfos[key].BlockHeight = block.number + 1;
        emit NotifyNameInfoChange(req.NameOwner, GetUrl(req.Header, req.URL));
    }

    function GetName(ReqInfo memory req) public view returns (NameInfo memory) {
        bytes memory key = concat(req.Header, req.URL);
        return nameInfos[key];
    }

    function GetHeader(ReqInfo memory req)
        public
        view
        returns (HeaderInfo memory)
    {
        return headerInfos[req.Header];
    }

    function DelDNS(ReqInfo memory req) public payable {
        bytes memory key = concat(req.Header, req.URL);
        NameInfo memory nameInfo = nameInfos[key];
        if (nameInfo.NameOwner != admin) {
            revert("not admin");
        }
        delete nameInfos[key];
    }

    function DelHeader(ReqInfo memory req) public payable {
        HeaderInfo memory headerInfo = headerInfos[req.Header];
        if (headerInfo.HeaderOwner != admin) {
            revert("not admin");
        }
        delete headerInfos[req.Header];
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

    // dns node
    function DNSNodeReg(DNSNodeInfo memory info) public payable {
        require(info.InitDeposit > 0, "index must > 0");
        require(msg.value >= info.InitDeposit, "deposit must > 0");
        if (info.WalletAddr != msg.sender) {
            revert("not owner");
        }
        PeerPoolItem memory item;
        item.PeerPubKey = info.PeerPubKey;
        item.WalletAddress = info.WalletAddr;
        item.Status = uint8(DNSStatus.RegisterCandidateStatus);
        item.TotalInitPos = info.InitDeposit;
        peerPool[info.PeerPubKey] = item;
        dnsNodeInfos[info.WalletAddr] = info;
        emit DNSNodeRegister(
            info.IP,
            info.Port,
            info.WalletAddr,
            info.InitDeposit
        );
    }

    function UnRegDNSNode(UnRegisterCandidateParam memory req) public {
        if (req.Address != msg.sender) {
            revert("not owner");
        }
        PeerPoolItem memory item = peerPool[req.PeerPubKey];
        if (item.Status != uint8(DNSStatus.RegisterCandidateStatus)) {
            revert("not register");
        }
        if (item.WalletAddress != req.Address) {
            revert("not owner");
        }
        payable(req.Address).transfer(item.TotalInitPos);
        delete peerPool[req.PeerPubKey];
        delete dnsNodeInfos[req.Address];
        emit DNSNodeUnReg(req.Address);
    }

    function ApproveDNSCandidate(string memory peerPubKey) public {
        PeerPoolItem memory item = peerPool[peerPubKey];
        if (item.TotalInitPos < 1) {
            revert("not enough init deposit");
        }
        if (item.Status != uint8(DNSStatus.RegisterCandidateStatus)) {
            revert("not register");
        }
        item.Status = uint8(DNSStatus.ConsensusStatus);
        peerPool[peerPubKey] = item;
    }

    function RejectDNSCandidate(string memory peerPubKey) public payable {
        PeerPoolItem memory item = peerPool[peerPubKey];
        if (item.Status != uint8(DNSStatus.RegisterCandidateStatus)) {
            revert("not register");
        }
        payable(item.WalletAddress).transfer(item.TotalInitPos);
        delete peerPool[peerPubKey];
        delete dnsNodeInfos[item.WalletAddress];
    }

    function QuitNode(QuitNodeParam memory req) public {
        if (req.Address != msg.sender) {
            revert("not owner");
        }
        PeerPoolItem memory item = peerPool[req.PeerPubKey];
        if (item.WalletAddress != req.Address) {
            revert("not owner");
        }
        if (
            item.Status != uint8(DNSStatus.ConsensusStatus) &&
            item.Status != uint8(DNSStatus.RegisterCandidateStatus)
        ) {
            revert("not consensus");
        }
        if (item.Status == uint8(DNSStatus.ConsensusStatus)) {
            item.Status = uint8(DNSStatus.QuitConsensusStatus);
        } else {
            item.Status = uint8(DNSStatus.QuitingStatus);
        }
        peerPool[req.PeerPubKey] = item;
        delete dnsNodeInfos[req.Address];
    }

    function AddInitPos(ChangeInitPosParam memory req) public payable {
        require(req.Pos > 0, "pos must > 0");
        require(msg.value >= req.Pos, "deposit must > 0");
        if (req.Address != msg.sender) {
            revert("not owner");
        }
        PeerPoolItem memory item = peerPool[req.PeerPubKey];
        if (item.WalletAddress != req.Address) {
            revert("not owner");
        }
        if (
            item.Status != uint8(DNSStatus.ConsensusStatus) &&
            item.Status != uint8(DNSStatus.RegisterCandidateStatus)
        ) {
            revert("not consensus");
        }
        item.TotalInitPos += req.Pos;
        peerPool[req.PeerPubKey] = item;
    }

    function ReduceInitPos(ChangeInitPosParam memory req) public payable {
        require(req.Pos > 0, "pos must > 0");
        if (req.Address != msg.sender) {
            revert("not owner");
        }
        PeerPoolItem memory item = peerPool[req.PeerPubKey];
        if (item.WalletAddress != req.Address) {
            revert("not owner");
        }
        if (item.TotalInitPos < req.Pos) {
            revert("not enough init deposit");
        }
        item.TotalInitPos -= req.Pos;
        peerPool[req.PeerPubKey] = item;
    }

    function GetPeerPoolMap() public pure returns (PeerPoolItem[] memory) {
        PeerPoolItem[] memory items = new PeerPoolItem[](0);
        // TODO
        return items;
    }

    function GetPeerPoolItem(string memory peerPubKey)
        public
        view
        returns (PeerPoolItem memory)
    {
        return peerPool[peerPubKey];
    }

    function GetDNSNodeByAddress(address addr)
        public
        view
        returns (DNSNodeInfo memory)
    {
        return dnsNodeInfos[addr];
    }

    function GetAllDnsNodes() public pure returns (DNSNodeInfo[] memory) {
        DNSNodeInfo[] memory items = new DNSNodeInfo[](0);
        // TODO
        return items;
    }

    function UpdateDNSNodesInfo(UpdateNodeParam memory info) public {
        require(info.IP.length > 0, "ip must not empty");
        require(info.Port.length > 0, "port must > 0");
        if (dnsNodeInfos[msg.sender].WalletAddr != msg.sender) {
            revert("not register");
        }
        dnsNodeInfos[msg.sender].IP = info.IP;
        dnsNodeInfos[msg.sender].Port = info.Port;
        emit DNSNodeRegister(
            info.IP,
            info.Port,
            msg.sender,
            dnsNodeInfos[msg.sender].InitDeposit
        );
    }

    function GetPluginList() public pure returns (NameInfo[] memory) {
        NameInfo[] memory items = new NameInfo[](0);
        // TODO
        return items;
    }
}
