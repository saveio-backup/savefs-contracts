import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { Node, Config, Sector } from "../typechain";
import { config, fs, node, space, sector } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {

  it(scriptName, async () => {
    const tx = sector.GetSectorInfos("0xdD2FD4581271e230360230F9337D5c0430Bf44C0");
    let res = await tx;
    // console.log(res)
    assert(res.length == 0);

    const tx2 = node.NodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0xdD2FD4581271e230360230F9337D5c0430Bf44C0",
      NodeAddr: "0xdD2FD4581271e230360230F9337D5c0430Bf44C0",
    },
      {
        value: 1000000
      }
    );
    let s = await (await tx2).wait();
    // console.log(s)

    const tx3 = sector.CreateSector({
      NodeAddr: "0xdD2FD4581271e230360230F9337D5c0430Bf44C0",
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
    let r3 = await (await tx3).wait()
    // console.log(r3)

    const tx4 = sector.GetSectorInfos("0xdD2FD4581271e230360230F9337D5c0430Bf44C0");
    let res4 = await tx4;
    assert(res4.length == 1);
    // console.log(res4)
  });
});
