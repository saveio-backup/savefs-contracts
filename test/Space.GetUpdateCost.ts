import { expect, assert } from "chai";
import { BigNumber } from "ethers";
import { ethers, network } from "hardhat";
import { addrs, print, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get update cost", async () => {
    const tx = space.GetUpdateCost({
      WalletAddr: addrs[50],
      Owner: addrs[50],
      Size: {
        Type: 1,
        Value: 1
      },
      BlockCount: {
        Type: 1,
        Value: (3600 * 24) / 5
      }
    });
    let res = await tx;
    // console.log(res)
    expect(tx).not.to.be.reverted;
  });
  
});
