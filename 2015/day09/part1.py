#!/usr/bin/python

from itertools import permutations
with open('input.txt') as f:
    roads = [line[:-1].split() for line in f.readlines()]
    cities = set([route[0] for route in roads] + [route[2] for route in roads])
    lcities = len(cities)
    dists = [[0 for _ in xrange(lcities)] for _ in xrange(lcities)]
    citynums = {}
    i = 0
    for city in cities:
        citynums[city] = i
        i += 1
    for route in roads:
        city1, city2, weight = citynums[route[0]], citynums[route[2]], int(route[4])
        dists[city1][city2] = weight
        dists[city2][city1] = weight
    mindist = 1 << 62 
    for perm in permutations(range(lcities)):
        current = perm[0]
        total = 0
        for nextcity in perm[1:]:
            total += dists[current][nextcity]
            current = nextcity
        mindist = total if total < mindist else mindist
    print mindist

