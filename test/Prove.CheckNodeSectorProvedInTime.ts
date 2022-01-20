import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space, prove } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res = prove.CheckNodeSectorProvedInTime({
      NodeAddr: addrs[24],
      SectorId: 1
    });
    expect(res).to.be.reverted;
  });
});
