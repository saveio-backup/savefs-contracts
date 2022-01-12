import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const res = space.GetUserSpace("0x0000000000000000000000000000000000000000");
    expect(res).to.be.reverted;
    // let tx = await res;
    // console.log(tx)
  });
});
