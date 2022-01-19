import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    let tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[7],
      NodeAddr: addrs[7],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;

    let tx2 = node.NodeCancel(addrs[7]);
    expect(tx2).to.be.reverted;
  });
});
