import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space, prove } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("check time", async () => {
    const res = prove.CheckNodeSectorProvedInTime({
      NodeAddr: addrs[24],
      SectorId: 1
    });
    expect(res).to.be.reverted;
  });

});
