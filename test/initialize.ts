import { assert, expect } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Sector, Space } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

let config: Config;
let node: Node;
let fs: FileSystem;
let sector: Sector;
let space: Space;

describe(scriptName, function () {

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

  it("Deploy Sector", async () => {
    const Sector = await ethers.getContractFactory("Sector");
    sector = await Sector.deploy();
    let res = sector.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Deploy Space", async () => {
    const Space = await ethers.getContractFactory("Space");
    space = await Space.deploy();
    let res = space.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Deploy FileSystem", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = fs.deployed();
    expect(res).to.not.be.reverted;
  });

  it("initialize Node", async () => {
    let tx = node.initialize(config.address);
    expect(tx).to.not.be.reverted;
  });

  it("initialize FileSystem", async () => {
    let tx = fs.initialize(config.address, node.address, space.address);
    expect(tx).to.not.be.reverted;
  });

  it("initialize Sector", async () => {
    let tx = sector.initialize(node.address);
    expect(tx).to.not.be.reverted;
  });

  it("initialize Space", async () => {
    let tx = space.initialize(config.address, fs.address);
    expect(tx).to.not.be.reverted;
  });

});

export {
  config,
  node,
  fs,
  sector,
  space
};
