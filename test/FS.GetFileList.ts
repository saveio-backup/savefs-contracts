import { assert, expect } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res = fs.StoreFile({
      FileHash: [1, 2, 3, 4],
      FileOwner: "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
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
    // console.log(await (await res).wait())

    const res2 = fs.GetFileList("0x70997970C51812dc3A010C7d01b50e0d17dc79C8");
    let tx = await res2;
    assert(tx.length == 1)
    // console.log(tx)
  });
});
