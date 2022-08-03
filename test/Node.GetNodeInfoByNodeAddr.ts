import { expect, assert } from "chai";
import { hexlify } from "ethers/lib/utils";
import { addrs, node, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  const nodeAddr = new TextEncoder().encode("tcp://127.0.0.1:10001")

  it("register", async () => {
    let tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[6],
      NodeAddr: nodeAddr,
    }, {
      value: 1000000
    });
    // print(tx)
    expect(tx).to.not.be.reverted;
  });

  it("get", async () => {
    const res = await node.GetNodeInfoByNodeAddr(nodeAddr);
    assert(res.NodeAddr == hexlify(nodeAddr));
  });

});
