import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { config, fs, node, space } from "./initialize";

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
      WalletAddr: "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
      NodeAddr: "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;

    const res = await node.GetNodeInfoByWalletAddr("0xbDA5747bFD65F08deb54cb465eB87D40e51B197E");
    assert(res.WalletAddr == "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E");
  });
});
