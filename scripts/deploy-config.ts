import { ethers, upgrades } from "hardhat";

async function main() {

  const Config = await ethers.getContractFactory("Config");
  let config = await upgrades.deployProxy(Config);
  let res = await config.deployed();
  console.log("Config deployed to:", config.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
