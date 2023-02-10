#!/usr/bin/env python3

import shubhlipi as sh, os

ENV = ["development", "production"]
ENV_NAME = [".env", f".env.{ENV[0]}", f".env.{ENV[1]}"]
if sh.args(0) == "login":
    sh.cmd("npx dotenv-vault login", False)
elif sh.args(0) == "push":
    for x in sh.argv[1:]:
        sh.cmd(f"npx dotenv-vault push {ENV[int(x)-1]}", False)
elif sh.args(0) == "pull":
    for x in sh.argv[1:]:
        sh.cmd(f"npx dotenv-vault pull {ENV[int(x)-1]}", False)
        if x == "1" and os.path.isfile(ENV_NAME[0]):
            if os.path.isfile(ENV_NAME[1]):
                os.remove(ENV_NAME[1])
            os.rename(ENV_NAME[0], ENV_NAME[1])
