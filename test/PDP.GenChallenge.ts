import assert from "assert";
import { expect } from "chai";
import { addrs, pdp, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("gen challenge", async () => {
    let tx = pdp.GenChallenge({
        WalletAddr: addrs[0],
        HashValue: "0x1234567890",
        FileBlockNum: 1,
        ProveNum: 1,
    });
    let res = await tx;
    // console.log(res);
  });

});
