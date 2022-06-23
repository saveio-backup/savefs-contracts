import { ethers, upgrades } from "hardhat";

async function main() {

  const Sector = await ethers.getContractFactory("Sector");
  let sector = await upgrades.deployProxy(Sector, [], { initializer: false });
  let res = await sector.deployed();
  console.log("Sector deployed to:", sector.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
