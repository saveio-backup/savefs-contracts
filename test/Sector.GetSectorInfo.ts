import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Sector } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  let node: Node;
  let config: Config;
  let sector: Sector;

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

  it("Deploy Sector", async () => {
    const S = await ethers.getContractFactory("Sector");
    sector = await S.deploy();
    let res = node.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Sector initialize", async () => {
    let tx = node.initialize(node.address);
    expect(tx).to.not.be.reverted;
  });
  // -------------------------------------------------------

  it(scriptName, async () => {
    const tx = sector.GetSectorInfo({
      SectorId: "1",
      NodeAddr: "0x0000000000000000000000000000000000000000",
    });
    let res = await tx;
    assert(res.Size.eq(0));
    // console.log(res);
  });
});
