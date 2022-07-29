import { expect } from "chai";
import { addrs, space } from "./initialize";

var path = require('path');
var name = path.basename(__filename);

describe(name, () => {

  it("get update cost", async () => {
    const tx = space.GetUpdateCost({
      WalletAddr: addrs[50],
      Owner: addrs[50],
      Size: {
        Type: 1,
        Value: 1024000
      },
      BlockCount: {
        Type: 1,
        Value: 172800
      }
    });
    expect(tx).not.to.be.reverted;

    let res = await tx;
    // console.log(res)

    let wait = await res.wait();
    if (wait.events){
      if (wait.events.length > 0) {
        let event = wait.events[0];
        if (event.args) {
          if (event.args.length > 0) {
            let arg = event.args[0];
            // console.log(arg)
            expect(arg.Value).not.to.be.eq(0);
          }
        }
      }
    }
  });
  
});
