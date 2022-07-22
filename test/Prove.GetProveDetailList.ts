import { assert } from "chai";
import { addrs, prove } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("get", async () => {
    const tx = prove.GetProveDetailList(addrs[17]);
    let res = await tx;
    // console.log(res)
    assert(res.length == 0)
  });

});
