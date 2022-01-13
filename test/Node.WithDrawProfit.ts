import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    let tx = node.NodeRegister({
      Pledge: 2,
      Profit: 2,
      Volume: 1000 * 1000,
      RestVol: 3,
      ServiceTime: 4,
      WalletAddr: addrs[3],
      NodeAddr: addrs[3],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;

    const res = node.NodeWithDrawProfit(addrs[3]);
    // must be recerted because profit is 0
    expect(res).to.be.reverted;
  });
});
