// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { ethers } from "hardhat";

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  // We get the contract to deploy
  const FS = await ethers.getContractFactory("FileSystem");
  const fs = await FS.deploy();
  await fs.deployed();
  console.log("FileSystem deployed to:", fs.address);

  let tx = fs.initialize();
  await (await tx).wait();

  let tx2 = fs.NodeRegister({
    Pledge: 0,
    Profit: 0,
    Volume: 1000 * 1000,
    RestVol: 0,
    ServiceTime: 0,
    WalletAddr: "0x8626f6940e2eb28930efb4cef49b2d1f2c9c1199",
    NodeAddr: "0x8626f6940e2eb28930efb4cef49b2d1f2c9c1199",
  }, {
    value: 1000000,
  });
  let res = await (await tx2).wait();
  console.log("NodeRegister done");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
