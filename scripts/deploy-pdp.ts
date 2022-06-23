import { ethers, upgrades } from "hardhat";

async function main() {

  const PDP = await ethers.getContractFactory("PDP");
  let pdp = await upgrades.deployProxy(PDP, [], { initializer: false });
  let res = await pdp.deployed();
  console.log("PDP deployed to:", pdp.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
