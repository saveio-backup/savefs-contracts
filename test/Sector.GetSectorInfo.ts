import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Sector } from "../typechain";
import { addrs, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  it(`${scriptName} 1`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[12],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.Size.eq(0));
  })

  it(`${scriptName} 2`, async () => {
    const tx1 = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[12],
      NodeAddr: addrs[12],
    },
      {
        value: 1000000
      }
    );
    expect(tx1).to.not.be.reverted;
  })

  it(`${scriptName} 3`, async () => {
    const tx2 = sector.CreateSector({
      NodeAddr: addrs[12],
      SectorID: 1024,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 1,
      GroupNum: 1014,
      IsPlots: false,
      FileList: []
    });
    expect(tx2).to.not.be.reverted;
  })

  it(`${scriptName} 4`, async () => {
    const tx = sector.GetSectorInfo({
      NodeAddr: addrs[12],
      SectorId: 1024,
    });
    let res = await tx;
    // console.log(res)
    assert(res.Size.eq(1));
  });
});
