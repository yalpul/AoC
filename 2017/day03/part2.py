#!/usr/bin/python

l = [0,1,1,2,4,5,10,11,23]
i = 9

def coords(num):
	def nums(layer):
		i = layer * 2 + 1
		step, base = i - 1, i*i
		first = base - 3*step
		second = base - 2*step
		third = base - step
		return first,second,third,base
	i = 1
	while i*i < num:
		i += 2

	base, layer, step = i*i, i/2, i-1
	first, second, third, _ = nums(layer)
	subf, subs, subt, subb = nums(layer-1)

	if num == base:
		return num-1,subb,subb+1
	elif num == base -1:
		return num-1,subb,subb-1,subb+1
	elif num < first:
		pivot, pivot2, subp, subp2 = first-i+2, first, subf-i+4, subf
	elif num < second:
		pivot, pivot2, subp, subp2 = first, second, subf, subs
	elif num < third:
		pivot, pivot2, subp, subp2 = second, third, subs, subt
	else:
		pivot, pivot2, subp, subp2 = third, base, subt, subb

	if num == pivot:
		return num-1, subp
	elif num == pivot + 1:
		return num-2,num-1,subp,subp+1
	elif num == pivot2 - 1:
		return num-1,subp2,subp2-1
	else:
		gap = pivot2 - num
		lowp = subp2 - gap + 1
		return num-1,lowp-1,lowp,lowp+1

while True:
	num = sum([l[j] for j in coords(i)])
	if num > 325489:
		print num
		break
	else:
		l.append(num)
	i += 1
