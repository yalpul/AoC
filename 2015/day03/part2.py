#!/usr/bin/python

world = {}

xs, ys = 0, 0
xr, yr = 0, 0
instr = raw_input()
santai = instr[::2]
roboi = instr[1::2]

world["0,0"] = True
for i in santai:
    def deliver():
        world[str(xs)+","+str(ys)] = True

    if i == "^":
        ys += 1
        deliver()
    elif i == "v":
        ys -= 1
        deliver()
    elif i == "<":
        xs -= 1
        deliver()
    elif i == ">":
        xs += 1
        deliver()

for i in roboi:
    def deliver():
        world[str(xr)+","+str(yr)] = True

    if i == "^":
        yr += 1
        deliver()
    elif i == "v":
        yr -= 1
        deliver()
    elif i == "<":
        xr -= 1
        deliver()
    elif i == ">":
        xr += 1
        deliver()

print len(world)
