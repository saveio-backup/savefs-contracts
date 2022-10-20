import assert from "assert";
import { addrs, prove } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("check time", async () => {
    const tx = prove.CheckNodeSectorProvedInTime({
      NodeAddr: addrs[24],
      SectorId: 1
    });
    let res = await (await tx).wait();
    if (res.events?.length == 1) {
      assert(res.events[0].event == "FsError");
    }
  });

});
