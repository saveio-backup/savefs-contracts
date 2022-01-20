import { assert, expect } from "chai";
import { ContractTransaction } from "ethers";
import { ethers, network } from "hardhat";
import {
  FileSystem, Node, Config, Sector, Space, List, Prove, PDP
} from "../typechain";

var path = require('path');
var name = path.basename(__filename);

let addrs: Array<string> = [];
let config: Config;
let node: Node;
let fs: FileSystem;
let sector: Sector;
let space: Space;
let list: List;
let prove: Prove;
let pdp: PDP;

describe(name, function () {
  it("Get addrs", async () => {
    const accounts = await ethers.getSigners();
    for (const account of accounts) {
      addrs.push(account.address)
    }
  })

  it("Deploy Config", async () => {
    const Config = await ethers.getContractFactory("Config");
    config = await Config.deploy();
    let res = config.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Node", async () => {
    const Node = await ethers.getContractFactory("Node");
    node = await Node.deploy();
    let res = node.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Sector", async () => {
    const Sector = await ethers.getContractFactory("Sector");
    sector = await Sector.deploy();
    let res = sector.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Space", async () => {
    const Space = await ethers.getContractFactory("Space");
    space = await Space.deploy();
    let res = space.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy FileSystem", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = fs.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy List", async () => {
    const List = await ethers.getContractFactory("List");
    list = await List.deploy();
    let res = list.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Prove", async () => {
    const Prove = await ethers.getContractFactory("Prove");
    prove = await Prove.deploy();
    let res = prove.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Prove", async () => {
    const PDP = await ethers.getContractFactory("PDP");
    pdp = await PDP.deploy();
    let res = pdp.deployed();
    await expect(res).not.to.be.reverted;
  });

  // ----------

  it("initialize Node", async () => {
    let tx = node.initialize(config.address, sector.address);
    await expect(tx).not.to.be.reverted;
  });

  it("initialize Sector", async () => {
    let tx = sector.initialize(node.address);
    await expect(tx).not.to.be.reverted;
  });

  it("initialize FileSystem", async () => {
    let tx = fs.initialize(
      config.address, node.address, space.address, sector.address, prove.address
    );
    await expect(tx).not.to.be.reverted;
  });

  it("initialize Space", async () => {
    let tx = space.initialize(config.address, fs.address);
    await expect(tx).not.to.be.reverted;
  });

  it("initialize Prove", async () => {
    let tx = prove.initialize(
      config.address,
      fs.address,
      node.address,
      pdp.address,
      sector.address
    );
    await expect(tx).not.to.be.reverted;
  });

});

var print = async (tx: Promise<ContractTransaction>) => {
  let res = await (await tx).wait();
  console.log(res)
}

export {
  addrs,
  config,
  node,
  fs,
  sector,
  space,
  list,
  prove,
  pdp,
  print
};
