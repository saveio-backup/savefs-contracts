import { expect, assert } from "chai";
import { addrs, node } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it("register", async () => {
    let tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[5],
      NodeAddr: addrs[5],
    }, {
      value: 1000000
    });
    await expect(tx).to.not.be.reverted;
  });

  it("query", async () => {
    const res = await node.GetNodeInfoByWalletAddr(addrs[5]);
    assert(res.WalletAddr == addrs[5]);
  });
});
