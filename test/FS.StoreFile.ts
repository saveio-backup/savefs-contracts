import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {
  let config: Config;
  let node: Node;
  let fs: FileSystem;
  let space: Space;

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

  it("initialize Node", async () => {
    let tx = node.initialize(config.address);
    expect(tx).to.not.be.reverted;
  });
  
  it("Deploy Space", async () => {
    const Space = await ethers.getContractFactory("Space");
    space = await Space.deploy();
    let res = space.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Space initialize", async () => {
    let tx = space.initialize();
    expect(tx).to.not.be.reverted;
  });

  it("Deploy FileSystem", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = fs.deployed();
    expect(res).to.not.be.reverted;
  });

  it("initialize FileSystem", async () => {
    let tx = fs.initialize(config.address, node.address, space.address);
    expect(tx).to.not.be.reverted;
  });
  // --------------------------------------------------

  it(scriptName, async () => {
    const res = fs.StoreFile({
      FileHash: [1, 2, 3, 4, 5, 6],
      FileOwner: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
      FileDesc: [],
      Privilege: 1,
      FileBlockNum: 1,
      FileBlockSize: 1,
      ProveInterval: 1,
      ProveTimes: 1,
      ExpiredHeight: 10000000,
      CopyNum: 1,
      Deposit: 1,
      FileProveParam: [],
      ProveBlockNum: 1,
      BlockHeight: 1,
      ValidFlag: false,
      StorageType_: 1,
      RealFileSize: 1,
      PrimaryNodes: { AddrNum: 1, AddrList: [] },
      CandidateNodes: { AddrNum: 1, AddrList: [] },
      ProveLevel_: 1,
      IsPlotFile: false,
      PlotInfo_: {
        NumberID: 1,
        StartNonce: 1,
        Nonces: 1,
      }
    }, {
      value: 1000000
    })
    expect(res).to.not.be.reverted;
  });
});
