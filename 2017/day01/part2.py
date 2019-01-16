#!/usr/bin/python

import sys

with open(sys.argv[1]) as f:
    content = f.read().strip()

l_sum = 0
lc = len(content)
step = lc/2
for i in range(len(content)):
    if content[i] == content[(i+step)%lc]:
        l_sum += int(content[i])
print(l_sum)
