import { assert, expect } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res = fs.DeleteUnsettledFiles(addrs[10])
    expect(res).to.not.be.reverted;
    // console.log(await (await res).wait())
  });
});
