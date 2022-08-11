import { expect } from "chai";
import { assert } from "console";
import { addrs, node, sector, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);


// addrs[7] = "0xF168345D34E76118b2280dBcF905DE98e2905e61"
// console.log(addrs[7])

describe(name, () => {

  it("register", async () => {
    let tx = node.Register({
      Pledge: 10000,
      Profit: 10000,
      Volume: 1000 * 1000,
      RestVol: 10000,
      ServiceTime: 10000,
      WalletAddr: addrs[7],
      NodeAddr: addrs[7],
    }, {
      value: 1000000
    });
    // print(tx);
    await expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[7],
      SectorID: 1,
      Size: 10000,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 0,
      NextProveHeight: 10000,
      TotalBlockNum: 10000,
      FileNum: 0,
      GroupNum: 1,
      IsPlots: true,
      FileList: []
    });
    await expect(tx).to.not.be.reverted;
  })

  it("query 1", async () => {
    const res = await node.GetNodeInfoByWalletAddr(addrs[7]);
    // console.log(res)
    assert(!res.Volume.eq(0))
  });

  it("cancel", async () => {
    let tx = node.Cancel(addrs[7]);
    // print(tx)
    await expect(tx).to.not.be.reverted;

    let res = await (await tx).wait()
    // console.log(res)
  });

  it("query 2", async () => {
    const res = await node.GetNodeInfoByWalletAddr(addrs[7]);
    // console.log(res)
    assert(res.Volume.eq(0))
  });

});
