import { provider, address, format, ethers, nodeAddress,fileAddress, signer, print } from "./config";

(async function main() {
  const File = await ethers.getContractFactory("File");
  let file = new ethers.Contract(fileAddress, File.interface, signer);

  const tx = file.DeleteFiles([
    new TextEncoder().encode("Savezb2rhijMiMQYPnWZ6QicCwCMqPdhwLeQviu9SBLQXDMe95Aug")
  ]);
  print(tx)

})();

