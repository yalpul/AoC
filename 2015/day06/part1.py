#!/usr/bin/python

world = [[False for i in range(1000)] for j in range(1000)]

def parse(inp):
    line = inp.split()
    if line[0] == "turn":
        c1 = map(int, line[2].split(","))
        c2 = map(int, line[4].split(","))
        if line[1] == "on":
            return [0] + c1 + c2
        else:
            return [1] + c1 + c2
    else:
        c1 = map(int, line[1].split(","))
        c2 = map(int, line[3].split(","))
        return [2] + c1 + c2

while True:
    try:
        inp = raw_input()
    except:
        break
    cmd, x1, y1, x2, y2 = parse(inp)
    if cmd == 0:
        for i in xrange(x1,x2+1):
            for j in xrange(y1,y2+1):
                world[i][j] = True
    elif cmd == 1:
        for i in xrange(x1,x2+1):
            for j in xrange(y1,y2+1):
                world[i][j] = False
    elif cmd == 2:
        for i in xrange(x1,x2+1):
            for j in xrange(y1,y2+1):
                world[i][j] ^= True
cnt = 0
for i in xrange(1000):
    for j in xrange(1000):
        if world[i][j]:
            cnt += 1
print cnt
