const { ethers, upgrades } = require("hardhat");

async function main() {
  // Deploying
  const Box = await ethers.getContractFactory("Greeter");
  const instance = await upgrades.deployProxy(Box, [42]);
  await instance.deployed();

  // Upgrading
  const BoxV2 = await ethers.getContractFactory("Greeter");
  const upgraded = await upgrades.upgradeProxy(instance.address, BoxV2);
}

main();