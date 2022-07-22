import { expect } from "chai";
import { addrs, list } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("add", async () => {
    const tx = list.WhiteListOperate({
      FileHash: [1, 2, 3],
      Op: 0,
      List: [
        {
          Addr: addrs[38],
          BaseHeight: 1,
          ExpireHeight: 100,
        }
      ]
    });
    await expect(tx).not.to.be.reverted;
  });

  it("del", async () => {
    const tx = list.WhiteListOperate({
      FileHash: [1, 2, 3],
      Op: 1,
      List: [
        {
          Addr: addrs[38],
          BaseHeight: 1,
          ExpireHeight: 100,
        }
      ]
    });
    await expect(tx).not.to.be.reverted;
  });

  it("add_cov", async () => {
    const tx = list.WhiteListOperate({
      FileHash: [1, 2, 3],
      Op: 2,
      List: [
        {
          Addr: addrs[38],
          BaseHeight: 1,
          ExpireHeight: 100,
        }
      ]
    });
    await expect(tx).not.to.be.reverted;
  });

  it("del_all", async () => {
    const tx = list.WhiteListOperate({
      FileHash: [1, 2, 3],
      Op: 3,
      List: [
        {
          Addr: addrs[38],
          BaseHeight: 1,
          ExpireHeight: 100,
        }
      ]
    });
    await expect(tx).not.to.be.reverted;
  });
});
