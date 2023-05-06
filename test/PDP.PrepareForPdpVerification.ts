import assert from "assert";
import { expect } from "chai";
import { addrs, pdp, print, node, sector, file, prove } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("register node", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[9],
      NodeAddr: addrs[9],
    },
      {
        value: 1000000
      }
    );
    expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[9],
      SectorID: 1,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 0,
      GroupNum: 0,
      IsPlots: false,
      FileList: [[1]]
    });
    expect(tx).to.not.be.reverted;
  });

  it("store file to sector", async () => {
    const tx = file.StoreFile({
      FileHash: [1, 6, 9],
      FileOwner: addrs[9],
      FileDesc: [],
      Privilege: 1,
      FileBlockNum: 1,
      FileBlockSize: 1,
      ProveInterval: 1,
      ProveTimes: 1,
      ExpiredHeight: 100,
      CopyNum: 0,
      Deposit: 0,
      FileProveParam_: {
        RootHash: [2],
        FileID: [1],
      },
      ProveBlockNum: 1,
      BlockHeight: 1,
      ValidFlag: false,
      BlocksRoot: [],
      StorageType_: 1,
      RealFileSize: 1,
      PrimaryNodes: [
        addrs[9]
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
      FileHash: [1, 6, 9],
      ProveData_: {
        Proofs: [1],
        BlockNum: 1,
        Tags: [[1]],
        MerklePath_: [{
          PathLen: 1,
          Path: [{ Layer: 1, Index: 1, Hash: [1] }],
        }],
      },
      BlockHeight: 1,
      NodeWallet: addrs[9],
      Profit: 1,
      SectorID: 1
    });
    // let res = await (await tx).wait();
    // console.log(res.events)
    await expect(tx).to.not.be.emit(prove, "FsError");
  });


  it("submit", async () => {
    const tx = pdp.PrepareForPdpVerification(
      {
        SectorInfo_: {
          NodeAddr: addrs[9],
          SectorID: 1,
          Size: 1,
          Used: 0,
          ProveLevel_: 1,
          FirstProveHeight: 1,
          NextProveHeight: 1,
          TotalBlockNum: 1,
          FileNum: 0,
          GroupNum: 0,
          IsPlots: false,
          FileList: [[1]]
        },
        Challenges: [{ Index: 0, Rand: 1 }],
        ProveData: {
          ProveFileNum: 1,
          Proofs: [1],
          BlockNum: 1,
          Tags: [[1]],
          MerklePath_: [{
            PathLen: 1,
            Path: [{ Layer: 1, Index: 1, Hash: [1] }],
          }],
          PlotData: [],
        }
      }
    )
    const res = await tx;
    // console.log(res.RootHashes)
  });

});
