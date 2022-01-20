import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { addrs, print, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("add", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[40],
      Owner: addrs[40],
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

  it("delete", async () => {
    const tx = space.DeleteUserSpace(addrs[40]);
    // print(tx)
    await expect(tx).to.be.reverted;
  });
});
