import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  let config: Config;
  let node: Node;
  let fs: FileSystem;

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

  it("Deploy FileSystem", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = fs.deployed();
    expect(res).to.not.be.reverted;
  });

  it("initialize FileSystem", async () => {
    let tx = fs.initialize(config.address, node.address);
    expect(tx).to.not.be.reverted;
  });
  // --------------------------------------------------

  it(`${scriptName} fileSize<0`, async () => {
    const res = fs.GetUploadStorageFee({
      FileDesc: [],
      FileSize: 0,
      ProveInterval: 1,
      ProveLevel: 1,
      ExpiredHeight: 100,
      Privilege: 0,
      CopyNum: 0,
      Encrypt: false,
      EncryptPassword: [],
      RegisterDNS: true,
      BindDNS: true,
      DnsURL: [],
      WhiteList_: {
        Num: 0,
        List: [
          {
            Addr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
            BaseHeight: 1,
            ExpireHeight: 100,
          },
        ],
      },
      Share: true,
      StorageType_: 0,
    });
    expect(res).to.be.reverted;
  });

  it(`${scriptName} fileSize>0 num==0`, async () => {
    const res2 = fs.GetUploadStorageFee({
      FileDesc: [],
      FileSize: 1,
      ProveInterval: 1,
      ProveLevel: 1,
      ExpiredHeight: 100,
      Privilege: 0,
      CopyNum: 0,
      Encrypt: false,
      EncryptPassword: [],
      RegisterDNS: true,
      BindDNS: true,
      DnsURL: [],
      WhiteList_: {
        Num: 0,
        List: [
          {
            Addr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
            BaseHeight: 1,
            ExpireHeight: 100,
          },
        ],
      },
      Share: true,
      StorageType_: 0,
    });
    expect(res2).to.not.be.reverted;
  });

  it(`${scriptName} fileSize>0 num>0`, async () => {
    const res2 = fs.GetUploadStorageFee({
      FileDesc: [],
      FileSize: 1,
      ProveInterval: 1,
      ProveLevel: 1,
      ExpiredHeight: 100,
      Privilege: 0,
      CopyNum: 0,
      Encrypt: false,
      EncryptPassword: [],
      RegisterDNS: true,
      BindDNS: true,
      DnsURL: [],
      WhiteList_: {
        Num: 1,
        List: [
          {
            Addr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
            BaseHeight: 1,
            ExpireHeight: 100,
          },
        ],
      },
      Share: true,
      StorageType_: 0,
    });
    expect(res2).to.not.be.reverted;
  });
});
