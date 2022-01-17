import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Sector } from "../typechain";
import { addrs, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {

    const tx = sector.DeleteSecotr({
      NodeAddr: addrs[15],
      SectorId: "100",
    });
    expect(tx).to.not.be.reverted;
    // let res = await tx;
    // console.log(res)
  });
});
