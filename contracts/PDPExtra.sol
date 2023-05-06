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

            uint64 start = offset;
            uint64 end = offset + sectorFileInfo.BlockCount - 1;

            for (uint64 j = curIndex; j < challenges.length; j++) {
                Challenge memory challenge = challenges[curIndex];
                if (challenge.Index >= start && challenge.Index <= end) {
                    if (curIndex == 0) {
                        pReturns.FileInfo_ = file.GetFileInfo(
                            sectorFileInfo.FileHash
                        );
                    }
                    pReturns.FileIDs[i] = file
                        .GetFileInfo(sectorFileInfo.FileHash)
                        .FileProveParam_
                        .FileID;
                    pReturns.RootHashes[i] = file
                        .GetFileInfo(sectorFileInfo.FileHash)
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
            offset += sectorFileInfo.BlockCount;
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
