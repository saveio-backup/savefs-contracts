const { provider, address, format } = require("./config.js");
import { ethers } from "hardhat";

let contractAddress = "0x0FCc3b2a1Fe79DAefF2DE1Fdc17D9cd3aEaf6d3f";

(async function main() {
  const Node = await ethers.getContractFactory("Node");
  let node = new ethers.Contract(contractAddress, Node.interface, provider);

  const tx = node.CalculateNodePledge({
    Pledge: 0,
    Profit: 0,
    Volume: 1000 * 1000,
    RestVol: 0,
    ServiceTime: 0,
    WalletAddr: address,
    NodeAddr: new Array(),
  });
  let res = await tx;
  format(res);

})();
