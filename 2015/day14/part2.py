import sys

seconds = int(sys.argv[2])
with open(sys.argv[1]) as f:
    lines = f.read().splitlines()

values = []
for line in lines:
    words = line.split()
    values.append([int(i) for i in (words[3], words[6], words[13])])

def distance(deer, seconds):
    period = deer[1]+deer[2]
    run = deer[0]*deer[1]
    num_run = seconds // period
    remainder = min(seconds % period, deer[1])

    length = (num_run * deer[1] + remainder) * deer[0]
    return length

points = [0] * len(values)
for second in range(1, seconds+1):
    dists = []
    for deer_i in range(len(values)):
        dists.append(distance(values[deer_i], second))
    zipped = list(reversed(sorted(zip(dists, range(len(values))))))
    max_val = zipped[0][0]
    points[zipped[0][1]] += 1
    for i in range(1, len(zipped)):
        if max_val == zipped[i][0]:
            points[zipped[i][1]] += 1
        else:
            break
        
print(max(points))
