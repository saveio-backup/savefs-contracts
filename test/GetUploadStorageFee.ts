import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { FileSystem } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  let fs: FileSystem;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = await fs.deployed();
    assert(res != undefined)
  });

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
      StorageType: 0,
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
      StorageType: 0,
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
      StorageType: 0,
    });
    expect(res2).to.not.be.reverted;
  });
});
