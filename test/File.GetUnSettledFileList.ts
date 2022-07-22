import { assert } from "chai";
import { addrs, file } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res2 = file.GetUnSettledFileList(addrs[7]);
    let tx = await res2;
    assert(tx.length == 0)
    // console.log(tx)
  });
});
