#!/usr/bin/python
import json

l = []

def recparse(lst):
	global l
	if type(lst) == list:
		for elem in lst:
			recparse(elem)
	elif type(lst) == dict:
		if 'red' in lst.values():
			return
		else:
			for vals in lst.values():
				recparse(vals)
	else:
		try:
			l.append(int(lst))
		except:
			pass
			
inp = raw_input()
lst = json.loads(inp)
recparse(lst)
print sum(l)
