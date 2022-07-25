import { expect } from "chai";
import { ContractTransaction } from "ethers";
import { ethers } from "hardhat";
import {
  File, Node, Config, Sector, Space, List, Prove, PDP, FileExtra, ProveExtra
} from "../typechain";
import { randomBytes } from "crypto";

var path = require('path');
var name = path.basename(__filename);

let addrs: Array<string> = [];
let config: Config;
let node: Node;
let file: File;
let fileExtra: FileExtra;
let sector: Sector;
let space: Space;
let list: List;
let prove: Prove;
let proveExtra: ProveExtra;
let pdp: PDP;


var print = async (tx: Promise<ContractTransaction>) => {
  let res = await (await tx).wait();
  console.log(res)
}

describe(name, function () {
  it("Get addrs", async () => {
    const accounts = await ethers.getSigners();
    for (const account of accounts) {
      addrs.push(account.address)
    }
  })

  // ----------

  it("Deploy Config", async () => {
    const Config = await ethers.getContractFactory("Config");
    config = await Config.deploy();
    let res = config.deployed();
    // console.log(await res)
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Node", async () => {
    const Node = await ethers.getContractFactory("Node");
    node = await Node.deploy();
    let res = node.deployed();
    // console.log(await res)
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Sector", async () => {
    const Sector = await ethers.getContractFactory("Sector");
    sector = await Sector.deploy();
    let res = sector.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Space", async () => {
    const Space = await ethers.getContractFactory("Space");
    space = await Space.deploy();
    let res = space.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy File", async () => {
    const File = await ethers.getContractFactory("File");
    file = await File.deploy();
    let res = file.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy FileExtra", async () => {
    const FileExtra = await ethers.getContractFactory("FileExtra");
    fileExtra = await FileExtra.deploy();
    let res = file.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy List", async () => {
    const List = await ethers.getContractFactory("List");
    list = await List.deploy();
    let res = list.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy Prove", async () => {
    const Prove = await ethers.getContractFactory("Prove");
    prove = await Prove.deploy();
    let res = prove.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy ProveExtra", async () => {
    const ProveExtra = await ethers.getContractFactory("ProveExtra");
    proveExtra = await ProveExtra.deploy();
    let res = proveExtra.deployed();
    await expect(res).not.to.be.reverted;
  });

  it("Deploy PDP", async () => {
    const PDP = await ethers.getContractFactory("PDP");
    pdp = await PDP.deploy();
    let res = pdp.deployed();
    await expect(res).not.to.be.reverted;
  });

  // ----------

  it("initialize Node", async () => {
    let tx = node.initialize(config.address, sector.address);
    await expect(tx).not.to.be.reverted;
  });

  it("initialize Sector", async () => {
    let tx = sector.initialize(node.address, {
      SECTOR_FILE_INFO_GROUP_MAX_LEN: 5000
    });
    await expect(tx).not.to.be.reverted;
  });

  it("initialize File", async () => {
    let tx = file.initialize(
      config.address, node.address, space.address, sector.address, prove.address,
      {
        DEFAULT_BLOCK_INTERVAL: 5,
        DEFAULT_PROVE_PERIOD: 3600 * 24,
        IN_SECTOR_SIZE: 1000 * 1000
      },
      fileExtra.address
    );
    await expect(tx).not.to.be.reverted;
  });

  it("initialize Space", async () => {
    let tx = space.initialize(config.address, file.address);
    await expect(tx).not.to.be.reverted;
  });

  it("initialize Prove", async () => {
    let tx = prove.initialize(
      config.address,
      file.address,
      node.address,
      pdp.address,
      sector.address,
      {
        SECTOR_PROVE_BLOCK_NUM: 32
      },
      proveExtra.address
    );
    await expect(tx).not.to.be.reverted;
  });

});


var mine = async (n: number) => {
  for (let index = 0; index < n; index++) {
    await ethers.provider.send('evm_mine', []);
  }
}

let files: Buffer[] = [];

; (() => {
  for (let index = 0; index < 100; index++) {
    files.push(randomBytes(32))
  }
})();

export {
  files,
  addrs,
  config,
  node,
  file,
  sector,
  space,
  list,
  prove,
  pdp,
  print,
  mine
};
