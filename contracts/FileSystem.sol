//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./IFileSystem.sol";
import "./Struct.sol";
import "./Error.sol";
import "./Enum.sol";

contract FileSystem is Initializable, IFileSystem {
    mapping(address => FsNodeInfo) nodesInfo;
    NodeList nodeList;
    mapping(address => SectorInfos) sectorInfos;

    /**********************************************************************
     * event define start *************************************************
     */
    event StoreFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        uint64 fileSize,
        address walletAddr,
        uint64 cost,
        bool isPlotFile
    );

    event DeleteFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        address walletAddr
    );

    event DeleteFilesEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        address walletAddr
    );

    event SetUserSpaceEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sizeType,
        uint64 size,
        uint64 countType,
        uint64 count
    );

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

    event ProveFileEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 profit
    );

    event FilePDPSuccessEvent(
        FsEvent eventType,
        uint256 blockHeight,
        string fileHash,
        address walletAddr
    );

    event CreateSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId,
        uint64 provLevel,
        uint64 size,
        bool isPlots
    );

    event DeleteSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId
    );
    /**
     * event define end ****************************************************
     ***********************************************************************/

    /************************************************************************
     * modifier define start ************************************************
     */
    modifier VolumeRequire(FsNodeInfo memory fsNodeInfo) {
        require(
            fsNodeInfo.Volume >= FsGetSettings().MinVolume,
            "Volume is too small"
        );
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

    /**
     * modifier define end ***************************************************
     *************************************************************************/

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

    /****************************************************************************
     * Node info mamanagement start *********************************************
     */
    function FsNodeRegister(FsNodeInfo memory fsNodeInfo)
        public
        payable
        override
        VolumeRequire(fsNodeInfo)
        NodeNotRegisted(fsNodeInfo.WalletAddr)
    {
        uint64 pledge = CalcLateNodePledge(fsNodeInfo);
        if (msg.value < pledge) {
            revert NotEnoughPledge(msg.value, pledge);
        }
        fsNodeInfo.Pledge = pledge;
        fsNodeInfo.Profit = 0;
        fsNodeInfo.RestVol = fsNodeInfo.Volume;
        nodesInfo[fsNodeInfo.WalletAddr] = fsNodeInfo;
        nodeList.AddrList.push(fsNodeInfo.WalletAddr);
        emit RegisterNodeEvent(
            FsEvent.EVENT_FS_REG_NODE,
            block.number,
            fsNodeInfo.WalletAddr,
            fsNodeInfo.NodeAddr,
            fsNodeInfo.Volume,
            fsNodeInfo.ServiceTime
        );
    }

    function CalcLateNodePledge(FsNodeInfo memory fsNodeInfo)
        public
        pure
        override
        returns (uint64)
    {
        FsSetting memory fsSetting = FsGetSettings();
        return
            fsSetting.FsGasPrice *
            fsSetting.GasPerGBPerBlock *
            fsNodeInfo.Volume;
    }

    function FsNodeQuery(address walletAddr)
        public
        view
        override
        returns (FsNodeInfo memory)
    {
        return nodesInfo[walletAddr];
    }

    function FsNodeUpdate(FsNodeInfo memory fsNodeInfo)
        public
        payable
        override
        VolumeRequire(fsNodeInfo)
        NodeRegisted(fsNodeInfo.WalletAddr)
    {
        require(
            nodesInfo[fsNodeInfo.WalletAddr].WalletAddr ==
                fsNodeInfo.WalletAddr,
            "Node walletAddr changed"
        );
        FsNodeInfo memory oldNode = nodesInfo[fsNodeInfo.WalletAddr];
        uint64 newPledge = CalcLateNodePledge(fsNodeInfo);
        uint64 oldPledge = oldNode.Pledge;
        if (newPledge < oldPledge) {
            payable(fsNodeInfo.WalletAddr).transfer(oldPledge - newPledge);
        } else {
            uint64 pledge = newPledge - oldPledge;
            if (msg.value < pledge) {
                revert NotEnoughPledge(msg.value, pledge);
            }
        }
        fsNodeInfo.Pledge = newPledge;
        fsNodeInfo.Profit = oldNode.Profit;
        fsNodeInfo.RestVol =
            oldNode.RestVol +
            fsNodeInfo.Volume -
            oldNode.Volume;
        nodesInfo[fsNodeInfo.WalletAddr] = fsNodeInfo;
    }

    function FsNodeCancel(address walletAddr)
        public
        override
        NodeRegisted(walletAddr)
    {
        FsNodeInfo memory fsNodeInfo = nodesInfo[walletAddr];
        if (fsNodeInfo.Pledge > 0) {
            payable(fsNodeInfo.WalletAddr).transfer(fsNodeInfo.Pledge + fsNodeInfo.Profit);
        }
        delete sectorInfos[fsNodeInfo.NodeAddr];
        delete nodesInfo[walletAddr];
        NodeListRemove(walletAddr);
        emit UnRegisterNodeEvent(
            FsEvent.EVENT_FS_UN_REG_NODE,
            block.number,
            walletAddr
        );
    }

    function NodeListRemove(address addr) public {
        for (uint256 i = 0; i < nodeList.AddrList.length - 1; i++) {
            if (nodeList.AddrList[i] == addr) {
                delete nodeList.AddrList[i];
                return;
            }
        }
    }

    function FsGetNodeList() public view returns (FsNodesInfo memory) {
        FsNodesInfo memory fsNodesInfo;
        fsNodesInfo.NodesInfo = new FsNodeInfo[](nodeList.AddrList.length);
        for (uint256 i = 0; i < nodeList.AddrList.length; i++) {
            FsNodeInfo memory info = nodesInfo[nodeList.AddrList[i]];
            fsNodesInfo.NodesInfo[i] = info;
            fsNodesInfo.NodeNum++;
        }
        return fsNodesInfo;
    }
    /**
     * Node info mamanagement end ************************************************
     *****************************************************************************/
}
