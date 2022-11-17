import { ethers, upgrades } from "hardhat";

async function main() {

  const Test = await ethers.getContractFactory("Test");
  let test = await Test.deploy();
  let res = await test.deployed();
  console.log("Test deployed to:", test.address);

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
