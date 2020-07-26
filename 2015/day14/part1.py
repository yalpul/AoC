import sys

seconds = int(sys.argv[2])
with open(sys.argv[1]) as f:
    lines = f.read().splitlines()

values = []
for line in lines:
    words = line.split()
    values.append([int(i) for i in (words[3], words[6], words[13])])

def distance(deer):
    period = deer[1]+deer[2]
    run = deer[0]*deer[1]
    num_run = seconds // period
    remainder = min(seconds % period, deer[1])

    length = (num_run * deer[1] + remainder) * deer[0]
    return length

distances = [distance(deer) for deer in values]
print(max(distances))
