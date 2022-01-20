import { assert, expect } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var get = path.basename(__filename);

describe(get, function () {

  it("get", async () => {
    const tx = fs.GetFileList(addrs[46]);
    let res = await tx;
    // console.log(tx)
    assert(res.length == 0)
  });

});
