import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem } from "../typechain";
import { addrs, config, fs, node, space, list } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get empty", async () => {
    const tx = list.GetWhiteList([1, 2, 3, 4]);
    await expect(tx).to.be.reverted; 
  });

  it("add", async () => {
    const tx = list.WhiteListOperate({
      FileHash: [1, 2, 3, 4],
      Op: 0,
      List: [
        {
          Addr: addrs[38],
          BaseHeight: 1,
          ExpireHeight: 100,
        }
      ]
    });
    await expect(tx).not.to.be.reverted;
  });

  it("get not empty", async () => {
    const tx = list.GetWhiteList([1, 2, 3, 4]);
    let res = await tx;
    // console.log(res);
    assert(res.length > 0);
  })
});
