import assert from "assert";
import { expect } from "chai";
import { addrs, pdp, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("submit", async () => {
    const tx = pdp.SubmitVerifyProofRequest(
      {
        Proof: {
          Version: 1,
          Proofs: [1],
          FileIds: [1],
          Tags: [1],
          RootHashes: [1]
        },
        State: true,
        LastUpdateHeight: 0,
      },
      [{ Index: 1, Rand: 1 }],
      [{
        PathLen: 1,
        Path: [{ Layer: 1, Index: 1, Hash: [1] }],
      }],
    )
    await expect(tx).to.not.be.reverted;
    let res = await (await tx).wait();
    if (res.events?.length == 1) {
      assert(res.events[0].event == "PDPVerifyEvent");
    }
  });

});
