import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Config } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {
  let config: Config;
  let fs: FileSystem;

  it("Deploy Config", async () => {
    const Config = await ethers.getContractFactory("Config");
    config = await Config.deploy();
    let res = config.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Deploy FileSystem", async () => {
    const FS = await ethers.getContractFactory("FileSystem");
    fs = await FS.deploy();
    let res = await fs.deployed();
    assert(res != undefined)
  });

  it(scriptName, async () => {
    let tx = fs.initialize();
    expect(tx).to.not.be.reverted;
  });

});
