import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Config, Node } from "../typechain";
import { addrs, config, fs, node, space, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("register", async () => {
    let tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[2],
      NodeAddr: addrs[2],
    }, {
      value: 1000000
    });
    await expect(tx).to.not.be.reverted;
  });

  it("node update success", async () => {
    let tx = node.NodeUpdate({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[2],
      NodeAddr: addrs[2],
    });
    // print(tx)
    await expect(tx).to.not.be.reverted;
  });

});
