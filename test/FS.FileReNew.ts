import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("store file", async () => {
    const tx = fs.StoreFile({
      FileHash: [1, 2, 3, 4, 5, 6, 7],
      FileOwner: addrs[65],
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
      BlocksRoot: [],
      StorageType_: 1,
      RealFileSize: 1,
      PrimaryNodes: [],
      CandidateNodes: [],
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
    await expect(tx).to.not.be.reverted;
  });

  it("get file info", async () => {
    const tx = fs.GetFileInfo([1, 2, 3, 4, 5, 6, 7]);
    let res = await tx;
    // console.log(res)
    assert(res.ExpiredHeight.eq(10000000))
  });

  it("renew file", async () => {
    const tx = fs.FileReNew({
      FileHash: [1, 2, 3, 4, 5, 6, 7],
      FromAddr: addrs[25],
      ReNewTimes: 123,
    }, {
      value: 1000000
    })
    let res = await tx;
    await expect(tx).to.not.be.reverted;
  });

  it("get file info again", async () => {
    const tx = fs.GetFileInfo([1, 2, 3, 4, 5, 6, 7]);
    let res = await tx;
    assert(res.ExpiredHeight.eq(12125440))
  });

});
