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

def rotate(tpl):
    s, num = tpl
    s = list(s)
    ns = ""
    for c in s:
        pos = ord(c)-ord('a')
        shf = (pos + num % 26)%26
        ns += chr(ord('a')+shf)
    return ns

with open('input.txt') as f:
    valids = []
    out = {}
    for line in map(lambda x:x.split('-'), f.readlines()):
        codes, sector_id, chksum = "".join(line[:-1]), int(line[-1][:3]), line[-1][4:-2]
        if max5(codes) == chksum:
            valids.append((codes, sector_id))
    for v in valids:
        out[rotate(v)] = v[1]
    print out['northpoleobjectstorage']

