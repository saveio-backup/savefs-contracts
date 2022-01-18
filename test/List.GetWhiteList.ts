import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem } from "../typechain";
import { addrs, config, fs, node, space, list } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(`${scriptName} 1`, async () => {
    const res = list.GetWhiteList([]);
    expect(res).to.be.reverted;  // because list is empty
  });

  it(`${scriptName} 2`, async () => {
    const res = list.GetWhiteList([1, 2, 3]);
    expect(res).to.be.reverted;  // because list is empty
  })
});
