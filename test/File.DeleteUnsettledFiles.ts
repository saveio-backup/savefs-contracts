import { expect } from "chai";
import { addrs, file } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, function () {

  it("delete unsettled files", async () => {
    const tx = file.DeleteUnsettledFiles(addrs[10])
    await expect(tx).to.not.be.reverted;
  });

});
