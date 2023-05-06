//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract Sector is Initializable, ISector, IFsEvent {
    INode node;
    uint64 SECTOR_FILE_INFO_GROUP_MAX_LEN;

    mapping(address => SectorInfo[]) sectorInfos; // nodeAddr => SectorInfo[]
    mapping(bytes => SectorFileInfoGroup) sectorFileInfoGroup; // nodeAddr + sectorId + groupId => groupKey
    mapping(bytes => SectorFileInfo[]) sectorFileInfoFileList; // groupKey => SectorFileInfo[]

    function initialize(
        INode _node,
        SectorConfig memory sectorConfig
    ) public initializer {
        node = _node;
        SECTOR_FILE_INFO_GROUP_MAX_LEN = sectorConfig
        .SECTOR_FILE_INFO_GROUP_MAX_LEN;
    }

    function CreateSector(
        SectorInfo memory sectorInfo
    ) public virtual override {
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

    function GetSectorInfo(
        SectorRef memory sectorRef
    ) public view virtual override returns (SectorInfo memory) {
        SectorInfo memory s;
        SectorInfo[] memory sectorInfo = sectorInfos[sectorRef.NodeAddr];
        for (uint64 i = 0; i < sectorInfo.length; i++) {
            if (sectorInfo[i].SectorID == sectorRef.SectorId) {
                s = sectorInfo[i];
                SectorFileInfoGroup memory g = getSectorFileInfoGroup(
                    sectorRef.NodeAddr,
                    s.SectorID,
                    s.GroupNum
                );
                bytes memory groupKey = GetGroupKey(
                    sectorRef.NodeAddr,
                    s.SectorID,
                    s.GroupNum
                );
                SectorFileInfo[] memory l = sectorFileInfoFileList[groupKey];
                bytes[] memory hashes = new bytes[](l.length);
                for (uint64 k = 0; k < l.length; k++) {
                    hashes[k] = l[k].FileHash;
                }
                s.FileList = hashes;
                break;
            }
        }
        return s;
    }

    function GetSectorsForNode(
        address nodeAddr
    ) public view virtual override returns (SectorInfo[] memory) {
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
                SectorInfo memory s = infos[i];
                SectorFileInfoGroup memory g = getSectorFileInfoGroup(
                    nodeAddr,
                    s.SectorID,
                    s.GroupNum
                );
                if (g.Null) {
                    continue;
                }
                bytes memory groupKey = GetGroupKey(
                    nodeAddr,
                    s.SectorID,
                    s.GroupNum
                );
                SectorFileInfo[] memory l = sectorFileInfoFileList[groupKey];
                bytes[] memory hashes = new bytes[](l.length);
                for (uint64 k = 0; k < l.length; k++) {
                    hashes[k] = l[k].FileHash;
                }
                s.FileList = hashes;
                res[j] = s;
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

    function UpdateSectorInfo(
        SectorInfo memory sector
    ) public payable virtual override {
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

        if (sectorInfo.FileList.length > 0) {
            bytes[] memory fl = new bytes[](sectorInfo.FileList.length - 1);
            uint64 j = 0;
            for (uint64 i = 0; i < sectorInfo.FileList.length; i++) {
                if (
                    keccak256(sectorInfo.FileList[i]) !=
                    keccak256(fileInfo.FileHash)
                ) {
                    fl[j] = sectorInfo.FileList[i];
                    j++;
                }
            }
            sectorInfo.FileList = fl;
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

    function GetSectorTotalSizeForNode(
        address nodeAddr
    ) private view returns (uint64) {
        uint64 totalSize = 0;
        SectorInfo[] memory nodeSectorInfos = sectorInfos[nodeAddr];
        for (uint256 i = 0; i < nodeSectorInfos.length; i++) {
            totalSize += nodeSectorInfos[i].Size;
        }
        return totalSize;
    }

    function getSectorFileInfoGroupNum(
        address nodeAddr,
        uint64 sectorId
    ) private view returns (uint64) {
        SectorInfo memory sectorInfo = GetSectorInfo(
            SectorRef({SectorId: sectorId, NodeAddr: nodeAddr})
        );
        return sectorInfo.GroupNum;
    }

    function deleteAllSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorId
    ) private {
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
        SectorFileInfo[] memory list,
        bytes memory fileHash
    ) private view returns (uint64, bool) {
        if (compareBytes(group.MinFileHash, fileHash) > 0 ||
            compareBytes(group.MaxFileHash, fileHash) < 0) {
            return (0, false);
        }
        uint64 start = 0;
        uint64 end = group.FileNum - 1;
        uint64 index = 0;
        while (start <= end) {
            index = (start + end) / 2;
            int result = compareBytes(list[index].FileHash, fileHash);
            if (result == 0) {
                return (index, true);
            } else if (result > 0) {
                end = index - 1;
            } else {
                start = index + 1;
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
        for (uint64 i = 1; i <= groupNum; i++) {
            SectorFileInfoGroup memory group = getSectorFileInfoGroup(
                nodeAddr,
                sectorId,
                i
            );
            bytes memory groupKey = GetGroupKey(
                nodeAddr,
                sectorId,
                i
            );
            SectorFileInfo[] memory list = sectorFileInfoFileList[groupKey];
            (uint64 index, bool found) = findFileInGroup(group, list, fileHash);
            if (found) {
                uint256 fileNum = sectorFileInfoFileList[groupKey].length;

                // delete element from sectorFileInfoFileList
                uint64 j = 0;
                for (uint64 k = 0; k < fileNum; k++) {
                    if (k != index) {
                        sectorFileInfoFileList[groupKey][
                        j
                        ] = sectorFileInfoFileList[groupKey][k];
                        j++;
                    }
                }
                sectorFileInfoFileList[groupKey].pop();

                if (index == 0) {
                    group.MinFileHash = fileHash;
                }
                if (index == fileNum - 1) {
                    group.MaxFileHash = fileHash;
                }
                group.FileNum--;
                setSectorFileInfoGroup(nodeAddr, sectorId, group);
                return false;
            }
        }
        return false;
    }

    function AddSectorInfo(
        address nodeAddr,
        SectorInfo memory sectorInfo
    ) private {
        sectorInfos[nodeAddr].push(sectorInfo);
    }

    function getSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorID,
        uint64 groupID
    ) private view returns (SectorFileInfoGroup memory) {
        bytes memory groupKey = abi.encodePacked(nodeAddr, sectorID, groupID);
        return sectorFileInfoGroup[groupKey];
    }

    function addSectorFileInfoToGroup(
        address nodeAddr,
        uint64 sectorID,
        SectorFileInfoGroup memory _sectorFileInfoGroup,
        SectorFileInfo memory sectorFileInfo
    ) private {
        bytes memory groupKey = GetGroupKey(nodeAddr, sectorID, _sectorFileInfoGroup.GroupId);
        uint64 index = _sectorFileInfoGroup.FileNum;
        for (uint64 i = 0; i < _sectorFileInfoGroup.FileNum; i++) {
            if (
                compareBytes(
                    sectorFileInfo.FileHash,
                    sectorFileInfoFileList[groupKey][i]
                    .FileHash
                ) < 0
            ) {
                index = i;
                break;
            }
        }
        sectorFileInfoFileList[groupKey].push(
            sectorFileInfo
        );
        // Add a new element at the end
        for (uint64 i = _sectorFileInfoGroup.FileNum; i > index; i--) {
            sectorFileInfoFileList[groupKey][
            i
            ] = sectorFileInfoFileList[groupKey][i - 1];
        }

        sectorFileInfoFileList[groupKey][index] = sectorFileInfo;

        if (_sectorFileInfoGroup.FileNum == 0) {
            _sectorFileInfoGroup.MinFileHash = sectorFileInfo.FileHash;
            _sectorFileInfoGroup.MaxFileHash = sectorFileInfo.FileHash;
        } else {
            if (
                compareBytes(
                    sectorFileInfo.FileHash,
                    _sectorFileInfoGroup.MinFileHash
                ) < 0
            ) {
                _sectorFileInfoGroup.MinFileHash = sectorFileInfo.FileHash;
            }

            if (
                compareBytes(
                    sectorFileInfo.FileHash,
                    _sectorFileInfoGroup.MaxFileHash
                ) > 0
            ) {
                _sectorFileInfoGroup.MaxFileHash = sectorFileInfo.FileHash;
            }
        }
        _sectorFileInfoGroup.FileNum++;
        sectorFileInfoGroup[groupKey] = _sectorFileInfoGroup;
    }

    function GetGroupKey(
        address nodeAddr,
        uint64 sectorID,
        uint64 groupID
    ) private pure returns (bytes memory) {
        return abi.encodePacked(nodeAddr, sectorID, groupID);
    }

    function compareBytes(
        bytes memory a,
        bytes memory b
    ) private pure returns (int) {
        for (uint i = 0; i < a.length && i < b.length; i++) {
            if (a[i] < b[i]) {
                return - 1;
            } else if (a[i] > b[i]) {
                return 1;
            }
        }
        if (a.length < b.length) {
            return - 1;
        } else if (a.length > b.length) {
            return 1;
        } else {
            return 0;
        }
    }

    function setSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorId,
        SectorFileInfoGroup memory group
    ) private {
        bytes memory groupKey = abi.encodePacked(
            nodeAddr,
            sectorId,
            group.GroupId
        );
        group.Null = false;
        sectorFileInfoGroup[groupKey] = group;
    }

    function deleteSectorFileInfoGroup(
        address nodeAddr,
        uint64 sectorId,
        uint64 groupId
    ) private {
        bytes memory groupKey = abi.encodePacked(nodeAddr, sectorId, groupId);
        delete sectorFileInfoGroup[groupKey];
        delete sectorFileInfoFileList[groupKey];
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
        bytes memory groupKey = GetGroupKey(nodeAddr, sectorID, groupNum);
        if (groupNum == 0) {
            groupNum = 1;
            groupCreated = true;
            groupInfo = SectorFileInfoGroup({
                FileNum: 0,
                GroupId: groupNum,
                MinFileHash: minFileHash,
                MaxFileHash: minFileHash,
                Null: false
            });
        } else {
            groupInfo = getSectorFileInfoGroup(nodeAddr, sectorID, groupNum);
            SectorFileInfo[] memory fileList = sectorFileInfoFileList[
            groupKey
            ];
            if (fileList.length == SECTOR_FILE_INFO_GROUP_MAX_LEN) {
                groupNum++;
                groupCreated = true;
                groupInfo = SectorFileInfoGroup({
                    FileNum: 0,
                    GroupId: groupNum,
                    MinFileHash: minFileHash,
                    MaxFileHash: minFileHash,
                    Null: false
                });
            }
        }
        addSectorFileInfoToGroup(nodeAddr, sectorID, groupInfo, sectorFileInfo);
        return groupCreated;
    }

    function IsSectorRefByFileInfo(
        address nodeAddr,
        uint64 sectorID
    ) public view virtual override returns (bool) {
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
