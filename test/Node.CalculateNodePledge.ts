import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get pledge", async () => {
    const tx = node.CalculateNodePledge({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[8],
      NodeAddr: addrs[8],
    });
    let res = await tx;
    assert(res.eq(1000000));
  });

});
