import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { addrs, print, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const tx = space.DeleteUserSpace(addrs[8]);
    await expect(tx).to.not.be.reverted; 
  });
});
