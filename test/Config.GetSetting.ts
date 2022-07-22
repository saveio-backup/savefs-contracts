import { assert } from "chai";
import { config } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get setting", async () => {
    const tx = config.GetSetting();
    let res = await tx;
    assert(res.MinVolume.eq(1000000))
  });

});
