import { ethers, upgrades } from "hardhat";

async function main() {

  const Test = await ethers.getContractFactory("Test");
  let test = await Test.deploy();
  let res = await test.deployed();
  console.log("Test deployed to:", test.address);

  const tx = test.Get();
  let res2 = await tx;
  console.log(res2)

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
