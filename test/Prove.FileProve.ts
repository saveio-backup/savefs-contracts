import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space, prove } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res = prove.FileProve({
      FileHash: [],
      ProveData: [],
      BlockHeight: 123,
      NodeWallet: addrs[21],
      Profit: 1,
      SectorID: 1
    });
    expect(res).to.be.reverted;
  });
});
