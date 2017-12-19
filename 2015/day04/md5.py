#/usr/bin/python

from hashlib import md5

# lim = 5
lim = 6
i = 1
while True:
        obj = md5("yzbqklnj"+str(i))
        if obj.hexdigest()[:lim] == "0"*lim:
                break
        i += 1
print i
