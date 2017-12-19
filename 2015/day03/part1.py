#!/usr/bin/python

world = {}

x, y = 0, 0
instr = raw_input()

world["0,0"] = True
for i in instr:
    def deliver():
        world[str(x)+","+str(y)] = True

    if i == "^":
        y += 1
        deliver()
    elif i == "v":
        y -= 1
        deliver()
    elif i == "<":
        x -= 1
        deliver()
    elif i == ">":
        x += 1
        deliver()
print len(world)
