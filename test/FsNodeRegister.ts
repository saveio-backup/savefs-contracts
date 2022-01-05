import { expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem } from "../typechain";

describe("FileSystem", () => {
  let fs: FileSystem;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    await fs.deployed();
  });

  it("FsNodeRegister event", async () => {
    const tx = await fs.FsNodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    });
    let res = tx.wait();
    // except(tx.logs[0].event).to.equal("FsNodeRegister");
    // console.log(await res)
  })

  it("FsNodeRegister", async () => {
    const tx = fs.FsNodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 0,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    });
    expect(tx).to.be.reverted;

    const tx2 = fs.FsNodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    });
    expect(tx2).to.not.be.reverted;
  });
});
