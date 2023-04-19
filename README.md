# SaveFS Contract Solidity Version

## Usage

### Deploy contract

```
./hardhat.sh run deploy/fs.ts  --network eth
./hardhat.sh run deploy/dns.ts  --network eth
```

### Upgrade contract

```
./hardhat.sh run deploy/upgrade-fs.ts  --network eth
./hardhat.sh run deploy/upgrade-dns.ts  --network eth
```

### Build golang code

```
python build.py
```
Output in ./build/go

## Test Coverage

```
npx hardhat coverage
```

Report:

```
  171 passing (14s)

-----------------|----------|----------|----------|----------|----------------|
File             |  % Stmts | % Branch |  % Funcs |  % Lines |Uncovered Lines |
-----------------|----------|----------|----------|----------|----------------|
 contracts/      |    58.84 |    37.99 |       68 |     58.2 |                |
  Config.sol     |    95.83 |    83.33 |      100 |    95.83 |             47 |
  Dns.sol        |    29.76 |     12.5 |    17.31 |    28.76 |... 731,739,740 |
  File.sol       |    67.27 |    42.31 |    86.21 |       64 |... 448,497,498 |
  FileExtra.sol  |    69.08 |    43.75 |    83.78 |    67.48 |... 466,477,500 |
  List.sol       |      100 |    81.25 |    66.67 |      100 |                |
  Node.sol       |    81.11 |    67.65 |      100 |    81.32 |... 188,193,194 |
  PDP.sol        |    82.94 |    53.85 |    67.39 |    80.09 |... 672,680,707 |
  Prove.sol      |    46.01 |    23.58 |    76.47 |    44.98 |... 573,575,583 |
  ProveExtra.sol |    56.84 |    20.83 |       60 |    57.43 |... 280,288,315 |
  Sector.sol     |    87.41 |    65.22 |    95.45 |    87.84 |... 392,393,420 |
  Space.sol      |    50.25 |    38.27 |    84.38 |    50.12 |... 1,1052,1053 |
  Test.sol       |        0 |      100 |        0 |        0 |              6 |
  interface.sol  |      100 |      100 |      100 |      100 |                |
  type.sol       |      100 |      100 |      100 |      100 |                |
-----------------|----------|----------|----------|----------|----------------|
All files        |    58.84 |    37.99 |       68 |     58.2 |                |
-----------------|----------|----------|----------|----------|----------------|
```

## Contract size

```
npx hardhat size-contracts 
```

Report:

```
  ·----------------------|--------------|----------------·
 |  Contract Name       ·  Size (KiB)  ·  Change (KiB)  │
 ·······················|··············|·················
 |  DNSNodeMapping      ·       0.084  ·                │
 ·······················|··············|·················
 |  NameInfoMapping     ·       0.084  ·                │
 ·······················|··············|·················
 |  PeerPoolMapping     ·       0.084  ·                │
 ·······················|··············|·················
 |  ChallengeMapping    ·       0.084  ·                │
 ·······················|··············|·················
 |  MerkleNodeMapping   ·       0.084  ·                │
 ·······················|··············|·················
 |  MerklePathMapping   ·       0.084  ·                │
 ·······················|··············|·················
 |  ProofsMapping       ·       0.084  ·                │
 ·······················|··············|·················
 |  IterableMapping     ·       0.084  ·                │
 ·······················|··············|·················
 |  console             ·       0.084  ·                │
 ·······················|··············|·················
 |  AddressUpgradeable  ·       0.084  ·                │
 ·······················|··············|·················
 |  Test                ·       0.115  ·                │
 ·······················|··············|·················
 |  Config              ·       1.455  ·                │
 ·······················|··············|·················
 |  List                ·       3.339  ·                │
 ·······················|··············|·················
 |  Node                ·      11.247  ·                │
 ·······················|··············|·················
 |  ProveExtra          ·      12.116  ·        -0.099  │
 ·······················|··············|·················
 |  Sector              ·      15.663  ·                │
 ·······················|··············|·················
 |  PDP                 ·      17.477  ·        -0.208  │
 ·······················|··············|·················
 |  FileExtra           ·      19.390  ·                │
 ·······················|··············|·················
 |  Space               ·      20.306  ·                │
 ·······················|··············|·················
 |  Dns                 ·      20.786  ·                │
 ·······················|··············|·················
 |  File                ·      22.626  ·                │
 ·······················|··············|·················
 |  Prove               ·      23.176  ·                │
 ·----------------------|--------------|----------------·
```
