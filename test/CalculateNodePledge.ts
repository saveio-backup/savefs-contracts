import { expect, assert } from "chai";
import { BigNumber } from "ethers";
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
    const tx = fs.CalculateNodePledge({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    });
    let res = await tx;
    assert(res.eq(1000000));
  });
});
