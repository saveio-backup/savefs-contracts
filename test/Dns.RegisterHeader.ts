import { assert, expect } from "chai";
import { dns, addrs } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("register header", async () => {
    const tx = dns.RegisterHeader({
      Header: new TextEncoder().encode("header"),
      NameOwner: addrs[80],
      Desc: new TextEncoder().encode("desc"),
      DesireTTL: 123,
    });
    expect(tx).to.not.be.reverted;
  });

});
