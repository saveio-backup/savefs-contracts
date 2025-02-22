//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract Node is Initializable, INode, IFsEvent {
    IConfig config;
    ISector sector;

    mapping(address => NodeInfo) nodesInfo; // walletAddr => NodeInfo
    address[] nodeList; // nodeAddr list

    function initialize(IConfig _config, ISector _sector) public initializer {
        config = _config;
        sector = _sector;
    }

    function CalculateNodePledge(NodeInfo memory nodeInfo)
        public
        view
        virtual
        override
        returns (uint64)
    {
        Setting memory setting = config.GetSetting();
        return setting.GasPrice * setting.GasPerGBPerBlock * nodeInfo.Volume;
    }

    function Register(NodeInfo memory nodeInfo)
        public
        payable
        virtual
        override
    {
        Setting memory setting = config.GetSetting();
        if (nodeInfo.Volume < setting.MinVolume) {
            emit FsError("Register", "Volume is too small");
            return;
        }
        if (nodesInfo[nodeInfo.WalletAddr].Volume != 0) {
            emit FsError("Register", "Node already registered");
            return;
        }
        uint64 pledge = CalculateNodePledge(nodeInfo);
        if (msg.value < pledge) {
            emit FsError("Register", "Insufficient pledge");
            return;
        }
        nodeInfo.Pledge = pledge;
        nodeInfo.Profit = 0;
        nodeInfo.RestVol = nodeInfo.Volume;
        nodesInfo[nodeInfo.WalletAddr] = nodeInfo;
        nodeList.push(nodeInfo.WalletAddr);
        emit RegisterNodeEvent(
            FsEvent.REG_NODE,
            block.number,
            nodeInfo.WalletAddr,
            nodeInfo.NodeAddr,
            nodeInfo.Volume,
            nodeInfo.ServiceTime
        );
    }

    function GetPledgeUpdate(NodeInfo memory nodeInfo)
        public
        view
        returns (int64)
    {
        NodeInfo memory oldNode = nodesInfo[nodeInfo.WalletAddr];
        uint64 newPledge = CalculateNodePledge(nodeInfo);
        uint64 oldPledge = oldNode.Pledge;
        return int64(newPledge) - int64(oldPledge);
    }

    function NodeUpdate(NodeInfo memory nodeInfo)
        public
        payable
        virtual
        override
    {
        require(
            nodesInfo[nodeInfo.WalletAddr].WalletAddr == nodeInfo.WalletAddr,
            "Node walletAddr changed"
        );
        Setting memory setting = config.GetSetting();
        if (nodeInfo.Volume < setting.MinVolume) {
            emit FsError("NodeUpdate", "Volume is too small");
            return;
        }
        if (nodesInfo[nodeInfo.WalletAddr].Volume == 0) {
            emit FsError("NodeUpdate", "Node not registered");
            return;
        }
        NodeInfo memory oldNode = nodesInfo[nodeInfo.WalletAddr];
        uint64 newPledge = CalculateNodePledge(nodeInfo);
        uint64 oldPledge = oldNode.Pledge;
        TransferState memory state;
        if (newPledge < oldPledge) {
            state.From = address(this);
            state.To = nodeInfo.WalletAddr;
            state.Value = oldNode.Pledge - newPledge;
        } else if (newPledge > oldPledge) {
            state.From = nodeInfo.WalletAddr;
            state.To = address(this);
            state.Value = newPledge - oldNode.Pledge;
        }
        if (newPledge != oldPledge) {
            if (state.To == address(this)) {
                require(msg.value >= state.Value, "Not enough pledge");
            } else {
                payable(state.From).transfer(state.Value);
            }
        }
        nodeInfo.Pledge = newPledge;
        nodeInfo.Profit = oldNode.Profit;
        nodeInfo.RestVol = oldNode.RestVol + nodeInfo.Volume - oldNode.Volume;
        nodesInfo[nodeInfo.WalletAddr] = nodeInfo;
    }

    function Cancel(address walletAddr) public virtual override {
        if (nodesInfo[walletAddr].Volume == 0) {
            emit FsError("Cancel", "Node not registered");
            return;
        }
        NodeInfo memory nodeInfo = nodesInfo[walletAddr];
        if (nodeInfo.Pledge > 0) {
            payable(nodeInfo.WalletAddr).transfer(
                nodeInfo.Pledge + nodeInfo.Profit
            );
        }
        delete nodesInfo[walletAddr];
        NodeListRemove(walletAddr);
        SectorInfo[] memory sectorInfos = sector.GetSectorsForNode(
            nodeInfo.WalletAddr
        );
        for (uint256 i = 0; i < sectorInfos.length; i++) {
            sector.DeleteSector(
                SectorRef({
                    NodeAddr: nodeInfo.WalletAddr,
                    SectorId: sectorInfos[i].SectorID
                })
            );
        }
        emit UnRegisterNodeEvent(FsEvent.UN_REG_NODE, block.number, walletAddr);
    }

    function GetNodeList()
        public
        view
        virtual
        override
        returns (NodeInfo[] memory)
    {
        NodeInfo[] memory _nodesInfo = new NodeInfo[](nodeList.length);
        for (uint256 i = 0; i < nodeList.length; i++) {
            _nodesInfo[i] = nodesInfo[nodeList[i]];
        }
        return _nodesInfo;
    }

    function GetNodeInfoByWalletAddr(address walletAddr)
        public
        view
        virtual
        override
        returns (NodeInfo memory)
    {
        return nodesInfo[walletAddr];
    }

    function GetNodeInfoByNodeAddr(bytes memory nodeAddr)
        public
        view
        virtual
        override
        returns (NodeInfo memory)
    {
        NodeInfo memory node;
        for (uint256 i = 0; i < nodeList.length; i++) {
            node = nodesInfo[nodeList[i]];
            if (keccak256(node.NodeAddr) == keccak256(nodeAddr)) {
                return node;
            }
        }
        return node;
    }

    function WithDrawProfit(address walletAddr) public virtual override {
        if (nodesInfo[walletAddr].Volume == 0) {
            emit FsError("WithDrawProfit", "Node not registered");
            return;
        }
        NodeInfo memory nodeInfo = nodesInfo[walletAddr];
        if (nodeInfo.Profit > 0) {
            payable(nodeInfo.WalletAddr).transfer(nodeInfo.Profit);
            nodeInfo.Profit = 0;
        } else {
            emit FsError("WithDrawProfit", "Zero profit");
            return;
        }
        nodesInfo[walletAddr] = nodeInfo;
    }

    function IsNodeRegisted(address walletAddr)
        public
        view
        virtual
        override
        returns (bool)
    {
        return nodesInfo[walletAddr].Volume != 0;
    }

    function UpdateNodeInfo(NodeInfo memory nodeInfo)
        public
        payable
        virtual
        override
    {
        nodesInfo[nodeInfo.WalletAddr] = nodeInfo;
    }

    function NodeListRemove(address addr) private {
        for (uint256 i = 0; i < nodeList.length; i++) {
            if (nodeList[i] == addr) {
                delete nodeList[i];
                return;
            }
        }
    }
}
