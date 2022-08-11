import { expect,assert } from "chai";
import { addrs, sector, node } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it(`get empty`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[70],
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
      WalletAddr: addrs[70],
      NodeAddr: addrs[70],
    },
      {
        value: 1000000
      }
    );
    await expect(tx1).to.not.be.reverted;
  })

  it(`create sector`, async () => {
    const tx2 = sector.CreateSector({
      NodeAddr: addrs[70],
      SectorID: 1024,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 0,
      GroupNum: 1014,
      IsPlots: false,
      FileList: []
    });
    expect(tx2).to.not.be.reverted;
  })

  it(`get not empty`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[70],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.Size.eq(1));
  });


  it("delete sector", async () => {
    const tx = sector.DeleteSector({
      NodeAddr: addrs[70],
      SectorId: "1024",
    });
    await expect(tx).to.not.be.reverted;
  });

  it(`get empty`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[70],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.Size.eq(0));
  })

});
