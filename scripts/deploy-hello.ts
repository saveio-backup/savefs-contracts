import { ethers, upgrades } from "hardhat";

async function main() {

  const Hello = await ethers.getContractFactory("Hello");
  let hello = await upgrades.deployProxy(Hello);
  let res = await hello.deployed();
  console.log("Hello deployed to:", hello.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
