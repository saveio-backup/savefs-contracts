import { expect, assert } from "chai";
import { addrs, node, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it('get 1', async () => {
    const tx = sector.GetSectorsForNode(addrs[13]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 0);
  });

  it('register node', async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[13],
      NodeAddr: addrs[13],
    },
      {
        value: 1000000
      }
    );
    await expect(tx).to.not.be.reverted;
  });

  it('create sector', async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[13],
      SectorID: 1,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 0,
      GroupNum: 1,
      IsPlots: false,
      FileList: []
    });
    await expect(tx).to.not.be.reverted;
  });

  it('get 2', async () => {
    const tx = sector.GetSectorsForNode(addrs[13]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 1);
  });

  it("delete sector", async () => {
    const tx = sector.DeleteSector({
      NodeAddr: addrs[13],
      SectorId: "1",
    });
    await expect(tx).to.not.be.reverted;
  });

  it('get 3', async () => {
    const tx = sector.GetSectorsForNode(addrs[13]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 0);
  });

});
