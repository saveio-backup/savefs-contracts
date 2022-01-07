# SaveFS Contract Solidity Version

## Coverage

You can got the coverage of the solidity code by running the following command:

```
npx hardhat coverage
```

There is the output:

```
Version
=======
> solidity-coverage: v0.7.17

Instrumenting for coverage...
=============================

> FileSystem.sol
> IFileSystem.sol
> Types.sol

Compilation:
============

Nothing to compile
No need to generate any newer typings.

Network Info
============
> HardhatEVM: v2.8.0
> network:    hardhat

No need to generate any newer typings.


  CalculateNodePledge.ts
    ✓ Deploy (92ms)
    ✓ CalculateNodePledge.ts

  GetFileInfo.ts
    ✓ Deploy (42ms)
    ✓ GetFileInfo.ts

  GetFileInfos.ts
    ✓ Deploy (60ms)
    ✓ GetFileInfos.ts

  GetFileList.ts
    ✓ Deploy (76ms)
    ✓ GetFileList.ts

  GetNodeInfoByNodeAddr.ts
    ✓ FDeploy (47ms)
    ✓ GetNodeInfoByNodeAddr.ts (83ms)

  GetNodeInfoByWalletAddr.ts
    ✓ Deploy (39ms)
    ✓ GetNodeInfoByWalletAddr.ts

  GetNodeList.ts
    ✓ Deploy
    ✓ GetNodeList.ts

  GetSettings.ts
    ✓ Deploy
    ✓ GetSettings.ts

  GetUploadStorageFee.ts
    ✓ Deploy
    ✓ GetUploadStorageFee.ts fileSize<0
    ✓ GetUploadStorageFee.ts fileSize>0 num==0
    ✓ GetUploadStorageFee.ts fileSize>0 num>0

  GetUserSpace.ts
    ✓ Deploy
    ✓ GetUserSpace.ts

  GetWhiteList.ts
    ✓ Deploy
    ✓ GetWhiteList.ts

  initialize.ts
    ✓ Deploy
initializer
    ✓ initialize.ts

  NodeCancel.ts
    ✓ Deploy
    ✓ NodeCancel.ts

  NodeRegister.ts
    ✓ Deploy
    ✓ NodeRegister.ts require 1
    ✓ NodeRegister.ts require 2
    ✓ NodeRegister.ts event

  NodeUpdate.ts
    ✓ Deploy (50ms)
    ✓ NodeUpdate.ts (55ms)

  NodeWithDrawProfit.ts
    ✓ Deploy
    ✓ NodeWithDrawProfit.ts


  36 passing (1s)

------------------|----------|----------|----------|----------|----------------|
File              |  % Stmts | % Branch |  % Funcs |  % Lines |Uncovered Lines |
------------------|----------|----------|----------|----------|----------------|
 contracts/       |    70.45 |    47.22 |    85.71 |    67.68 |                |
  FileSystem.sol  |    70.45 |    47.22 |    85.71 |    67.68 |... 366,389,402 |
  IFileSystem.sol |      100 |      100 |      100 |      100 |                |
  Types.sol       |      100 |      100 |      100 |      100 |                |
------------------|----------|----------|----------|----------|----------------|
All files         |    70.45 |    47.22 |    85.71 |    67.68 |                |
------------------|----------|----------|----------|----------|----------------|

> Istanbul reports written to ./coverage/ and ./coverage.json
```

## How to use

Try running some of the following tasks:

```shell
npx hardhat accounts
npx hardhat compile
npx hardhat clean
npx hardhat test
npx hardhat node
npx hardhat help
REPORT_GAS=true npx hardhat test
npx hardhat coverage
npx hardhat run scripts/deploy.ts
TS_NODE_FILES=true npx ts-node scripts/deploy.ts
npx eslint '**/*.{js,ts}'
npx eslint '**/*.{js,ts}' --fix
npx prettier '**/*.{json,sol,md}' --check
npx prettier '**/*.{json,sol,md}' --write
npx solhint 'contracts/**/*.sol'
npx solhint 'contracts/**/*.sol' --fix
```
