import { ethers, upgrades } from "hardhat";

async function main() {

  let configAddress = "0x1463bFB7d7F7c5A45B2502F903F8C3364988d827"
  let nodeAddress = "0x7029e8faC0de5B6b50B12951e9c274dE9D91Ca8d"
  let sectorAddress = "0xD814d6F57Dd7b767ADe6537BF84F442904084E25"
  let spaceAddress = "0x8cf9E9cB187B9ED64113C7852941992499Be2789"
  let fileAddress = "0x1B691dA3fAD45aEc92319025959585D2392B3440"
  let fileExtraAddress = "0xE2B1685000185C1bc5F5eDD5e34d3Ba00FbA365e"
  let listAddress = "0xEB1ee9ff9cAab2259C93b4b75ddc788b72C0A0BC"
  let proveAddress = "0x8A56564ecb7223BFC6dc9F0DB4c7198292108220"
  let proveExtraAddress = "0x9Bc985d31C849D2307344a8199D04B0D6D33c56E"
  let pdpAddress = "0x57AbC9065B6E9163D25D9d30DE2ab9D2dec8f526"
  let pdpExtraAddress = "0xa85fFeA4C9E062124B46a0D817f5bE0b87b8b334"

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

  const PDPExtra = await ethers.getContractFactory("PDPExtra");
  let pdpExtra = await upgrades.upgradeProxy(pdpExtraAddress, PDPExtra);
  console.log("PDPExtra upgraded");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
