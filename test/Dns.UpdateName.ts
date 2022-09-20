import { assert, expect } from "chai";
import { dns, addrs } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("register name", async () => {
    const tx = dns.RegisterName({
       Type: 0,
       Header: new TextEncoder().encode("header"),
       URL: new TextEncoder().encode("url"),
       Name: new TextEncoder().encode("name1"),
       NameOwner: addrs[80],
       Desc: new TextEncoder().encode("desc"),
       DesireTTL: 123,
    });
    expect(tx).to.not.be.reverted;
  });

  it("update name", async () => {
    const tx = dns.UpdateName({
       Type: 0,
       Header: new TextEncoder().encode("header"),
       URL: new TextEncoder().encode("url"),
       Name: new TextEncoder().encode("name1"),
       NameOwner: addrs[80],
       Desc: new TextEncoder().encode("desc"),
       DesireTTL: 123,
    });
    expect(tx).to.not.be.reverted;
  });

  it("update name failed", async () => {
    const tx = dns.UpdateName({
       Type: 0,
       Header: new TextEncoder().encode("header"),
       URL: new TextEncoder().encode("url"),
       Name: new TextEncoder().encode("name1"),
       NameOwner: addrs[81],
       Desc: new TextEncoder().encode("desc"),
       DesireTTL: 123,
    });
    expect(tx).to.be.reverted;
  });

});
