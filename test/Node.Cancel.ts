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
      WalletAddr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
      NodeAddr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;

    let tx2 = node.NodeCancel("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266");
    expect(tx2).to.be.reverted;
  });
});
