import { expect, assert } from "chai";
import { addrs, node } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("node register success", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[69],
      NodeAddr: addrs[69],
    }, {
      value: 1000000
    });
    await expect(tx).to.not.be.reverted;
  });

  it(`get 1`, async () => {
    let tx = node.GetNodeInfoByNodeAddr(addrs[69])
    let res = await tx;
    // console.log(res)
    assert(res.Volume.eq(1000 * 1000));
  });

  it(`update info`, async () => {
    let tx = node.UpdateNodeInfo({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1001,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[69],
      NodeAddr: addrs[69],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it(`get 2`, async () => {
    let tx = node.GetNodeInfoByNodeAddr(addrs[69])
    let res = await tx;
    // console.log(res)
    assert(res.Volume.eq(1000 * 1001));
  });

});
