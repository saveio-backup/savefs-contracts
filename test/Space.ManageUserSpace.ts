import { expect } from "chai";
import {
  addrs, space
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
    await expect(tx).to.be.reverted;
  });

  it("space 1-1", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[32],
      Owner: addrs[32],
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
      WalletAddr: addrs[32],
      Owner: addrs[32],
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
    await expect(tx).to.not.be.reverted;
    await expect(tx).to.emit(space, "SetUserSpaceEvent");
  });

  let value = 0;
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
    let wait = await res.wait();
    if (wait.events){
      if (wait.events.length > 0) {
        let event = wait.events[0];
        if (event.args) {
          if (event.args.length > 0) {
            let arg = event.args[0];
            value = arg.Value;
            // console.log(arg)
            expect(arg.Value).not.to.be.eq(0);
          }
        }
      }
    }
  });

  it("space 1-1", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[32],
      Owner: addrs[32],
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
