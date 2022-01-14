import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { addrs, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const res = space.DeleteUserSpace(addrs[8]);
    expect(res).to.be.reverted; // userspace not exist
    // let tx = await (await res).wait();
    // console.log(tx)
  });
});
