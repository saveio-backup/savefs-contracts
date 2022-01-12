import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { Space } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const res = space.GetUserSpace("0x0000000000000000000000000000000000000000");
    // because empty
    expect(res).to.be.reverted;
    // let tx = await res;
    // console.log(tx)
  });
});
