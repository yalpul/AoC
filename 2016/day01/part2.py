#!/usr/bin/python

from sys import exit

line = raw_input()
line +=','
coords = map(lambda x : (x[0], int(x[1:-1])), line.split())

trail = set()
state = ['N', 0, 0]
def up(num):
    global trail
    global state
    x, y = state[1], state[2]
    state[2] = y + num
    for i in range(y, y+num):
        if (x,i) in trail:
            print abs(x) + abs(i)
            exit(0)
        else:
            trail.add((x,i))
def down(num):
    global trail
    global state
    x, y = state[1], state[2]
    state[2] = y - num
    for i in range(y-num+1, y+1):
        if (x,i) in trail:
            print abs(x) + abs(i)
            exit(0)
        else:
            trail.add((x,i))
def left(num):
    global trail
    global state
    x, y = state[1], state[2]
    state[1] = x - num
    for i in range(x-num+1, x+1):
        if (i,y) in trail:
            print abs(i) + abs(y)
            exit(0)
        else:
            trail.add((i,y))
def right(num):
    global trail
    global state
    x, y = state[1], state[2]
    state[1] = x + num
    for i in range(x, x+num):
        if (i,y) in trail:
            print abs(i) + abs(y)
            exit(0)
        else:
            trail.add((i,y))
for c in coords:
    if state[0] == 'N':
        if c[0] == 'R':
            state[0] = 'E'
            right(c[1])
        else:
            state[0] = 'W'
            left(c[1])
    elif state[0] == 'S':
        if c[0] == 'R':
            state[0] = 'W'
            left(c[1])
        else:
            state[0] = 'E'
            right(c[1])
    elif state[0] == 'W':
        if c[0] == 'R':
            state[0] = 'N'
            up(c[1])
        else:
            state[0] = 'S'
            down(c[1])
    elif state[0] == 'E':
        if c[0] == 'R':
            state[0] = 'S'
            down(c[1])
        else:
            state[0] = 'N'
            up(c[1])
