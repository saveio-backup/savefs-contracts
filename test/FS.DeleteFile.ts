import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { FileSystem, Node, Config, Space } from "../typechain";
import { files, addrs, config, fs, node, space, print, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("create space", async () => {
    const tx = space.ManageUserSpace({
      WalletAddr: addrs[67],
      Owner: addrs[67],
      Size: {
        Type: 1,
        Value: 1
      },
      BlockCount: {
        Type: 1,
        Value: (3600 * 24) / 5
      }
    });
    // await print(tx);
    await expect(tx).to.not.be.reverted;
    await expect(tx).to.emit(space, "SetUserSpaceEvent");
  });

  it("store file to space", async () => {
    const tx = fs.StoreFile({
      FileHash: files[0],
      FileOwner: addrs[67],
      FileDesc: [],
      Privilege: 1,
      FileBlockNum: 1,
      FileBlockSize: 1,
      ProveInterval: 1,
      ProveTimes: 1,
      ExpiredHeight: 10000000,
      CopyNum: 1,
      Deposit: 1,
      FileProveParam: [],
      ProveBlockNum: 1,
      BlockHeight: 1,
      ValidFlag: false,
      BlocksRoot: [],
      StorageType_: 1,
      RealFileSize: 1,
      PrimaryNodes: [],
      CandidateNodes: [],
      ProveLevel_: 1,
      IsPlotFile: false,
      PlotInfo_: {
        NumberID: 1,
        StartNonce: 1,
        Nonces: 1,
      }
    }, {
      value: 1000000
    })
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("delete file fail because space used = 0", async () => {
    const tx = fs.DeleteFile(files[0]);
    // await print(tx)
    await expect(tx).to.be.reverted;
  });

  it("node register success", async () => {
    const tx = node.Register({
      Pledge: 0,
      Profit: 0,
      Volume: 1000 * 1000,
      RestVol: 0,
      ServiceTime: 0,
      WalletAddr: addrs[67],
      NodeAddr: addrs[67],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it("create sector", async () => {
    const tx = sector.CreateSector({
      NodeAddr: addrs[67],
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
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("store file to sector", async () => {
    const tx = fs.StoreFile({
      FileHash: files[1],
      FileOwner: addrs[67],
      FileDesc: [],
      Privilege: 1,
      FileBlockNum: 0,
      FileBlockSize: 0,
      ProveInterval: 1,
      ProveTimes: 1,
      ExpiredHeight: 100,
      CopyNum: 0,
      Deposit: 0,
      FileProveParam: [],
      ProveBlockNum: 1,
      BlockHeight: 1,
      ValidFlag: false,
      BlocksRoot: [],
      StorageType_: 1,
      RealFileSize: 1,
      PrimaryNodes: [],
      CandidateNodes: [],
      ProveLevel_: 1,
      IsPlotFile: false,
      PlotInfo_: {
        NumberID: 1,
        StartNonce: 1,
        Nonces: 1,
      }
    }, {
      value: 1000000
    })
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

  it("delete file success", async () => {
    const tx = fs.DeleteFile(files[1]);
    // await print(tx)
    await expect(tx).to.not.be.reverted;
  });

});
