import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space, prove, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("store file to sector", async () => {
    const tx = fs.StoreFile({
      FileHash: [3, 2, 1],
      FileOwner: addrs[49],
      FileDesc: [],
      Privilege: 1,
      FileBlockNum: 0,
      FileBlockSize: 0,
      ProveInterval: 1,
      ProveTimes: 1,
      ExpiredHeight: 100,
      CopyNum: 0,
      Deposit: 0,
      FileProveParam: [],
      ProveBlockNum: 1,
      BlockHeight: 1,
      ValidFlag: false,
      StorageType_: 1,
      RealFileSize: 1,
      PrimaryNodes: [
        addrs[49]
      ],
      CandidateNodes: [],
      ProveLevel_: 1,
      IsPlotFile: true,
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

  it("prove", async () => {
    const tx = prove.FileProve({
      FileHash: [3, 2, 1],
      ProveData: [],
      BlockHeight: 123,
      NodeWallet: addrs[49],
      Profit: 1,
      SectorID: 1
    });
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

});
