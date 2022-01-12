import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { Config } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const tx = config.GetSetting();
    let res = await tx;
    assert(res.MinVolume.eq(1000000))
  });
});
