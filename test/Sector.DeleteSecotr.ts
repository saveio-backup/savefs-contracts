import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Sector } from "../typechain";
import { addrs, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("delete sector", async () => {
    const tx = sector.DeleteSecotr({
      NodeAddr: addrs[70],
      SectorId: "100",
    });
    await expect(tx).to.not.be.reverted;
  });

});
