import { ethers, upgrades } from "hardhat";

async function main() {

  const List = await ethers.getContractFactory("List");
  let list = await upgrades.deployProxy(List, [], { initializer: false });
  let res = await list.deployed();
  console.log("List deployed to:", list.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
