#!/usr/bin/python

import sys

with open(sys.argv[1]) as f:
    content = f.read().splitlines()
nums = [int(i) for i in content]
print(sum(nums))
