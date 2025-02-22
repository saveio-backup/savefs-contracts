import { assert } from "chai";
import { config } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("high", async () => {
    const tx = config.GetSettingWithProveLevel(0);
    let res = await tx;
    // console.log(res)
    assert(res.DefaultProvePeriod.eq(17280))
  });

  it("medium", async () => {
    const tx = config.GetSettingWithProveLevel(1);
    let res = await tx;
    // console.log(res)
    assert(res.DefaultProvePeriod.eq(34560))
  });

  it("low", async () => {
    const tx = config.GetSettingWithProveLevel(2);
    let res = await tx;
    // console.log(res)
    assert(res.DefaultProvePeriod.eq(138240))
  });

});
