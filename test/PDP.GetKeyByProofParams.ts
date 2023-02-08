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
      // Challenges: [{ Index: 1, Rand: 1 }],
      // MerklePath_: [{
      //   PathLen: 1,
      //   Path: [{ Layer: 1, Index: 1, Hash: [1] }],
      // }],
      RootHashes: [1],
    })
    let res = await tx;
    // console.log(res);
    // assert(res == "0x4309c3f69b948561504fd848880897b1ddc34f335338afa19c9180ef5e562d45")
  });

  it("get key 2", async () => {
    let tx = pdp.GetKeyByProofParams({
      Version: 1,
      Proofs: [1],
      FileIds: [[2]],
      Tags: [[1]],
      // Challenges: [{ Index: 1, Rand: 1 }],
      // MerklePath_: [{
      //   PathLen: 1,
      //   Path: [{ Layer: 1, Index: 1, Hash: [1] }],
      // }],
      RootHashes: [1],
    })
    let res = await tx;
    // console.log(res);
    // assert(res == "0xb9a5947c460c1c6acaa988611d74f1bbdd943494d33901a6add5c97b2a2939d3")
  });

});
