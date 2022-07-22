import { assert } from "chai";
import { addrs, file } from "./initialize";

var path = require('path');
var get = path.basename(__filename);

describe(get, function () {

  it("get", async () => {
    const tx = file.GetFileList(addrs[46]);
    let res = await tx;
    // console.log(tx)
    assert(res.length == 0)
  });

});
