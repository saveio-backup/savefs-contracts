import { ethers } from "hardhat";

async function main() {

  const Config = await ethers.getContractFactory("Config");
  let config = await Config.deploy();
  await config.deployed();
  console.log("Config deployed to:", config.address);

  const Node = await ethers.getContractFactory("Node");
  let node = await Node.deploy();
  await node.deployed();
  console.log("Node deployed to:", node.address);

  const Sector = await ethers.getContractFactory("Sector");
  let sector = await Sector.deploy();
  await sector.deployed();
  console.log("Sector deployed to:", sector.address);

  const Space = await ethers.getContractFactory("Space");
  let space = await Space.deploy();
  await space.deployed();
  console.log("Space deployed to:", space.address);

  const FS = await ethers.getContractFactory("FileSystem");
  let fs = await FS.deploy();
  await fs.deployed();
  console.log("FS deployed to:", fs.address);

  const List = await ethers.getContractFactory("List");
  let list = await List.deploy();
  await list.deployed();
  console.log("List deployed to:", list.address);

  const Prove = await ethers.getContractFactory("Prove");
  let prove = await Prove.deploy();
  await prove.deployed();
  console.log("Prove deployed to:", prove.address);

  const PDP = await ethers.getContractFactory("PDP");
  let pdp = await PDP.deploy();
  await pdp.deployed();
  console.log("PDP deployed to:", pdp.address);

  await node.initialize(config.address, sector.address);
  await sector.initialize(node.address);
  await fs.initialize(config.address, node.address, space.address, sector.address, prove.address);
  await space.initialize(config.address, fs.address);
  await prove.initialize(config.address, fs.address, node.address, pdp.address, sector.address);
  console.log("Initialize finished");

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
