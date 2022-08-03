import os
import json

list = [
    "Config",
    "File",
    "List",
    "Node",
    "PDP",
    "Prove",
    "Sector",
    "Space"
]

for i in list:
    source = "artifacts/contracts/{}.sol/{}.json".format(i, i)
    abi = "build/{}.abi".format(i)
    bin = "build/{}.bin".format(i)
    go = "build/go/{}/{}.go".format(i, i)
    os.makedirs(os.path.dirname(go), exist_ok=True)

    with open(source, "r") as f:

        if os.path.isfile(abi):
            os.remove(abi)
        if os.path.isfile(bin):
            os.remove(bin)
        if os.path.isfile(go):
            os.remove(go)

        data = json.load(f)

        a = data["abi"]
        w = open(abi, "w")
        w.write(json.dumps(a))
        w.close()

        b = data["bytecode"]
        w = open(bin, "w")
        w.write(b)
        w.close()

        c = "abigen --bin=build/{}.bin --abi=build/{}.abi --pkg=store --out=build/go/{}/{}.go".format(
            i, i, i, i)
        os.system(c)
