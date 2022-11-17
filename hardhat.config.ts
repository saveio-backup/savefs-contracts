import * as dotenv from "dotenv";

import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "solidity-coverage";
import '@openzeppelin/hardhat-upgrades';
import 'hardhat-contract-sizer';

dotenv.config();

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.4",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
        details: {
          yul: false
        }
      },
    },
  },
  networks: {
    ropsten: {
      url: process.env.ROPSTEN_URL || "",
      accounts:
        process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : [],
    },
    hardhat: {
      // allowUnlimitedContractSize: true,
      accounts: {
        count: 100
      }
    },
    ganache: {
      url: `http://127.0.0.1:7545/`,
      accounts: [`58d923527aac8f3e792f341e7e186f5aa9a8555ec7426bb639415710015460d2`],
    },
    op: {
      url: "http://localhost:8545",
      accounts:["0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"],
    },
    opeth: {
      url: "http://localhost:9545",
      accounts:["0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"],
    },
    eth: {
      url: "http://localhost:8545",
      accounts:["644fe9176025b45f853f27592b3f6038a61ee063018deff9c94ffa8552635d12"],
    },
    dev: {
      url: "http://117.50.17.39:38545",
      accounts:["b6e21cc1e3a9e2892bdeb82943eab2e0194a460c0878edca67584c3ac7f8ce66"],
    },
    
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};

export default config;
