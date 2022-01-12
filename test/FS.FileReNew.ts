import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res = fs.StoreFile({
      FileHash: [1, 2, 3, 4, 5, 6, 7],
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

    const res2 = fs.GetFileInfo([1, 2, 3, 4, 5, 6, 7]);
    let tx2 = await res2;
    assert(tx2.ExpiredHeight.eq(10000000))
    // console.log(tx2.ExpiredHeight)

    const res3 = fs.FileReNew({
      FileHash: [1, 2, 3, 4, 5, 6, 7],
      FromAddr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
      ReNewTimes: 123,
    }, {
      value: 1000000
    })
    let tx3 = await res3;
    // console.log(tx3)

    const res4 = fs.GetFileInfo([1, 2, 3, 4, 5, 6, 7]);
    let tx4 = await res4;
    assert(tx4.ExpiredHeight.eq(12125440))
    // console.log(tx4.ExpiredHeight)
  });
});
