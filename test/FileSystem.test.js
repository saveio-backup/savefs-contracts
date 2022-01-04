const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FileSystem", () => {

  let fs;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    await fs.deployed();
  })

  it("FsGetSettings", async () => {
    const setting = await fs.FsGetSettings();
  });

  it("FsNodeRegister", async () => {
    const res = fs.FsNodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 0,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    });
    expect(res).to.be.reverted;

    const res2 = fs.FsNodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
      NodeAddr: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    });
    expect(res2).to.not.be.reverted;
  });
});
