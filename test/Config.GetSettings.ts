import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { Config } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  let config: Config;

  it("Deploy", async () => {
    const C = await ethers.getContractFactory("Config");
    config = await C.deploy();
    let res = await config.deployed();
    assert(res != undefined)
  });

  it(scriptName, async () => {
    const tx = config.GetSetting();
    let res = await tx;
    assert(res.MinVolume.eq(1000000))
  });
});
