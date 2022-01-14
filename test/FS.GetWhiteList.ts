import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const res = fs.GetWhiteList([]);
    expect(res).to.be.reverted;

    const res2 = fs.GetWhiteList([1, 2, 3]);
    expect(res2).to.be.reverted;  // because list is empty
  });
});
