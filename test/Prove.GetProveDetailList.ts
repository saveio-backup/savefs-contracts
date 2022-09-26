import { assert, expect } from "chai";
import { addrs, file, node, prove, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("get 0", async () => {
    const tx = prove.GetProveDetailList(addrs[17]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 0)
  });

  it("register node", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[17],
      NodeAddr: addrs[17],
    },
      {
        value: 1000000
      }
    );
    expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[17],
      SectorID: 1,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 1,
      GroupNum: 1,
      IsPlots: true,
      FileList: []
    });
    expect(tx).to.not.be.reverted;
  });

  it("store file to sector", async () => {
    const tx = file.StoreFile({
      FileHash: [3, 2, 1, 4, 7, 8],
      FileOwner: addrs[17],
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
      BlocksRoot: [],
      StorageType_: 1,
      RealFileSize: 1,
      PrimaryNodes: [
        addrs[17]
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

  it("file prove", async () => {
    const tx = prove.FileProve({
      FileHash: [3, 2, 1, 4, 7, 8],
      ProveData: [],
      BlockHeight: 123,
      NodeWallet: addrs[17],
      Profit: 1,
      SectorID: 1
    });
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });
  
  it("get larger than 0", async () => {
    const tx = prove.GetProveDetailList([3, 2, 1, 4, 7, 8]);
    let res = await tx;
    // console.log(res)
    assert(res.length > 0)
  });

});
