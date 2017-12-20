#!/usr/bin/python

constr1 = lambda x : sum(map(x.count, "aeiou")) >= 3

def constr2(s):
    prev = None
    for i in s:
        if i == prev:
            return True
        prev = i
    return False

constr3 = lambda x : all(map(lambda y : y not in x, ["ab","cd","pq","xy"]))

cnt = 0
while True:
   try:
    inp = raw_input()
   except:
    break
   if constr1(inp) and constr2(inp) and constr3(inp):
    cnt += 1

print cnt
