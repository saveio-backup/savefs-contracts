import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
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
