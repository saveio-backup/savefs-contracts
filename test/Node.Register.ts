import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { addrs, config, fs, node, space, print } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it("register failed because not pay enough", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[0],
      NodeAddr: addrs[0],
    });
    // print(tx)
    expect(tx).to.be.reverted;
  });

  it("register success", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[0],
      NodeAddr: addrs[0],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it("event", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[1],
      NodeAddr: addrs[1],
    }, {
      value: 1000000
    });
    let res = await (await tx).wait();
    // console.log(res)
    if (res.events?.length == 1) {
      assert(res.events[0].event == "RegisterNodeEvent");
    }
  })

});
