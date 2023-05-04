//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract PDPExtra {
    function checkSectorProveData(
        SectorInfo memory sectorInfo,
        SectorProveData memory proveData
    ) public pure returns (string memory) {
        if (proveData.ProveFileNum > getSectorFileNum(sectorInfo)) {
            return
                "[checkSectorProveData] proveFileNum larger than file num in sector";
        }
        if (proveData.ProveFileNum > proveData.BlockNum) {
            return
                "[checkSectorProveData] proveFileNum larger than challenged block num in sector";
        }
        if (proveData.BlockNum > sectorInfo.TotalBlockNum) {
            return
                "[checkSectorProveData] challenged block num larger than total block num in sector";
        }
        if (sectorInfo.IsPlots && proveData.PlotData.length == 0) {
            return "[checkSectorProveData] ";
        }
        return "";
    }

    function getSectorFileNum(
        SectorInfo memory sectorInfo
    ) internal pure returns (uint256) {
        return sectorInfo.FileNum;
    }

    function GetSectorFileInfosForSector(
        ISector sector,
        address nodeAddr,
        uint64 sectorId,
        uint64 fileNum
    ) public view returns (SectorFileInfo[] memory) {
        SectorFileInfo[] memory sectorFileInfos = new SectorFileInfo[](fileNum);
        // TODO merge group
        SectorInfo[] memory sectorInfos = sector.GetSectorsForNode(nodeAddr);
        SectorInfo memory sectorInfo;
        for (uint64 i = 0; i < sectorInfos.length; i++) {
            if (sectorInfos[i].SectorID == sectorId) {
                sectorInfo = sectorInfos[i];
                break;
            }
        }
        for (uint64 i = 0; i < fileNum; i++) {
            sectorFileInfos[i].FileHash = sectorInfo.FileList[i];
        }
        return sectorFileInfos;
    }

    function PrepareForPdpVerification1(
        IFile file,
        uint64 fileNum,
        SectorFileInfo[] memory sectorFileInfos,
        Challenge[] memory challenges,
        PdpVerificationReturns memory pReturns
    ) public view returns (PdpVerificationReturns memory) {
        uint64 offset = 0;
        uint64 curIndex = 0;

        for (uint64 i = 0; i < fileNum; i++) {
            SectorFileInfo memory sectorFileInfo = sectorFileInfos[i];
            bytes memory fileHash = sectorFileInfo.FileHash;
            uint64 blockCount = sectorFileInfo.BlockCount;

            uint64 start = offset;
            uint64 end = offset + blockCount - 1;

            for (uint64 j = curIndex; j < challenges.length; j++) {
                Challenge memory challenge = challenges[curIndex];
                if (challenge.Index >= start && challenge.Index <= end) {
                    if (curIndex == 0) {
                        pReturns.FileInfo_ = file.GetFileInfo(fileHash);
                    }
                    pReturns.FileIDs[i] = file
                        .GetFileInfo(fileHash)
                        .FileProveParam_
                        .FileID;
                    pReturns.RootHashes[i] = file
                        .GetFileInfo(fileHash)
                        .FileProveParam_
                        .RootHash;
                    curIndex++;
                    if (curIndex >= challenges.length) {
                        return pReturns;
                    }
                    continue;
                }
                break;
            }
            offset += blockCount;
        }
        return pReturns;
    }

    function PrepareForPdpVerification2(
        uint64 fileNum,
        SectorFileInfo[] memory sectorFileInfos,
        Challenge[] memory challenges,
        SectorProveData memory proveData,
        PdpVerificationReturns memory pReturns
    ) public pure returns (PdpVerificationReturns memory) {
        uint64 offset = 0;
        uint64 curIndex = 0;

        for (uint64 i = 0; i < fileNum; i++) {
            SectorFileInfo memory sectorFileInfo = sectorFileInfos[i];
            uint64 blockCount = sectorFileInfo.BlockCount;

            uint64 start = offset;
            uint64 end = offset + blockCount - 1;

            for (uint64 j = curIndex; j < challenges.length; j++) {
                Challenge memory challenge = challenges[curIndex];
                if (challenge.Index >= start && challenge.Index <= end) {
                    pReturns.Tags[i] = proveData.Tags[curIndex];
                    pReturns.Path[i] = proveData.MerklePath_[curIndex];
                    pReturns.UpdatedChal[i] = Challenge(
                        challenge.Index - uint32(start),
                        challenge.Rand
                    );
                    curIndex++;
                    if (curIndex >= challenges.length) {
                        return pReturns;
                    }
                    continue;
                }
                break;
            }
            offset += blockCount;
        }
        return pReturns;
    }
}
