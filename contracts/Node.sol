//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./FileSystem.sol";
import "./Config.sol";
import "./Sector.sol";

contract Node is Initializable {
    Config config;
    Sector sector;

    mapping(address => NodeInfo) nodesInfo; // walletAddr => NodeInfo
    address[] nodeList; // nodeAddr list

    event RegisterNodeEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        address nodeAddr,
        uint64 volume,
        uint64 serviceTime
    );

    event UnRegisterNodeEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr
    );

    modifier VolumeRequire(NodeInfo memory nodeInfo, Setting memory setting) {
        require(nodeInfo.Volume >= setting.MinVolume, "Volume is too small");
        _;
    }

    modifier NodeRegisted(address walletAddr) {
        require(nodesInfo[walletAddr].Volume != 0, "Node not registered");
        _;
    }

    modifier NodeNotRegisted(address walletAddr) {
        require(nodesInfo[walletAddr].Volume == 0, "Node already registered");
        _;
    }

    error NotEnoughPledge(uint256 got, uint256 want);
    error ZeroProfit();

    function initialize(Config _config, Sector _sector) public initializer {
        config = _config;
        sector = _sector;
    }

    function CalculateNodePledge(NodeInfo memory nodeInfo)
        public
        view
        returns (uint64)
    {
        Setting memory setting = config.GetSetting();
        return setting.GasPrice * setting.GasPerGBPerBlock * nodeInfo.Volume;
    }

    function Register(NodeInfo memory nodeInfo)
        public
        payable
        VolumeRequire(nodeInfo, config.GetSetting())
        NodeNotRegisted(nodeInfo.WalletAddr)
    {
        uint64 pledge = CalculateNodePledge(nodeInfo);
        if (msg.value < pledge) {
            revert NotEnoughPledge(msg.value, pledge);
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

    function NodeUpdate(NodeInfo memory nodeInfo)
        public
        payable
        VolumeRequire(nodeInfo, config.GetSetting())
        NodeRegisted(nodeInfo.WalletAddr)
    {
        require(
            nodesInfo[nodeInfo.WalletAddr].WalletAddr == nodeInfo.WalletAddr,
            "Node walletAddr changed"
        );
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

    function NodeCancel(address walletAddr) public NodeRegisted(walletAddr) {
        NodeInfo memory nodeInfo = nodesInfo[walletAddr];
        if (nodeInfo.Pledge > 0) {
            payable(nodeInfo.WalletAddr).transfer(
                nodeInfo.Pledge + nodeInfo.Profit
            );
        }
        delete nodesInfo[walletAddr];
        NodeListRemove(walletAddr);
        SectorInfo[] memory sectorInfos = sector.GetSectorsForNode(
            nodeInfo.NodeAddr
        );
        for (uint256 i = 0; i < sectorInfos.length; i++) {
            sector.DeleteSecotr(
                SectorRef({
                    NodeAddr: nodeInfo.NodeAddr,
                    SectorId: sectorInfos[i].SectorID
                })
            );
        }
        emit UnRegisterNodeEvent(FsEvent.UN_REG_NODE, block.number, walletAddr);
    }

    function NodeListRemove(address addr) private {
        for (uint256 i = 0; i < nodeList.length; i++) {
            if (nodeList[i] == addr) {
                delete nodeList[i];
                return;
            }
        }
    }

    /**
     * description: Actually I can't understand why do this
     *              mixed use nodeAddr and walletAddr
     *
     * @return nodeAddr => NodeInfo
     */
    function GetNodeList() public view returns (NodeInfo[] memory) {
        NodeInfo[] memory _nodesInfo = new NodeInfo[](nodeList.length);
        for (uint256 i = 0; i < nodeList.length; i++) {
            _nodesInfo[i] = nodesInfo[nodeList[i]];
        }
        return _nodesInfo;
    }

    function GetNodeInfoByWalletAddr(address walletAddr)
        public
        view
        returns (NodeInfo memory)
    {
        return nodesInfo[walletAddr];
    }

    function GetNodeInfoByNodeAddr(address nodeAddr)
        public
        view
        returns (NodeInfo memory)
    {
        return nodesInfo[nodeAddr];
    }

    function WithDrawProfit(address walletAddr)
        public
        NodeRegisted(walletAddr)
    {
        NodeInfo memory nodeInfo = nodesInfo[walletAddr];
        if (nodeInfo.Profit > 0) {
            payable(nodeInfo.WalletAddr).transfer(nodeInfo.Profit);
            nodeInfo.Profit = 0;
        } else {
            revert ZeroProfit();
        }
        nodesInfo[walletAddr] = nodeInfo;
    }

    function IsNodeRegisted(address walletAddr) public view returns (bool) {
        return nodesInfo[walletAddr].Volume != 0;
    }

    function UpdateNodeInfo(NodeInfo memory nodeInfo) public payable {
        nodesInfo[nodeInfo.WalletAddr] = nodeInfo;
    }
}
