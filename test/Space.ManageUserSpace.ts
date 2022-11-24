import { assert, expect } from "chai";
import {
  addrs, space, print
} from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("param error", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[32],
      Owner: addrs[32],
      Size: {
        Type: 0,
        Value: 0
      },
      BlockCount: {
        Type: 0,
        Value: 0
      }
    });
    // await print(tx);
    let res = await (await tx).wait();
    if (res.events?.length == 1) {
      assert(res.events[0].event == "FsError");
    }
  });

  it("space 1-1", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[70],
      Owner: addrs[70],
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

  it("space 1-2", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[71],
      Owner: addrs[71],
      Size: {
        Type: 1,
        Value: 1
      },
      BlockCount: {
        Type: 2,
        Value: (3600 * 24) / 5
      }
    });
    // await print(tx);
    let res = await (await tx).wait();
    if (res.events?.length == 1) {
      assert(res.events[0].event == "FsError");
    }
  });

  let value = 0;
  it("get update cost", async () => {
    const tx = space.GetUpdateCost({
      WalletAddr: addrs[72],
      Owner: addrs[72],
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
    value = res.Value.toNumber();
  });

  it("space 1-1", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[73],
      Owner: addrs[73],
      Size: {
        Type: 1,
        Value: 1024000
      },
      BlockCount: {
        Type: 1,
        Value: 172800
      }
    }, {
      value: value
    });
    // await print(tx);
    await expect(tx).to.not.be.reverted;
    await expect(tx).to.emit(space, "SetUserSpaceEvent");
  });

});
