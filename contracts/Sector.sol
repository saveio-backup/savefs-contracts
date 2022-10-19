//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract Sector is Initializable, ISector, IFsEvent {
    struct SectorFileInfo {
        bytes FileHash;
        uint64 BlockCount;
    }

    struct SectorFileInfoGroup {
        uint64 GroupId;
        bytes MinFileHash;
        bytes MaxFileHash;
    }

    INode node;
    uint64 SECTOR_FILE_INFO_GROUP_MAX_LEN;

    mapping(address => SectorInfo[]) sectorInfos; // nodeAddr => SectorInfo[]
    mapping(string => SectorFileInfoGroup) sectorFileInfoGroup; // nodeAddr + sectorId + groupId => groupId
    mapping(uint64 => SectorFileInfo[]) sectorFileInfoFileList; // GroupId => SectorFileInfo[]

    function initialize(INode _node, SectorConfig memory sectorConfig)
        public
        initializer
    {
        node = _node;
        SECTOR_FILE_INFO_GROUP_MAX_LEN = sectorConfig
            .SECTOR_FILE_INFO_GROUP_MAX_LEN;
    }

    function CreateSector(SectorInfo memory sectorInfo)
        public
        virtual
        override
    {
        if (sectorInfo.SectorID == 0) {
            emit FsError("CreateSector", "SectorID is 0");
            return;
        }
        if (sectorInfo.Size == 0) {
            emit FsError("CreateSector", "Size is 0");
            return;
        }
        if (!node.IsNodeRegisted(sectorInfo.NodeAddr)) {
            emit FsError("CreateSector", "Node not registered");
            return;
        }
        if (
            GetSectorInfo(
                SectorRef({
                    SectorId: sectorInfo.SectorID,
                    NodeAddr: sectorInfo.NodeAddr
                })
            ).Size != 0
        ) {
            emit FsError("CreateSector", "Sector already exists");
            return;
        }
        NodeInfo memory nodeInfo = node.GetNodeInfoByWalletAddr(
            sectorInfo.NodeAddr
        );
        uint64 totalSize = GetSectorTotalSizeForNode(sectorInfo.NodeAddr);
        if (sectorInfo.Size + totalSize > nodeInfo.Volume) {
            emit FsError("CreateSector", "Insufficient volume");
            return;
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
        virtual
        override
        returns (SectorInfo memory)
    {
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
        virtual
        override
        returns (SectorInfo[] memory)
    {
        SectorInfo[] memory infos = sectorInfos[nodeAddr];
        uint64 num = 0;
        for (uint64 i = 0; i < infos.length; i++) {
            if (infos[i].Size > 0) {
                num++;
            }
        }
        SectorInfo[] memory res = new SectorInfo[](num);
        uint64 j = 0;
        for (uint64 i = 0; i < infos.length; i++) {
            if (infos[i].Size > 0) {
                res[j] = infos[i];
                j++;
            }
        }
        return res;
    }

    function DeleteSector(SectorRef memory sectorRef) public virtual override {
        SectorInfo memory sectorInfo = GetSectorInfo(sectorRef);
        if (sectorInfo.FileNum > 0) {
            emit FsError("DeleteSector", "NotEmptySector");
            return;
        }
        deleteSector(sectorRef.NodeAddr, sectorRef.SectorId);
        emit DeleteSectorEvent(
            FsEvent.DELETE_SECTOR,
            block.number,
            msg.sender,
            sectorRef.SectorId
        );
    }

    function UpdateSectorInfo(SectorInfo memory sector)
        public
        payable
        virtual
        override
    {
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
    ) public payable virtual override {
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

    function AddFileToSector(
        SectorInfo memory sectorInfo,
        FileInfo memory fileInfo
    ) public payable virtual override {
        if (
            sectorInfo.Used + fileInfo.FileBlockNum * fileInfo.FileBlockSize >
            sectorInfo.Size
        ) {
            emit FsError("AddFileToSector", "NotEnoughSpace");
            return;
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

    function GetSectorTotalSizeForNode(address nodeAddr)
        private
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

    function getSectorFileInfoGroupNum(address nodeAddr, uint64 sectorId)
        private
        view
        returns (uint64)
    {
        SectorInfo memory sectorInfo = GetSectorInfo(
            SectorRef({SectorId: sectorId, NodeAddr: nodeAddr})
        );
        return sectorInfo.GroupNum;
    }

    function deleteAllSectorFileInfoGroup(address nodeAddr, uint64 sectorId)
        private
    {
        uint64 groupNum = getSectorFileInfoGroupNum(nodeAddr, sectorId);
        for (uint64 i = 0; i < groupNum; i++) {
            deleteSectorFileInfoGroup(nodeAddr, sectorId, i);
        }
    }

    function deleteSectorInfo(address nodeAddr, uint64 sectorId) private {
        for (uint64 i = 0; i < sectorInfos[nodeAddr].length; i++) {
            if (sectorInfos[nodeAddr][i].SectorID == sectorId) {
                delete sectorInfos[nodeAddr][i];
                break;
            }
        }
    }

    function deleteSector(address nodeAddr, uint64 sectorId) private {
        deleteAllSectorFileInfoGroup(nodeAddr, sectorId);
        deleteSectorInfo(nodeAddr, sectorId);
    }

    function findFileInGroup(
        SectorFileInfoGroup memory group,
        bytes memory fileHash
    ) private view returns (uint64, bool) {
        SectorFileInfo[] memory fileList = sectorFileInfoFileList[
            group.GroupId
        ];
        if (fileList.length == 0) {
            return (0, false);
        }
        if (keccak256(group.MinFileHash) != keccak256(fileHash)) {
            return (0, false);
        }
        for (uint64 i = 0; i < fileList.length; i++) {
            if (keccak256(fileList[i].FileHash) == keccak256(fileHash)) {
                return (i, true);
            }
        }
        return (0, false);
    }

    function deleteSectorFileInfo(
        address nodeAddr,
        uint64 sectorId,
        bytes memory fileHash
    ) private returns (bool) {
        uint64 groupNum = getSectorFileInfoGroupNum(nodeAddr, sectorId);
        for (uint64 i = 0; i < groupNum; i++) {
            SectorFileInfoGroup memory group = getSectorFileInfoGroup(
                nodeAddr,
                sectorId,
                i
            );
            (uint64 index, bool found) = findFileInGroup(group, fileHash);
            if (found) {
                uint256 fileNum = sectorFileInfoFileList[group.GroupId].length;
                delete sectorFileInfoFileList[group.GroupId][index];
                if (index == 0) {
                    group.MinFileHash = fileHash;
                }
                if (index == fileNum - 1) {
                    group.MaxFileHash = fileHash;
                }
                setSectorFileInfoGroup(nodeAddr, sectorId, group);
                return true;
            }
        }
        return false;
    }

    function AddSectorInfo(address nodeAddr, SectorInfo memory sectorInfo)
        private
    {
        sectorInfos[nodeAddr].push(sectorInfo);
    }

    function getSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorID,
        uint64 groupID
    ) private view returns (SectorFileInfoGroup memory) {
        string memory groupKey = string(
            abi.encodePacked(nodeAddr, sectorID, groupID)
        );
        return sectorFileInfoGroup[groupKey];
    }

    function addSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorID,
        SectorFileInfoGroup memory _sectorFileInfoGroup,
        SectorFileInfo memory sectorFileInfo
    ) private {
        string memory groupKey = string(
            abi.encodePacked(nodeAddr, sectorID, _sectorFileInfoGroup.GroupId)
        );
        sectorFileInfoGroup[groupKey] = _sectorFileInfoGroup;
        sectorFileInfoFileList[_sectorFileInfoGroup.GroupId].push(
            sectorFileInfo
        );
    }

    function setSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorId,
        SectorFileInfoGroup memory group
    ) private {
        string memory groupKey = string(
            abi.encodePacked(nodeAddr, sectorId, group.GroupId)
        );
        sectorFileInfoGroup[groupKey] = group;
    }

    function deleteSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorId,
        uint64 groupId
    ) private {
        string memory groupKey = string(
            abi.encodePacked(nodeAddr, sectorId, groupId)
        );
        delete sectorFileInfoGroup[groupKey];
        delete sectorFileInfoFileList[groupId];
    }

    function addSectorFileInfo(
        address nodeAddr,
        uint64 sectorID,
        SectorFileInfo memory sectorFileInfo
    ) private returns (bool) {
        SectorFileInfoGroup memory groupInfo;
        uint64 groupNum;
        bool groupCreated;
        SectorInfo memory sectorInfo = GetSectorInfo(
            SectorRef({SectorId: sectorID, NodeAddr: nodeAddr})
        );
        bytes memory minFileHash;
        groupNum = sectorInfo.GroupNum;
        if (groupNum == 0) {
            groupNum = 1;
            groupCreated = true;
            groupInfo = SectorFileInfoGroup({
                GroupId: groupNum,
                MinFileHash: minFileHash,
                MaxFileHash: minFileHash
            });
        } else {
            groupInfo = getSectorFileInfoGroup(nodeAddr, sectorID, groupNum);
            SectorFileInfo[] memory fileList = sectorFileInfoFileList[
                groupInfo.GroupId
            ];
            if (fileList.length == SECTOR_FILE_INFO_GROUP_MAX_LEN) {
                groupNum++;
                groupCreated = true;
                groupInfo = SectorFileInfoGroup({
                    GroupId: groupNum,
                    MinFileHash: minFileHash,
                    MaxFileHash: minFileHash
                });
            }
        }
        addSectorFileInfoGroup(nodeAddr, sectorID, groupInfo, sectorFileInfo);
        return groupCreated;
    }

    function IsSectorRefByFileInfo(address nodeAddr, uint64 sectorID)
        public
        view
        virtual
        override
        returns (bool)
    {
        SectorInfo[] memory sectors = GetSectorsForNode(nodeAddr);
        for (uint256 i = 0; i < sectors.length; i++) {
            if (
                sectors[i].NodeAddr == nodeAddr &&
                sectors[i].SectorID == sectorID
            ) {
                return true;
            }
        }
        return false;
    }
}
