const { ethers } = require("hardhat");

const url = "http://127.0.0.1:8545";
const address = "0xd369b085e8907cBa039dD2BFA7b32E31a43273D3";
const privateKey =
  "07c1ba70f06ffb7752d0cf6edb150148ed5113d425ad5bec6d820ab12197d486";

let nodeAddress = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9";

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

