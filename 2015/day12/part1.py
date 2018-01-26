#!/usr/bin/python

import re

inp = raw_input()
lst = re.findall('-?\d+', inp)
print sum([int(i) for i in lst])
