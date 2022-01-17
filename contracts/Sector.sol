//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./Node.sol";

contract Sector is Initializable {
    Node node;

    mapping(address => SectorInfo[]) sectorInfos; // nodeAddr => SectorInfo[]
    mapping(address => mapping(uint64 => uint64)) sectorFileInfoGroup; // nodeAddr => sectorId => groupId

    event DeleteSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId
    );

    error NotEnoughVolume(uint64 got, uint64 want);
    error NotEmptySector(uint64 got, uint64 want);

    function initialize(Node _node) public initializer {
        node = _node;
    }

    function getSectorTotalSizeForNode(address nodeAddr)
        public
        view
        returns (uint64)
    {
        uint64 totalSize = 0;
        SectorInfo[] memory nodeSectorInfos = sectorInfos[nodeAddr];
        for (uint256 i = 0; i < nodeSectorInfos.length; i++) {
            totalSize += nodeSectorInfos[i].Size;
        }
        return totalSize;
    }

    function CreateSector(SectorInfo memory sectorInfo) public payable {
        require(sectorInfo.SectorID > 0, "sectorId is wrong");
        require(sectorInfo.Size > 0, "sector size is wrong");
        require(
            node.IsNodeRegisted(sectorInfo.NodeAddr),
            "node not registered"
        );
        require(
            GetSectorInfo(
                SectorRef({
                    SectorId: sectorInfo.SectorID,
                    NodeAddr: sectorInfo.NodeAddr
                })
            ).Size == 0,
            "sector already exists"
        );
        NodeInfo memory nodeInfo = node.GetNodeInfoByNodeAddr(
            sectorInfo.NodeAddr
        );
        uint64 totalSize = getSectorTotalSizeForNode(sectorInfo.NodeAddr);
        if (sectorInfo.Size + totalSize > nodeInfo.Volume) {
            revert NotEnoughVolume(
                nodeInfo.Volume,
                sectorInfo.Size + totalSize
            );
        }
        sectorInfos[sectorInfo.NodeAddr].push(sectorInfo);
    }

    function GetSectorInfo(SectorRef memory sectorRef)
        public
        view
        returns (SectorInfo memory)
    {
        require(sectorRef.SectorId > 0, "sectorId must be greater than 0");
        SectorInfo[] memory sectorInfo = sectorInfos[sectorRef.NodeAddr];
        for (uint64 i = 0; i < sectorInfo.length; i++) {
            if (sectorInfo[i].SectorID == sectorRef.SectorId) {
                return sectorInfo[i];
            }
        }
        SectorInfo memory emptySectorInfo;
        return emptySectorInfo;
    }

    function GetSectorInfos(address nodeAddr)
        public
        view
        returns (SectorInfo[] memory)
    {
        return sectorInfos[nodeAddr];
    }

    function getSectorFileInfoGroupNum(address nodeAddr, uint64 sectorId)
        public
        view
        returns (uint64)
    {
        SectorInfo memory sectorInfo = GetSectorInfo(
            SectorRef({SectorId: sectorId, NodeAddr: nodeAddr})
        );
        return sectorInfo.GroupNum;
    }

    function deleteSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorId,
        uint64 groupId
    ) public {
        mapping(uint64 => uint64)
            storage _sectorFileInfoGroup = sectorFileInfoGroup[nodeAddr];
        delete _sectorFileInfoGroup[groupId];
    }

    function deleteAllSectorFileInfoGroup(address nodeAddr, uint64 sectorId)
        public
    {
        uint64 groupNum = getSectorFileInfoGroupNum(nodeAddr, sectorId);
        for (uint64 i = 0; i < groupNum; i++) {
            deleteSectorFileInfoGroup(nodeAddr, sectorId, i);
        }
    }

    function deleteSectorInfo(address nodeAddr, uint64 sectorId) public {
        for (uint64 i = 0; i < sectorInfos[nodeAddr].length; i++) {
            if (sectorInfos[nodeAddr][i].SectorID == sectorId) {
                delete sectorInfos[nodeAddr][i];
                break;
            }
        }
    }

    function deleteSector(address nodeAddr, uint64 sectorId) public payable {
        deleteAllSectorFileInfoGroup(nodeAddr, sectorId);
        deleteSectorInfo(nodeAddr, sectorId);
    }

    function DeleteSecotr(SectorRef memory sectorRef) public {
        SectorInfo memory sectorInfo = GetSectorInfo(sectorRef);
        if (sectorInfo.FileNum > 0) {
            revert NotEmptySector(0, sectorInfo.FileNum);
        }
        deleteSector(sectorRef.NodeAddr, sectorRef.SectorId);
        emit DeleteSectorEvent(
            FsEvent.DELETE_SECTOR,
            block.number,
            msg.sender,
            sectorRef.SectorId
        );
    }
}
