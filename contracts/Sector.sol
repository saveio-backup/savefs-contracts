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

    function initialize(Node _node) public initializer {
        node = _node;
    }

    function getSectorTotalSizeForNode(address nodeAddr)
        public
        pure
        returns (uint64)
    {
        // uint64 totalSize = 0;
        // mapping(uint64 => SectorInfo) memory nodeSectorInfos = sectorInfos[nodeAddr];
    }

    function CreateSector(SectorInfo memory sectorInfo) public view {
        require(sectorInfo.SectorID > 0, "sectorId is wrong");
        require(sectorInfo.Size > 0, "sector size is wrong");
        require(
            node.IsNodeRegisted(sectorInfo.NodeAddr),
            "node not registered"
        );
        // TODO
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
