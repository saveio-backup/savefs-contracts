import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  let fs: FileSystem;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = await fs.deployed();
    assert(res != undefined)
  });

  it(scriptName, async () => {
    const res = fs.GetWhiteList([]);
    expect(res).to.be.reverted;

    const res2 = fs.GetWhiteList([1, 2, 3]);
    // because list is empty
    expect(res2).to.be.reverted;
  });
});
