#!/usr/bin/python

import sys

with open(sys.argv[1]) as f:
    content = [int(i) for i in f.read().splitlines()]

s = set()
r = 0
l = len(content)
i = 0
p = content[i]
while r not in s:
    s.add(r)
    r += p
    i += 1
    p = content[i%l]
print(r)
