#!/usr/bin/python

def incr(s):
	if not s:
		return ''

	for i in xrange(len(s)):
		if s[i] in 'oli':
			s = s[:i] + chr(ord(s[i])+1) + ('a'*(len(s)-i-1))
			return s

	if s[-1] == 'z':
		return incr(s[:-1]) + 'a'
	return s[:-1] + chr(ord(s[-1])+1)

def constr1(s):
	if len(s) < 3:
		return False
	num1, num2, num3 = [ord(i) for i in s[:3]]
	num2 -= num1
	num3 -= num1
	return (num2,num3) == (1,2) or constr1(s[1:])

def constr3(s):
	cnt = 0
	i = 0
	while i < len(s)-1:
		if s[i] == s[i+1]:
			cnt += 1
			i += 2
		else:
			i += 1
	return cnt > 1

string = 'hepxcrrq'
while not (constr1(string) and constr3(string)):
	string = incr(string)

print string
