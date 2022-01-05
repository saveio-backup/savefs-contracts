import { ethers } from "hardhat";

describe("FileSystem", function () {
  let fs;

  it("FsDeploy", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    await fs.deployed();
  });
});
