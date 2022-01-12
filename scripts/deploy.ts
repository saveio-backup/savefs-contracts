import { ethers } from "hardhat";

async function main() {
  
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
