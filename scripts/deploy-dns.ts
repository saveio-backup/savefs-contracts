import { ethers, upgrades } from "hardhat";

async function main() {

  const Dns = await ethers.getContractFactory("Dns");
  let dns = await upgrades.deployProxy(Dns);
  let res = await dns.deployed();
  console.log("Dns deployed to:", dns.address);

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
