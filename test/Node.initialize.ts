import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {
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

});
