import { assert, expect } from "chai";
import { addrs, node, sector } from "./initialize";

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
      FileNum: 1,
      GroupNum: 1,
      IsPlots: false,
      FileList: []
    });
    expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[9],
      SectorID: 2,
      Size: 1000 * 1001,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 1,
      GroupNum: 1,
      IsPlots: false,
      FileList: []
    });
    // sector.Size > node.Volume
    let res = await (await tx).wait();
    if (res.events?.length == 1) {
      assert(res.events[0].event == "FsError");
    }
  });

});
