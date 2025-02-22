import { expect, assert } from "chai";
import { addrs, space, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get empty", async () => {
    const tx = space.GetUserSpace(addrs[37]);
    let res = await tx;
    // console.log(tx)
    assert(res.Remain.eq(0));
  });

  it("space 1-1", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[37],
      Owner: addrs[37],
      Size: {
        Type: 1,
        Value: 1
      },
      BlockCount: {
        Type: 1,
        Value: (3600 * 24) / 5
      }
    });
    // await print(tx);
    await expect(tx).to.not.be.reverted;
    await expect(tx).to.emit(space, "SetUserSpaceEvent");
  });

  it("get not empty", async () => {
    const tx = space.GetUserSpace(addrs[37]);
    let res = await tx;
    // console.log(tx)
    assert(res.Remain.eq(1));
  });

});
