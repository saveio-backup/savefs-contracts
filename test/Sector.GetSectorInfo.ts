import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { FileSystem, Node, Config, Sector } from "../typechain";
import { config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {

    const tx0 = sector.GetSectorInfo({
      NodeAddr: "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
      SectorId: "1",
    });
    let res0 = await tx0;
    // console.log(res0)
    assert(res0.Size.eq(0));

    const tx2 = node.NodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
      NodeAddr: "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
    },
      {
        value: 1000000
      }
    );
    let s = await (await tx2).wait();
    // console.log(s)

    const tx = sector.CreateSector({
      NodeAddr: "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
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
      FileList_: {
        FileNum: 0,
        List: []
      }
    });
    let r2 = await (await tx).wait()
    // console.log(r2)

    const tx3 = sector.GetSectorInfo({
      NodeAddr: "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
      SectorId: "1",
    });
    let res3 = await tx3;
    assert(res3.Size.eq(1));
  });
});
