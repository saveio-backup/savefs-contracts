import { ethers, upgrades } from "hardhat";

function writeToFile(fileName: string, data: any) {
  const fs = require('fs');
  fs.writeFile(fileName, data, (err: any) => {
    if (err) {
      console.error(err);
      return;
    };
    console.log("File has been created");
  });
}

function getNetworkName() {
  const network = require('hardhat').network.name;
  return network;
}

async function main() {

  const Config = await ethers.getContractFactory("Config");
  let config = await upgrades.deployProxy(Config);
  await config.deployed();
  console.log("Config deployed to:", config.address);

  const Node = await ethers.getContractFactory("Node");
  let node = await upgrades.deployProxy(Node, [], { initializer: false });
  await node.deployed();
  console.log("Node deployed to:", node.address);

  const Sector = await ethers.getContractFactory("Sector");
  let sector = await upgrades.deployProxy(Sector, [], { initializer: false });
  await sector.deployed();
  console.log("Sector deployed to:", sector.address);

  const Space = await ethers.getContractFactory("Space");
  let space = await upgrades.deployProxy(Space, [], { initializer: false });
  await space.deployed();
  console.log("Space deployed to:", space.address);

  const File = await ethers.getContractFactory("File");
  let file = await upgrades.deployProxy(File, [], { initializer: false });
  await file.deployed();
  console.log("File deployed to:", file.address);

  const FileExtra = await ethers.getContractFactory("FileExtra");
  let fileExtra = await upgrades.deployProxy(FileExtra, [], { initializer: false });
  await fileExtra.deployed();
  console.log("fileExtra deployed to:", fileExtra.address);

  const List = await ethers.getContractFactory("List");
  let list = await upgrades.deployProxy(List, [], { initializer: false });
  await list.deployed();
  console.log("List deployed to:", list.address);

  const Prove = await ethers.getContractFactory("Prove");
  let prove = await upgrades.deployProxy(Prove, [], { initializer: false });
  await prove.deployed();
  console.log("Prove deployed to:", prove.address);

  const ProveExtra = await ethers.getContractFactory("ProveExtra");
  let proveExtra = await upgrades.deployProxy(ProveExtra, [], { initializer: false });
  await proveExtra.deployed();
  console.log("proveExtra deployed to:", proveExtra.address);

  const PDP = await ethers.getContractFactory("PDP");
  let pdp = await upgrades.deployProxy(PDP, [], { initializer: false });
  await pdp.deployed();
  console.log("PDP deployed to:", pdp.address);

  await node.initialize(config.address, sector.address);
  await sector.initialize(
    node.address,
    {
      SECTOR_FILE_INFO_GROUP_MAX_LEN: 5000
    }
  );
  await file.initialize(
    config.address,
    node.address,
    space.address,
    sector.address,
    prove.address,
    {
      DEFAULT_BLOCK_INTERVAL: 5,
      DEFAULT_PROVE_PERIOD: 3600 * 24,
      IN_SECTOR_SIZE: 1000 * 1000
    },
    fileExtra.address
  );
  await space.initialize(config.address, file.address);
  await prove.initialize(
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
  await pdp.initialize(file.address, sector.address);
  console.log("Initialize finished");

  let log = ''
  log += "------------golang----------"
  log +=
    `
var ConfigAddress = ethCommon.HexToAddress("${config.address}")
var NodeAddress = ethCommon.HexToAddress("${node.address}")
var SectorAddress = ethCommon.HexToAddress("${sector.address}")
var SpaceAddress = ethCommon.HexToAddress("${space.address}")
var FileAddress = ethCommon.HexToAddress("${file.address}")
var FileExtraAddress = ethCommon.HexToAddress("${fileExtra.address}")
var ListAddress = ethCommon.HexToAddress("${list.address}")
var ProveAddress = ethCommon.HexToAddress("${prove.address}")
var ProveExtraAddress = ethCommon.HexToAddress("${proveExtra.address}")
var PDPAddress = ethCommon.HexToAddress("${pdp.address}")
    `

  log += "------------javascript----------"
  log +=
    `
let configAddress = "${config.address}"
let nodeAddress = "${node.address}"
let sectorAddress = "${sector.address}"
let spaceAddress = "${space.address}"
let fileAddress = "${file.address}"
let fileExtraAddress = "${fileExtra.address}"
let listAddress = "${list.address}"
let proveAddress = "${prove.address}"
let proveExtraAddress = "${proveExtra.address}"
let pdpAddress = "${pdp.address}"
    `

  console.log(log)

  let network = getNetworkName()
  writeToFile(`./deploy/.address-${network}.txt`, log)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
