import { expect } from "chai";
import { addrs, prove } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("check time", async () => {
    const res = prove.CheckNodeSectorProvedInTime({
      NodeAddr: addrs[24],
      SectorId: 1
    });
    expect(res).to.be.reverted;
  });

});
