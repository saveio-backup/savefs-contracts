import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const res = fs.WhiteListOperate({
      FileHash: [1, 2, 3],
      Op: 0,
      List: []
    });
    expect(res).not.to.be.reverted;

  });
});
