import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { Node, Config, Sector } from "../typechain";
import { addrs, config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);
var title = (s: string) => {
  return scriptName + " " + s;
}

describe(scriptName, () => {

  it(title("get 1"), async () => {
    const tx = sector.GetSectorsForNode(addrs[13]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 0);
  });

  it(title("register node"), async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[13],
      NodeAddr: addrs[13],
    },
      {
        value: 1000000
      }
    );
    // let s = await (await tx).wait();
    // console.log(s)
    expect(tx).to.not.be.reverted;
  });

  it(title("create sector"), async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[13],
      SectorID: 1,
      Size: 1,
      Used: 0,
      ProveLevel_: 1,
      FirstProveHeight: 1,
      NextProveHeight: 1,
      TotalBlockNum: 1,
      FileNum: 1,
      GroupNum: 1,
      IsPlots: false,
      FileList: []
    });
    // let r3 = await (await tx).wait()
    // console.log(r3)
    expect(tx).to.not.be.reverted;
  });

  it(title("get 2"), async () => {
    const tx = sector.GetSectorsForNode(addrs[13]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 1);
  });
});
