#!/usr/bin/python

import sys

with open(sys.argv[1]) as f:
    content = f.read().strip()

l_sum = 0
for i in range(len(content)-1):
    if content[i] == content[i+1]:
        l_sum += int(content[i])
if content[-1] == content[0]:
    l_sum += int(content[0])
print(l_sum)
