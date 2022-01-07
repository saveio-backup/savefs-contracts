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

  it("GetUserSpace", async () => {
    const res = fs.GetUserSpace("0x0000000000000000000000000000000000000000");
    expect(res).to.be.reverted;
  });
});
