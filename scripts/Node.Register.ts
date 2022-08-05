import { provider, address, format, ethers, nodeAddress, signer } from "./config";

(async function main() {
  const Node = await ethers.getContractFactory("Node");
  let node = new ethers.Contract(nodeAddress, Node.interface, signer);

  let tx = node.Register({
    Pledge: 0,
    Profit: 0,
    Volume: 1000 * 1000,
    RestVol: 0,
    ServiceTime: 0,
    WalletAddr: address,
    NodeAddr: new Array(),
  }, {
    value: 1000000
  });

  let res = await tx;
  console.log(res)

})();

