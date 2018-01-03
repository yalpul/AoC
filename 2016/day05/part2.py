#!/usr/bin/python
from hashlib import *
import sys
s = "reyedfim"
i = 0
k = 0
l = list('_'*8)

while k < 8:
    st = md5(s+str(i))
    digest = st.hexdigest()
    if digest[:5] == '00000':
        [pos, char] = digest[5:7]
        if '0' <= pos <= '7' and l[int(pos)] == '_':
            l[int(pos)] = char
            k += 1
            print '\r%s' % ("".join(l)),
            sys.stdout.flush()
    i += 1
