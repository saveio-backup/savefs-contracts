import { ethers, upgrades } from "hardhat";

async function main() {

  let dnsAddress = '0x7cD105A603C15fe40A78dA1555051E8fBD950931'

  const Dns = await ethers.getContractFactory("Dns");
  let dns = await upgrades.upgradeProxy(dnsAddress, Dns);
  console.log("Dns upgraded");

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
