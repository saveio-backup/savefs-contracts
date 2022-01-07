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

  it("GetFileInfo", async () => {
    const res = fs.GetFileInfo([65, 66, 67, 68, 69, 70]);
    expect(res).to.not.be.reverted;
  });
});
