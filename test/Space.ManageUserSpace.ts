import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import {
  addrs, print, config, fs, node, space, sector
} from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("manage failed", async () => {
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

  it("manage add", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[32],
      Owner: addrs[32],
      Size: {
        Type: 1,
        Value: 1
      },
      BlockCount: {
        Type: 1,
        Value: 1
      }
    });
    // await print(tx);
    await expect(tx).to.not.be.reverted;
    await expect(tx).to.emit(space, "SetUserSpaceEvent");
  });

});
