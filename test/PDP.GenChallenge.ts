import assert from "assert";
import {expect} from "chai";
import {addrs, pdp, print} from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

    it("gen challenge", async () => {
        let tx = pdp.GenChallenge({
            WalletAddr: addrs[0],
            HashValue: "0x8a76d03b40c30b3ee7f3904b93d1405e7db0f68f1b46e8b97c8c4eaa19b8725a",
            FileBlockNum: 1,
            ProveNum: 1,
        });
        let res = await tx;
        // console.log(res);
        // console.log(addrs[0]);
        assert(res.length > 0);
    });

});
