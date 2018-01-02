#!/usr/bin/python

def max5(code):
    def comp(x,y):
        if code.count(x) > code.count(y):
            return -1
        if code.count(x) == code.count(y) and x < y:
            return -1
        return 1

    l = list(set(code))
    l.sort(comp)
    return "".join(l[:5])

with open('input.txt') as f:
    sums = 0
    for line in map(lambda x:x.split('-'), f.readlines()):
        codes, sector_id, chksum = "".join(line[:-1]), int(line[-1][:3]), line[-1][4:-2]
        if max5(codes) == chksum:
            sums += sector_id
    print sums

