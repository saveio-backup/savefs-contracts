import { provider, address, format, ethers, nodeAddress,sectorAddress, signer } from "./config";

(async function main() {
  const Sector = await ethers.getContractFactory("Sector");
  let sector = new ethers.Contract(sectorAddress, Sector.interface, signer);

  const tx = sector.GetSectorsForNode("0x8203B0C002FEf4f85f9EDce29f23719f06f5FcAC");
  // console.log(tx)
  let res = await tx;
  console.log(res)

})();

