#!/usr/bin/env python3

import shubhlipi as sh

for x in sh.argv:
    if x == "push":
        sh.cmd("cd py && python3 karah.py deploy", direct=False)
        sh.cmd("space push", direct=False)
