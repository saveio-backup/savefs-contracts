import { expect, assert } from "chai";
import { addrs, node, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it(`get 1`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[12],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.Size.eq(0));
  })

  it(`register node`, async () => {
    const tx1 = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[12],
      NodeAddr: addrs[12],
    },
      {
        value: 1000000
      }
    );
    await expect(tx1).to.not.be.reverted;
  })

  it(`create sector`, async () => {
    const tx2 = sector.CreateSector({
      NodeAddr: addrs[12],
      SectorID: 1024,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 1,
      GroupNum: 1014,
      IsPlots: false,
      FileList: []
    });
    expect(tx2).to.not.be.reverted;
  })

  it(`get 2`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[12],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.Size.eq(1));
  });

});
