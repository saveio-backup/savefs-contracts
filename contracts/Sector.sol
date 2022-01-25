//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./Node.sol";

contract Sector is Initializable {
    uint64 SECTOR_FILE_INFO_GROUP_MAX_LEN = 5000;

    Node node;

    mapping(address => SectorInfo[]) sectorInfos; // nodeAddr => SectorInfo[]
    mapping(address => mapping(uint64 => mapping(uint64 => uint64))) sectorFileInfoGroup; // nodeAddr => sectorId => groupId

    event CreateSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId,
        ProveLevel provLevel,
        uint64 size,
        bool isPlots
    );

    event DeleteSectorEvent(
        FsEvent eventType,
        uint256 blockHeight,
        address walletAddr,
        uint64 sectorId
    );

    error NotEnoughVolume(uint64 got, uint64 want);
    error NotEmptySector(uint64 got, uint64 want);
    error NotEnoughSpace();

    function initialize(Node _node) public initializer {
        node = _node;
    }

    function GetSectorTotalSizeForNode(address nodeAddr)
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
        uint64 totalSize = GetSectorTotalSizeForNode(sectorInfo.NodeAddr);
        if (sectorInfo.Size + totalSize > nodeInfo.Volume) {
            revert NotEnoughVolume(
                nodeInfo.Volume,
                sectorInfo.Size + totalSize
            );
        }
        sectorInfos[sectorInfo.NodeAddr].push(sectorInfo);
        emit CreateSectorEvent(
            FsEvent.CREATE_SECTOR,
            block.number,
            sectorInfo.NodeAddr,
            sectorInfo.SectorID,
            sectorInfo.ProveLevel_,
            sectorInfo.Size,
            sectorInfo.IsPlots
        );
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

    function GetSectorsForNode(address nodeAddr)
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
        mapping(uint64 => mapping(uint64 => uint64))
            storage _sectorFileInfoGroup = sectorFileInfoGroup[nodeAddr];
        mapping(uint64 => uint64) storage groups = _sectorFileInfoGroup[
            sectorId
        ];
        delete groups[groupId];
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

    function DeleteSector(address nodeAddr, uint64 sectorId) public payable {
        deleteAllSectorFileInfoGroup(nodeAddr, sectorId);
        deleteSectorInfo(nodeAddr, sectorId);
    }

    function DeleteSecotr(SectorRef memory sectorRef) public {
        SectorInfo memory sectorInfo = GetSectorInfo(sectorRef);
        if (sectorInfo.FileNum > 0) {
            revert NotEmptySector(0, sectorInfo.FileNum);
        }
        DeleteSector(sectorRef.NodeAddr, sectorRef.SectorId);
        emit DeleteSectorEvent(
            FsEvent.DELETE_SECTOR,
            block.number,
            msg.sender,
            sectorRef.SectorId
        );
    }

    function deleteSectorFileInfo(
        address nodeAddr,
        uint64 sectorId,
        bytes memory fileHash
    ) public returns (bool) {
        uint64 groupNum = getSectorFileInfoGroupNum(nodeAddr, sectorId);
        for (uint64 i = 0; i < groupNum; i++) {}
        // TODO
    }

    function UpdateSectorInfo(SectorInfo memory sector) public payable {
        SectorInfo[] storage _sectorInfos = sectorInfos[sector.NodeAddr];
        for (uint64 i = 0; i < _sectorInfos.length; i++) {
            if (_sectorInfos[i].SectorID == sector.SectorID) {
                _sectorInfos[i] = sector;
                break;
            }
        }
        sectorInfos[sector.NodeAddr] = _sectorInfos;
    }

    function DeleteFileFromSector(
        SectorInfo memory sectorInfo,
        FileInfo memory fileInfo
    ) public payable {
        bool groupDeleted = deleteSectorFileInfo(
            sectorInfo.NodeAddr,
            sectorInfo.SectorID,
            fileInfo.FileHash
        );
        sectorInfo.FileNum--;
        sectorInfo.TotalBlockNum -= fileInfo.FileBlockNum;
        sectorInfo.Used -= fileInfo.FileBlockNum * fileInfo.FileBlockSize;
        if (groupDeleted) {
            sectorInfo.GroupNum--;
        }
        if (sectorInfo.FileNum == 0) {
            sectorInfo.NextProveHeight = 0;
        }
        UpdateSectorInfo(sectorInfo);
    }

    function getSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorID,
        uint64 groupID
    ) public returns (SectorFileInfoGroup memory) {
        // TODO
    }

    function addSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorID,
        SectorFileInfoGroup memory _sectorFileInfoGroup,
        SectorFileInfo memory sectorFileInfo
    ) public {
        // TODO
    }

    struct SectorFileInfo {
        bytes FileHash;
        uint64 BlockCount;
    }

    struct SectorFileInfoGroup {
        uint64 FileNum;
        uint64 GroupId;
        bytes MinFileHash;
        bytes MaxFileHash;
        SectorFileInfo[] FileList;
    }

    function addSectorFileInfo(
        address nodeAddr,
        uint64 sectorID,
        SectorFileInfo memory sectorFileInfo
    ) public payable returns (bool) {
        SectorFileInfoGroup memory groupInfo;
        uint64 groupNum;
        bool groupCreated;
        SectorInfo memory sectorInfo = GetSectorInfo(
            SectorRef({SectorId: sectorID, NodeAddr: nodeAddr})
        );
        bytes memory minFileHash;
        SectorFileInfo[] memory fileList;
        groupNum = sectorInfo.GroupNum;
        if (groupNum == 0) {
            groupNum = 1;
            groupCreated = true;
            groupInfo = SectorFileInfoGroup({
                FileNum: 0,
                GroupId: groupNum,
                MinFileHash: minFileHash,
                MaxFileHash: minFileHash,
                FileList: fileList
            });
        } else {
            groupInfo = getSectorFileInfoGroup(nodeAddr, sectorID, groupNum);
            if (groupInfo.FileNum == SECTOR_FILE_INFO_GROUP_MAX_LEN) {
                groupNum++;
                groupCreated = true;
                groupInfo = SectorFileInfoGroup({
                    FileNum: 0,
                    GroupId: groupNum,
                    MinFileHash: minFileHash,
                    MaxFileHash: minFileHash,
                    FileList: fileList
                });
            }
        }
        addSectorFileInfoGroup(nodeAddr, sectorID, groupInfo, sectorFileInfo);
        return groupCreated;
    }

    function AddFileToSector(
        SectorInfo memory sectorInfo,
        FileInfo memory fileInfo
    ) public payable {
        if (
            sectorInfo.Used + fileInfo.FileBlockNum * fileInfo.FileBlockSize >
            sectorInfo.Size
        ) {
            revert NotEnoughSpace();
        }
        bool groupCreated = addSectorFileInfo(
            sectorInfo.NodeAddr,
            sectorInfo.SectorID,
            SectorFileInfo({
                FileHash: fileInfo.FileHash,
                BlockCount: fileInfo.FileBlockNum
            })
        );
        sectorInfo.FileNum++;
        sectorInfo.Used += fileInfo.FileBlockNum * fileInfo.FileBlockSize;
        sectorInfo.TotalBlockNum += fileInfo.FileBlockNum;
        if (groupCreated) {
            sectorInfo.GroupNum++;
        }
        UpdateSectorInfo(sectorInfo);
    }
}
