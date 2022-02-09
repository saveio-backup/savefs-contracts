import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config } from "../typechain";
import { addrs, config, fs, node, space, sector, print } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("register", async () => {
    let tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[7],
      NodeAddr: addrs[7],
    }, {
      value: 1000000
    });
    // print(tx);
    await expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[7],
      SectorID: 1,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 10000,
      NextProveHeight: 10000,
      TotalBlockNum: 1,
      FileNum: 0,
      GroupNum: 1,
      IsPlots: false,
      FileList: []
    });
    await expect(tx).to.not.be.reverted;
  })

  it("cancel", async () => {
    let tx = node.Cancel(addrs[7]);
    // print(tx)
    await expect(tx).to.not.be.reverted;
  });

});
