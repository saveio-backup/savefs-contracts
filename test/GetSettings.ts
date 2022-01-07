import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { FileSystem } from "../typechain";

describe("FileSystem", () => {
  let fs: FileSystem;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = await fs.deployed();
    assert(res != undefined)
  });

  it("GetSettings", async () => {
    const setting = await fs.GetSetting();
    expect(setting).to.not.equal(null);
  });
});
