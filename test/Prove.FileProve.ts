import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space, prove, print } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const tx = prove.FileProve({
      FileHash: [1, 2, 3],
      ProveData: [],
      BlockHeight: 123,
      NodeWallet: addrs[21],
      Profit: 1,
      SectorID: 1
    });
    // await print(tx)
    await expect(tx).to.be.reverted;
  });
});
