import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  let node: Node;
  let config: Config;

  it("Deploy Config", async () => {
    const Config = await ethers.getContractFactory("Config");
    config = await Config.deploy();
    let res = config.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Deploy Node", async () => {
    const Node = await ethers.getContractFactory("Node");
    node = await Node.deploy();
    let res = node.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Node initialize", async () => {
    let tx = node.initialize(config.address);
    expect(tx).to.not.be.reverted;
  });

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
