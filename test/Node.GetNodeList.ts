import { expect, assert } from "chai";
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
      WalletAddr: addrs[4],
      NodeAddr: addrs[4],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it("get list", async () => {
    const res = await node.GetNodeList();
    assert(res.length > 0);
  });

});
