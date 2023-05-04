import { assert, expect } from "chai";
import { addrs, file } from "./initialize";
import { randomBytes } from "crypto";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  let fileHash1 = randomBytes(32)
  let fileHash2 = randomBytes(32)

  it("store file 1", async () => {
    const tx = file.StoreFile({
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
      FileProveParam_: {
        RootHash: [1],
        FileID: [1],
      },
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
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("store file 2", async () => {
    const tx = file.StoreFile({
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
      FileProveParam_: {
        RootHash: [1],
        FileID: [1],
      },
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
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("get file infos", async () => {
    const tx = file.GetFileInfos(
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
