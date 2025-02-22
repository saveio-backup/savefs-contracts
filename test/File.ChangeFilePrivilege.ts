import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { config, file, node, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("change file privilege", async () => {
    const tx = file.ChangeFilePrivilege({
      fileHash: [65, 66, 67, 68, 69, 70],
      privilege: 1
    });
    await expect(tx).to.not.be.reverted;
  });

});
