import { assert, expect } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("delete unsettled files", async () => {
    const tx = fs.DeleteUnsettledFiles(addrs[10])
    await expect(tx).to.not.be.reverted;
  });

});
