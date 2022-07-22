import { expect } from "chai";
import { addrs, node } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("register", async () => {
    let tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[2],
      NodeAddr: addrs[2],
    }, {
      value: 1000000
    });
    await expect(tx).to.not.be.reverted;
  });

  it("node update with same pledge", async () => {
    let tx = node.NodeUpdate({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[2],
      NodeAddr: addrs[2],
    });
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("node update with different pledge", async () => {
    let nodeInfo = {
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000 * 100,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[2],
      NodeAddr: addrs[2],
    }

    let query = node.GetPledgeUpdate(nodeInfo)
    let res = await query
    // console.log(res)

    let tx = node.NodeUpdate(nodeInfo, {value: res});
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

});
