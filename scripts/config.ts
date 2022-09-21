const { ethers } = require("hardhat");

const url = "http://127.0.0.1:8545";
const address = "0x0bd81846A232E263bC289A0B2206Ae4ae26A08A2";
const privateKey =
  "5e2434f0a266601bc7b03313dfada0d6bb7706aa78723c01df0e7bbf92eb804e";

let nodeAddress = "0x0bd81846A232E263bC289A0B2206Ae4ae26A08A2";

let provider = new ethers.providers.JsonRpcProvider(url);
let signer = new ethers.Wallet(privateKey, provider);

let format = (x: any) => console.log(ethers.utils.formatUnits(x, 0));

export {
  address,
  ethers, 
  provider, 
  signer, 
  format,
  nodeAddress 
};

