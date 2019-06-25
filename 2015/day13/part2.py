#/usr/bin/python

import sys

def convert(line):
    words = line[:-1].split()
    number = int(words[3])
    info = (words[0],
            number if words[2] == "gain" else -number,
            words[-1])
    return info

def find_n(l):
    n = 2
    while n*(n-1) != l:
        n += 1
    return n

def give_ids(rels):
    id_dict = {}
    id_dict[rels[0][0]] = 0
    i = 1
    for _, _, p2 in rels:
        id_dict[p2] = i
        i += 1
    return id_dict

def to_table(rels):
    n = find_n(len(rels))
    table = [[0] * (n+1) for _ in range(n+1)]
    people = {}
    id_dict = give_ids(rels[:n-1])
    for p1, v, p2 in rels:
        p1_i = id_dict[p1]
        p2_i = id_dict[p2]
        table[p1_i][p2_i] = v
    return table

bond_energy = lambda t, e1, e2 : t[e1][e2] + t[e2][e1]

def solve(table):
    from itertools import permutations

    max_h = 0
    for perm in permutations(range(len(table))):
        bond = bond_energy(table, perm[0], perm[-1])
        for i in range(len(table)-1):
            bond += bond_energy(table, perm[i], perm[i+1])
        if bond > max_h:
            max_h = bond
    return max_h

with open(sys.argv[1]) as f:
    lines = f.read().splitlines()

relations = [convert(line) for line in lines]

table = to_table(relations)
result = solve(table)
print(result)
