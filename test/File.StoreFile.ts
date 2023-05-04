import { assert, expect } from "chai";
import { addrs, file, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("space 1-1", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[43],
      Owner: addrs[43],
      Size: {
        Type: 1,
        Value: 1
      },
      BlockCount: {
        Type: 1,
        Value: (3600 * 24) / 5
      }
    });
    // await print(tx);
    await expect(tx).to.not.be.reverted;
    await expect(tx).to.emit(space, "SetUserSpaceEvent");
  });

  it("store file to space", async () => {
    const tx = file.StoreFile({
      FileHash: [1, 2, 3],
      FileOwner: addrs[43],
      FileDesc: [],
      Privilege: 1,
      FileBlockNum: 0,
      FileBlockSize: 0,
      ProveInterval: 1,
      ProveTimes: 1,
      ExpiredHeight: 100,
      CopyNum: 0,
      Deposit: 0,
      FileProveParam_: {
        RootHash: [1],
        FileID: [1],
      },
      ProveBlockNum: 1,
      BlockHeight: 1,
      ValidFlag: false,
      BlocksRoot: [],
      StorageType_: 0,
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

  it("store file to sector", async () => {
    const tx = file.StoreFile({
      FileHash: [1, 2, 3, 4, 5],
      FileOwner: addrs[43],
      FileDesc: [],
      Privilege: 1,
      FileBlockNum: 0,
      FileBlockSize: 0,
      ProveInterval: 1,
      ProveTimes: 1,
      ExpiredHeight: 100,
      CopyNum: 0,
      Deposit: 0,
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

  it("get file list", async () => {
    const tx = file.GetFileList(addrs[43]);
    let res = await tx;
    // console.log(tx)
    assert(res.length == 2)
  });

});
