import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { print, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get update cost", async () => {
    const tx = space.GetUpdateCost();
    expect(tx).to.not.be.reverted;
  });
});
