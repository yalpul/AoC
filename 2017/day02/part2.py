#!/usr/bin/python

import sys

def find_diff(l):
    sl = sorted(l)
    for i in range(len(sl)):
        pivot = sl[-i-1]
        for j in range(len(sl)-i-1):
            if pivot % sl[j] == 0:
                return pivot / sl[j]

with open(sys.argv[1]) as f:
    content = f.read().splitlines()

content = [[int(i) for i in row.split()] for row in content]
vals = [find_diff(row) for row in content]
print(sum(vals))
