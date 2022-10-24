import { ethers, upgrades } from "hardhat";

async function main() {

  let configAddress = "0x3ad4Aa72049FA58FCAD2bb462e530B4314935935"
  let nodeAddress = "0xff5C50e7080a5e116d0EF8F11e519Ac9De4EC2a7"
  let sectorAddress = "0x251201a0CDc5e011b35FA39a84BCA284Dcefc8d4"
  let spaceAddress = "0x1246B61f24CEAcAD9bAdE4a07b71248F4905371a"
  let fileAddress = "0x83AFEFFd26CF0ba8BF1AfF5605dDa7812f6AEF10"
  let fileExtraAddress = "0xF3D9D2E306dA5C2c15A1e3B5006e70f952E61d77"
  let listAddress = "0xfe56D4853f006f168ef746727555323026176e2E"
  let proveAddress = "0x9111f68a213909eF7aa62C1517F77530a8610329"
  let proveExtraAddress = "0x45771759191173D99C9b35931a716732aA9037d1"
  let pdpAddress = "0xDb96e5bbC0785A22Ef73F835C59F4424f38A15E0"

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
