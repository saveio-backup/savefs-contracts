import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("change file privilege", async () => {
    const tx = fs.ChangeFilePivilege({
      fileHash: [65, 66, 67, 68, 69, 70],
      privilege: 1
    });
    await expect(tx).to.not.be.reverted;
  });

});
