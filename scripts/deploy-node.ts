import { ethers, upgrades } from "hardhat";

async function main() {

  const Node = await ethers.getContractFactory("Node");
  let node = await upgrades.deployProxy(Node, [], { initializer: false });
  let res = await node.deployed();
  console.log("Node deployed to:", node.address);

  // console.log(res.deployTransaction)
  console.log(res.deployTransaction.gasLimit)
  console.log(res.deployTransaction.gasPrice)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
