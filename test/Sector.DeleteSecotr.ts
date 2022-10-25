import { expect,assert } from "chai";
import { addrs, sector, node } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it(`get empty 1`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[81],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.SectorID.eq(0));
  })

  it(`register node`, async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[81],
      NodeAddr: addrs[81],
    },
      {
        value: 1000000
      }
    );
    expect(tx).to.not.be.reverted;
  })

  it(`create sector`, async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[81],
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
    expect(tx).to.not.be.reverted;
  })

  it(`get not empty`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[81],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.SectorID.eq(1024));
  });

  it("delete sector", async () => {
    const tx = sector.DeleteSector({
      NodeAddr: addrs[81],
      SectorId: "1024",
    });
    await expect(tx).to.not.be.reverted;
  });

  it(`get empty 2`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[81],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.SectorID.eq(0));
  })

});
