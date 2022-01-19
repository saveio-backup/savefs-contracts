import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { name, addrs, config, fs, node, space } from "./initialize";

describe(name, () => {

  it("register", async () => {
    let tx = node.Register({
      Pledge: 1,
      Profit: 2,
      Volume: 1000 * 1000,
      RestVol: 3,
      ServiceTime: 4,
      WalletAddr: addrs[26],
      NodeAddr: addrs[26],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it("withdraw", async () => {
    const res = node.WithDrawProfit(addrs[26]);
    expect(res).to.be.reverted; // profit is 0
  });

  it("update", async () => {
    let tx = node.UpdateNodeInfo({
      Pledge: 1,
      Profit: 2,
      Volume: 1000 * 1000,
      RestVol: 3,
      ServiceTime: 4,
      WalletAddr: addrs[26],
      NodeAddr: addrs[26],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it("withdraw 2", async () => {
    const res = node.WithDrawProfit(addrs[26]);
    expect(res).to.not.be.reverted; // profit is not 0
  });
});
