import { assert } from "chai";
import { addrs, node } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get pledge", async () => {
    const tx = node.CalculateNodePledge({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[0],
      NodeAddr: new Array(),
    });
    let res = await tx;
    assert(res.eq(1000000));
  });

});
