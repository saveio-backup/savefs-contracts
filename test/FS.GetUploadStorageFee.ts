import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { print, config, fs, node, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it(`fileSize<0`, async () => {
    const tx = fs.GetUploadStorageFee({
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
      WhiteList_: [
        {
          Addr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
          BaseHeight: 1,
          ExpireHeight: 100,
        },
      ],
      Share: true,
      StorageType_: 0,
    });
    expect(tx).to.be.reverted;
  });

  it(`fileSize>0 num==0`, async () => {
    const tx = fs.GetUploadStorageFee({
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
      WhiteList_: [
        {
          Addr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
          BaseHeight: 1,
          ExpireHeight: 100,
        },
      ],
      Share: true,
      StorageType_: 0,
    });
    expect(tx).to.not.be.reverted;
  });

  it(`fileSize>0 num>0`, async () => {
    const tx = fs.GetUploadStorageFee({
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
      WhiteList_: [
        {
          Addr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
          BaseHeight: 1,
          ExpireHeight: 100,
        },
      ],
      Share: true,
      StorageType_: 0,
    });
    expect(tx).to.not.be.reverted;
  });
  
});
