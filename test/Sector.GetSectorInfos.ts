import { expect, assert } from "chai";
import { ethers, network } from "hardhat";
import { Node, Config, Sector } from "../typechain";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, () => {
  let config: Config;
  let node: Node;
  let sector: Sector;

  it("Deploy Config", async () => {
    const Config = await ethers.getContractFactory("Config");
    config = await Config.deploy();
    let res = config.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Deploy Node", async () => {
    const Node = await ethers.getContractFactory("Node");
    node = await Node.deploy();
    let res = node.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Node initialize", async () => {
    let tx = node.initialize(config.address);
    expect(tx).to.not.be.reverted;
  });

  it("Deploy Sector", async () => {
    const S = await ethers.getContractFactory("Sector");
    sector = await S.deploy();
    let res = sector.deployed();
    expect(res).to.not.be.reverted;
  });

  it("Sector initialize", async () => {
    let tx = sector.initialize(node.address);
    expect(tx).to.not.be.reverted;
  });
  // -------------------------------------------------------

  it(scriptName, async () => {
    const tx = sector.GetSectorInfos("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266");
    let res = await tx;
    assert(res.length == 0);

    const tx2 = node.NodeRegister({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
      NodeAddr: "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
    },
      {
        value: 1000000
      }
    );
    let s = await (await tx2).wait();
    // console.log(s)

    const tx3 = sector.CreateSector({
      NodeAddr: "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
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
    // console.log(r2)

    const tx4 = sector.GetSectorInfos("0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199");
    let res4 = await tx4;
    assert(res4.length == 1);
  });
});
