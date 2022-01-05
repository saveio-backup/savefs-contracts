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

  it("FsGetSettings", async () => {
    const setting = await fs.FsGetSettings();
    expect(setting).to.not.equal(null);
  });
});
