#!/usr/bin/python

constr1 = lambda x : len(x) > 3 and (x[:2] in x[2:] or constr1(x[1:]))
constr2 = lambda x : len(x) > 2 and (x[0] == x[2] or constr2(x[1:]))

cnt = 0
while True:
    try:
        inp = raw_input()
    except:
        break
    if constr1(inp) and constr2(inp):
        cnt += 1
print cnt
