//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./Config.sol";
import "./Node.sol";

contract Sector is Initializable {
    Node node;

    mapping(address => SectorInfo[]) sectorInfos; // nodeAddr => SectorInfo[]

    error NotEnoughVolume(uint64 got, uint64 want);

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
        console.log(2);
        NodeInfo memory nodeInfo = node.GetNodeInfoByNodeAddr(
            sectorInfo.NodeAddr
        );
        console.log(3);
        uint64 totalSize = getSectorTotalSizeForNode(sectorInfo.NodeAddr);
        console.log(4);
        console.log("totalSize", totalSize,sectorInfo.Size + totalSize > nodeInfo.Volume);
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
}
