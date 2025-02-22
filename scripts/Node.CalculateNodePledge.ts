import { provider, address, format, ethers, nodeAddress } from "./config";


(async function main() {
  const Node = await ethers.getContractFactory("Node");
  let node = new ethers.Contract(nodeAddress, Node.interface, provider);

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
