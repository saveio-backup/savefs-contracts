import { expect } from "chai";
import { addrs, sector } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("delete sector", async () => {
    const tx = sector.DeleteSector({
      NodeAddr: addrs[70],
      SectorId: "100",
    });
    await expect(tx).to.not.be.reverted;
  });

});
