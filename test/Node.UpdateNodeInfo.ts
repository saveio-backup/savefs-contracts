import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Config, Node } from "../typechain";
import { addrs, config, fs, node, space } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(`${scriptName} register`, async () => {
    let tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[25],
      NodeAddr: addrs[25],
    }, {
      value: 1000000
    });
    await (await tx).wait();
  });

  it(`${scriptName} get 1`, async () => {
    let tx = node.GetNodeInfoByNodeAddr(addrs[25])
    let res = await tx;
    // console.log(res)
    assert(res.Volume.eq(1000 * 1000));
  });

  it(`${scriptName} update info`, async () => {
    let tx = node.UpdateNodeInfo({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1001,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[25],
      NodeAddr: addrs[25],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it(`${scriptName} get 2`, async () => {
    let tx = node.GetNodeInfoByNodeAddr(addrs[25])
    let res = await tx;
    // console.log(res)
    assert(res.Volume.eq(1000 * 1001));
  });
});
