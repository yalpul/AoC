#!/usr/bin/python
import sys

knowledge_base = {}
with open(sys.argv[1]) as f:
    for line in f.readlines():
        expr, dest = line.split(' -> ')
        knowledge_base[dest[:-1]] = expr.split()

def evl(arg):
    if len(arg) == 1:
        return int(arg[0])
    if len(arg) == 2:
        return ~int(arg[1])
    
    op1, op, op2 = int(arg[0]), arg[1], int(arg[2])
    if op == "LSHIFT": return op1 << op2
    if op == "RSHIFT": return op1 >> op2
    if op == "AND"   : return op1 & op2
    if op == "OR"    : return op1 | op2

def isint(arg):
    try:
        int(arg)
        return True
    except:
        return False

def numeric(arg):
    if len(arg) == 1 and isint(arg[0])  or len(arg) == 2 and isint(arg[1]) \
                     or isint(arg[0]) and isint(arg[2]):
        return True
    return False
    
knowledge_base["b"] = ['3176']
while True:
    for key in knowledge_base.keys():
        val = knowledge_base[key]
        if numeric(val):
            val = evl(val)
            if key == "a":
                print key, val
                sys.exit()
            for key2 in knowledge_base.keys():
                knowledge_base[key2] = map(lambda x: str(val) if x == key else x, knowledge_base[key2])
            del knowledge_base[key]
