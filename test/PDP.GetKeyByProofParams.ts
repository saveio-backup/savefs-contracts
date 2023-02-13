import assert from "assert";
import { expect } from "chai";
import { addrs, pdp, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

    it("get key 1", async () => {
        let tx = pdp.GetKeyByProofParams({
            Version: 1,
            Proofs: [1],
            FileIds: [[1]],
            Tags: [[1]],
            RootHashes: [1],
        },
            [{ Index: 1, Rand: 1 }],
            [{
                PathLen: 1,
                Path: [{ Layer: 1, Index: 1, Hash: [1] }],
            }],
        )
        let res = await tx;
        // console.log(res);
        assert(res == "0xeb85fde30c2060675fffe325840d0c3621c099c64c2ce6c8307b6bf5019dca6d")
    });

    it("get key 2", async () => {
        let tx = pdp.GetKeyByProofParams({
            Version: 1,
            Proofs: [1],
            FileIds: [[2]],
            Tags: [[1]],
            RootHashes: [1],
        },
            [{ Index: 1, Rand: 1 }],
            [{
                PathLen: 1,
                Path: [{ Layer: 1, Index: 1, Hash: [1] }],
            }],
        )
        let res = await tx;
        // console.log(res);
        assert(res == "0xfd7bc8aa5d5e205b43c0e9ae827cb4d75704837bfbc8276f36c8977df54aa429")
    });

});
