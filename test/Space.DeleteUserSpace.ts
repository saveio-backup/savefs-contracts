import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { addrs, print, config, fs, node, space, sector, mine } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("space 1-1", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[32],
      Owner: addrs[32],
      Size: {
        Type: 1,
        Value: 1024
      },
      BlockCount: {
        Type: 1,
        Value: (3600 * 24) / 5
      }
    }, {
      value: 1251
    });
    // await print(tx);
    await expect(tx).to.not.be.reverted;
    await expect(tx).to.emit(space, "SetUserSpaceEvent");
  });

  it("delete", async () => {
    await mine(17294 * 2);
    const tx = space.DeleteUserSpace(addrs[32]);
    // print(tx)
    await expect(tx).to.not.be.reverted;
  });
  
});
