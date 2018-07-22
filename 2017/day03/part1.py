#!/usr/bin/python

num = input()

i = 1
while i*i < num:
	i += 2

if num < i*i - 3*(i-1):
	print abs(i*i - 3*(i-1)-i/2-num) + i/2
elif num < i*i - 2*(i-1):
	print abs(i*i - 2*(i-1)-i/2-num) + i/2
elif num < i*i - (i-1):
	print abs(i*i - (i-1)-i/2-num) + i/2
else:
	print abs(i*i-i/2-num) + i/2

