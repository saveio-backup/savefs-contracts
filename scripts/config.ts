import { ContractTransaction } from "ethers";

const { ethers } = require("hardhat");

const url = "http://127.0.0.1:8545";
const address = "0x114c3ea85f6e7e623AEB5b66C639A2161B66E76F";
const privateKey =
  "644fe9176025b45f853f27592b3f6038a61ee063018deff9c94ffa8552635d12";

let nodeAddress = "0x0bd81846A232E263bC289A0B2206Ae4ae26A08A2";
let sectorAddress = "0x158bf3D044c8ECAc5e96165E280c2d46d7771091";
let fileAddress = "0x2BD421687AEaa170D7a12d1a776f3e441aaeAa11";

let provider = new ethers.providers.JsonRpcProvider(url);
let signer = new ethers.Wallet(privateKey, provider);

let format = (x: any) => console.log(ethers.utils.formatUnits(x, 0));
var print = async (tx: Promise<ContractTransaction>) => {
  let res = await (await tx).wait();
  console.log(res)
}

export {
  address,
  ethers, 
  provider, 
  signer, 
  format,
  print,
  nodeAddress,
  sectorAddress,
  fileAddress
};

