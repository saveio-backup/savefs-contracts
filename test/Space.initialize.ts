import { expect, assert } from "chai";
import { ethers } from "hardhat";
import { Space } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {
  let space: Space;

  it("Deploy Space", async () => {
    const Space = await ethers.getContractFactory("Space");
    space = await Space.deploy();
    let res = space.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Space initialize", async () => {
    let tx = space.initialize();
    expect(tx).to.not.be.reverted;
  });

});
