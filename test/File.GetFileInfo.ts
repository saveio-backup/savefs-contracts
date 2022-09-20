import { assert, expect } from "chai";
import { addrs, file, print } from "./initialize";
import { randomBytes } from "crypto";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  let fileHash = randomBytes(32)

  it("get file info fail", async () => {
    const tx = file.GetFileInfo(fileHash);
    let res = await tx;
    assert(res.ExpiredHeight.eq(0))
  });

  it("store file", async () => {
    const tx = file.StoreFile({
      FileHash: fileHash,
      FileOwner: addrs[66],
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
    // await print(tx);
    await expect(tx).not.to.be.reverted;
  });

  it("get file info success", async () => {
    const tx = file.GetFileInfo(fileHash);
    let res = await tx
    // console.log(res)
    assert(res.FileBlockNum.eq(1))
  });

});
