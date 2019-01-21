#!/usr/bin/python

import sys

with open(sys.argv[1]) as f:
    content = f.read().splitlines()
    passes = [[''.join(sorted(y)) for y in x.split()] for x in content]


passes = filter(lambda x : len(list(set(x))) == len(x), passes)
print(len(passes))
