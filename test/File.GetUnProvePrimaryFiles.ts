import { expect } from "chai";
import { addrs, file } from "./initialize";

var path = require('path');
var scriptName = path.basename(__filename);

describe(scriptName, function () {

  it(scriptName, async () => {
    const res = file.GetUnProvePrimaryFiles(addrs[15]);
    expect(res).to.not.be.reverted;
  });
});
