import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const res = space.ManageUserSpace({
      WalletAddr: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
      Owner: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
      Size: {
        Type: 0,
        Value: 0
      },
      BlockCount: {
        Type: 0,
        Value: 0
      }
    });
    expect(res).to.be.reverted; // params error
    // let tx = await (await res).wait();
    // console.log(tx)
  });
});
