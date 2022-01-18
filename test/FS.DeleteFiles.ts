import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res = fs.DeleteFiles([
      [65, 66, 67, 68, 69, 70],
      [65, 66, 67, 68, 69, 71],
    ]);
    expect(res).to.not.be.reverted;
  });
});
