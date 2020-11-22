#!/usr/bin/python3

import sys

total = 101

with open(sys.argv[1]) as f:
    data = f.read().splitlines()
    values = []
    for line in data:
        _, _, cap, _, dur, _, fla, _, tex, _, _  = line.split()
        cap = int(cap[:-1])
        dur = int(dur[:-1])
        fla = int(fla[:-1])
        tex = int(tex[:-1])
        values.append([cap, dur, fla, tex])

best = 0
for i in range(total):
    for j in range(total-i):
        for k in range(total-i-j):
            l = total - 1 - i - j - k
            first = [num*i for num in values[0]]
            second = [num*j for num in values[1]]
            third = [num*k for num in values[2]]
            fourth = [num*l for num in values[3]]

            capacity = first[0] + second[0] + third[0] + fourth[0]
            durability = first[1] + second[1] + third[1] + fourth[1]
            flavor = first[2] + second[2] + third[2] + fourth[2]
            texture = first[3] + second[3] + third[3] + fourth[3]

            result = 1
            for value in (capacity, durability, flavor, texture):
                result *= max(0, value)
            if result > best:
                best = result

print(best)

