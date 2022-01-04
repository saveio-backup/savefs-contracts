const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FileSystem", function () {
  it("Should get default setting", async function () {
    const FS = await ethers.getContractFactory("FileSystem");
    const fs = await FS.deploy();
    await fs.deployed();

    const setting = await fs.GetSettings();
    expect(setting).to.not.equal("");
  });
});
