import { assert, expect } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res2 = fs.GetUnSettledFileList(addrs[7]);
    let tx = await res2;
    assert(tx.length == 0)
    // console.log(tx)
  });
});
