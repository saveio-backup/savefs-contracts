import { expect } from "chai";
import { assert } from "console";
import { addrs, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get update cost", async () => {
    const tx = space.GetUpdateCost({
      WalletAddr: addrs[50],
      Owner: addrs[50],
      Size: {
        Type: 1,
        Value: 1024000
      },
      BlockCount: {
        Type: 1,
        Value: 172800
      }
    });
    expect(tx).not.to.be.reverted;

    let res = await tx;
    // console.log(res)
    assert(res.Value.gt(0));
  });
  
});
