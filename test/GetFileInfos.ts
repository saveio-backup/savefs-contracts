import { expect, assert } from "chai";
import { BigNumber } from "ethers";
import { isBytesLike } from "ethers/lib/utils";
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

  it("GetFileInfos", async () => {
    const res = fs.GetFileInfos({
      FileNum: 0,
      List: []
    });
    expect(res).to.be.reverted;

    const res2 = fs.GetFileInfos({
      FileNum: 1,
      List: [
        [65, 66, 67, 68, 69, 70]
      ]
    });
    expect(res2).to.be.reverted;
  });
});
