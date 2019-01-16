#!/usr/bin/python

import sys

with open(sys.argv[1]) as f:
    content = f.read().splitlines()

content = [[int(i) for i in row.split()] for row in content]
vals = [max(row)-min(row) for row in content]
print(sum(vals))
