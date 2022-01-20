import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { Config } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("high", async () => {
    const tx = config.GetSettingWithProveLevel(0);
    let res = await tx;
    // console.log(res)
    assert(res.DefaultProvePeriod.eq(17280))
  });

  it("medieum", async () => {
    const tx = config.GetSettingWithProveLevel(1);
    let res = await tx;
    // console.log(res)
    assert(res.DefaultProvePeriod.eq(34560))
  });

  it("low", async () => {
    const tx = config.GetSettingWithProveLevel(2);
    let res = await tx;
    // console.log(res)
    assert(res.DefaultProvePeriod.eq(138240))
  });

});
