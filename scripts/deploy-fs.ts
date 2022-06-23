import { ethers, upgrades } from "hardhat";

async function main() {

  const FS = await ethers.getContractFactory("FileSystem");
  let fs = await upgrades.deployProxy(FS, [], { initializer: false });
  let res = await fs.deployed();
  console.log("FS deployed to:", fs.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
