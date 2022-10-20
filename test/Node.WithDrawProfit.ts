import { expect } from "chai";
import { addrs, node } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("register", async () => {
    let tx = node.Register({
      Pledge: 1,
      Profit: 2,
      Volume: 1000 * 1000,
      RestVol: 3,
      ServiceTime: 4,
      WalletAddr: addrs[72],
      NodeAddr: addrs[72],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it("withdraw", async () => {
    const tx = node.WithDrawProfit(addrs[72]);
    expect(tx).to.be.reverted; // profit is 0
  });

  it("update", async () => {
    let tx = node.UpdateNodeInfo({
      Pledge: 1,
      Profit: 2,
      Volume: 1000 * 1000,
      RestVol: 3,
      ServiceTime: 4,
      WalletAddr: addrs[72],
      NodeAddr: addrs[72],
    }, {
      value: 1000000
    });
    expect(tx).to.not.be.reverted;
  });

  it("withdraw 2", async () => {
    const tx = node.WithDrawProfit(addrs[72]);
    expect(tx).to.not.be.reverted; // profit is not 0
  });

});
