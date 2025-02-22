import assert from "assert";
import { expect } from "chai";
import { addrs, file, node, prove, sector, pdp } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("register node", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[99],
      NodeAddr: addrs[99],
    },
      {
        value: 1000000
      }
    );
    expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[99],
      SectorID: 1,
      Size: 100,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 200,
      TotalBlockNum: 0,
      FileNum: 0,
      GroupNum: 0,
      IsPlots: false,
      FileList: []
    });
    expect(tx).to.not.be.reverted;
  });

  it("store file to sector", async () => {
    const tx = file.StoreFile({
      FileHash: [3, 2, 1, 5, 3],
      FileOwner: addrs[99],
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
      PrimaryNodes: [
        addrs[99]
      ],
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

  it("file prove", async () => {
    const tx = prove.FileProve({
      FileHash: [3, 2, 1, 5, 3],
      ProveData_: {
        Proofs: [],
        BlockNum: 1,
        Tags: [[1]],
        MerklePath_: [{
          PathLen: 1,
          Path: [{ Layer: 1, Index: 1, Hash: [1] }],
        }],
      },
      BlockHeight: 123,
      NodeWallet: addrs[99],
      Profit: 1,
      SectorID: 1
    });
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it('get sector info', async () => {
    const tx = sector.GetSectorsForNode(addrs[99]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 1);
    assert(res[0].FileList.length == 1);
  });

  it("delete files", async () => {
    const tx = file.DeleteFiles([
      [3, 2, 1, 5, 3]
    ]);
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it('get sector info 2', async () => {
    const tx = sector.GetSectorsForNode(addrs[99]);
    let res = await tx;
    // console.log(res[0].FileList)
    assert(res.length == 1);
    assert(res[0].FileList.length == 0);
  });

});
