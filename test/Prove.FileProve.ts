import assert from "assert";
import { expect } from "chai";
import { addrs, file, node, prove, sector } from "./initialize";

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
      WalletAddr: addrs[49],
      NodeAddr: addrs[49],
    },
      {
        value: 1000000
      }
    );
    expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[49],
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
      BlocksRoot: [],
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

  it("file prove", async () => {
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

  it('get sector info', async () => {
    const tx = sector.GetSectorsForNode(addrs[49]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 1);
  });

});
