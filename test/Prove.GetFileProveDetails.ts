import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space, prove } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("get", async () => {
    const tx = prove.GetFileProveDetails(addrs[17]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 0)
  });

});
