import { ethers, upgrades } from "hardhat";

async function main() {

  const Space = await ethers.getContractFactory("Space");
  let space = await upgrades.deployProxy(Space, [], { initializer: false });
  let res = await space.deployed();
  console.log("Space deployed to:", space.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
