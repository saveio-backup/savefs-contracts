import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem } from "../typechain";

describe("FileSystem", () => {
  let fs: FileSystem;

  it("FsDeploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    await fs.deployed();
  });

  it("FsNodeQuery", async () => {
    const res = await fs.FsNodeQuery("0x2546BcD3c84621e976D8185a91A922aE77ECEc30");
    assert(res.Volume.eq(0))
  });
});
