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
      Pledge: 2,
      Profit: 2,
      Volume: 1000 * 1000,
      RestVol: 3,
      ServiceTime: 4,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;

    const res = node.NodeWithDrawProfit("0x5FbDB2315678afecb367f032d93F642f64180aa3");
    // must be recerted because profit is 0
    expect(res).to.be.reverted;
  });
});
