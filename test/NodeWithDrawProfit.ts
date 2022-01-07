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
    await fs.deployed();
  });

  it(scriptName, async () => {
    let tx = fs.NodeRegister({
      Pledge: 2,
      Profit: 2,
      Volume: 1000 * 1000,
      RestVol: 3,
      ServiceTime: 4,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    }, {
      value: 1000000
    });
    let r = await (await tx).wait();
    // console.log(r)

    const res = fs.NodeWithDrawProfit("0x5FbDB2315678afecb367f032d93F642f64180aa3");
    // must be recerted because profit is 0
    expect(res).to.be.reverted;
  });
});
