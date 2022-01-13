import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    let tx = node.NodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[6],
      NodeAddr: addrs[6],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;

    const res = await node.GetNodeInfoByNodeAddr(addrs[6]);
    assert(res.NodeAddr == addrs[6]);
  });
});
