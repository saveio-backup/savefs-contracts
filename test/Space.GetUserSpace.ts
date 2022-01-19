import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { addrs, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get", async () => {
    const res = space.GetUserSpace(addrs[37]);
    await expect(res).to.not.be.reverted;
  });
});
