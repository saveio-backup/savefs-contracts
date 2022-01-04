const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FileSystem", () => {
  let fs;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    await fs.deployed();
  });

  it("FsGetSettings", async () => {
    const setting = await fs.FsGetSettings();
    expect(setting).to.not.equal(null);
  });
});
