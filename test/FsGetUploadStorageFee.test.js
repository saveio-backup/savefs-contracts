const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FileSystem", () => {
  let fs;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    await fs.deployed();
  });

  it("FsGetUploadStorageFee", async () => {
    const res = fs.FsGetUploadStorageFee({
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
      WhiteList: {
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

    const res2 = fs.FsGetUploadStorageFee({
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
      WhiteList: {
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
});
