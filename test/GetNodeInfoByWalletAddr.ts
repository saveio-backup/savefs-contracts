import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem } from "../typechain";

describe("FileSystem", () => {
  let fs: FileSystem;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = await fs.deployed();
    assert(res != undefined)
  });

  it("GetNodeInfoByWalletAddr", async () => {
    let tx = fs.NodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    }, {
      value: 1000000
    });
    await (await tx).wait();

    const res = await fs.GetNodeInfoByWalletAddr("0x5FbDB2315678afecb367f032d93F642f64180aa3");
    assert(res.WalletAddr == "0x5FbDB2315678afecb367f032d93F642f64180aa3");
  });
});
