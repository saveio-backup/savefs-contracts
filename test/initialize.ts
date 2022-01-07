import { assert } from "console";
import { ethers } from "hardhat";
import { FileSystem } from "../typechain";

describe("FileSystem", function () {
  let fs: FileSystem;

  it("Deploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = await fs.deployed();
    assert(res != undefined)
  });

  it("initialize", async () => {
    let tx = fs.initialize();
    await (await tx).wait();
  });

});
