import { ethers, upgrades } from "hardhat";

async function main() {

  const Prove = await ethers.getContractFactory("Prove");
  let prove = await upgrades.deployProxy(Prove, [], { initializer: false });
  let res = await prove.deployed();
  console.log("Prove deployed to:", prove.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
