import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";
import { randomBytes } from "crypto";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  let fileHash1 = randomBytes(32)
  let fileHash2 = randomBytes(32)

  it("store file 1", async () => {
    const tx = fs.StoreFile({
      FileHash: fileHash1,
      FileOwner: addrs[67],
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
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("store file 2", async () => {
    const tx = fs.StoreFile({
      FileHash: fileHash2,
      FileOwner: addrs[67],
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
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("get file infos", async () => {
    const tx = fs.GetFileInfos(
      [
        fileHash1,
        fileHash2,
      ]
    );
    let res = await tx
    // console.log(res)
    assert(res.length == 2)
  });

});
