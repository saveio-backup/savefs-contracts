import { ethers, upgrades } from "hardhat";

async function main() {

  let configAddress = "0x786bd101fA16d0243d085f5850D5666154eB1277"
  let nodeAddress = "0x361eE58bbDFd0b320C063b900932A4E3dD9D3112"
  let sectorAddress = "0xa05Ad8cE3a256A05BEA6bB85CECFdBcb8EfEd001"
  let spaceAddress = "0x48b4d0e8b626c1fBcCB62258D939fd8D151BB89e"
  let fileAddress = "0x15083fBC432F7f5aC9F34Afddd441c614eB44639"
  let fileExtraAddress = "0x23f936c0d405c067309E2DA6bCb74A81a306E4d8"
  let listAddress = "0x599a1590D10A70D575C27e3ffa4ED145082AeF8C"
  let proveAddress = "0x1F86902045355e344e5F957F44341c239158f794"
  let proveExtraAddress = "0xa15Db8Bea785eF6771f2a20Ab8c9F6DA272706CD"
  let pdpAddress = "0x49065229EA79e183677CcfeDb006CD290b59ea30"

  const Config = await ethers.getContractFactory("Config");
  let config = await upgrades.upgradeProxy(configAddress, Config);
  console.log("Config upgraded");

  const Node = await ethers.getContractFactory("Node");
  let node = await upgrades.upgradeProxy(nodeAddress, Node);
  console.log("Node upgraded");

  const Sector = await ethers.getContractFactory("Sector");
  let sector = await upgrades.upgradeProxy(sectorAddress, Sector);
  console.log("Sector upgraded");

  const Space = await ethers.getContractFactory("Space");
  let space = await upgrades.upgradeProxy(spaceAddress, Space);
  console.log("Space upgraded");

  const File = await ethers.getContractFactory("File");
  let file = await upgrades.upgradeProxy(fileAddress, File);
  console.log("File upgraded");

  const FileExtra = await ethers.getContractFactory("FileExtra");
  let fileExtra = await upgrades.upgradeProxy(fileExtraAddress, FileExtra);
  console.log("fileExtra upgraded");

  const List = await ethers.getContractFactory("List");
  let list = await upgrades.upgradeProxy(listAddress, List);
  console.log("List upgraded");

  const Prove = await ethers.getContractFactory("Prove");
  let prove = await upgrades.upgradeProxy(proveAddress, Prove);
  console.log("Prove upgraded");

  const ProveExtra = await ethers.getContractFactory("ProveExtra");
  let proveExtra = await upgrades.upgradeProxy(proveExtraAddress, ProveExtra);
  console.log("ProveExtra upgraded");

  const PDP = await ethers.getContractFactory("PDP");
  let pdp = await upgrades.upgradeProxy(pdpAddress, PDP);
  console.log("PDP upgraded");

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
