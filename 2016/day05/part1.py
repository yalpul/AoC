#!/usr/bin/python
from hashlib import *

s = "reyedfim"
i = 0
l = []

while len(l) < 8:
    st = md5(s+str(i))
    digest = st.hexdigest()
    if digest[:5] == '00000':
        l.append(digest[5])
    i += 1
print "".join(l)
