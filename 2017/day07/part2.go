package main

import (
        "fmt"
        "os"
        "bufio"
        "strings"
)

const NULL = ""

type Node struct {
        weight int
        totalWeight int
        subs []string
}

func main() {
        f, _ := os.Open(os.Args[1])
        scanner := bufio.NewScanner(f)

        tower := make(map[string]Node)

        for scanner.Scan() {
                line := scanner.Text()
                words := strings.Fields(line)
                elem := words[0]
                var weight int
                fmt.Sscanf(words[1], "(%d)", &weight)

                if len(words) > 2 {
                        subs := extract(words[3:])
                        tower[elem] = Node{weight, weight, subs}
                } else {
                        tower[elem] = Node{weight, weight, nil}
                }
        }

        var start string
        for k := range tower {
                start = k
                break
        }

        node := findParent(tower, start)
        var root string
        for node != NULL {
                root = node
                node = findParent(tower, node)
        }

        fillWeights(tower, root)
        odd := findOdd(tower, root)
        diff := findDifferent(tower, tower[root].subs)
        var surplus int
        if tower[root].subs[0] == diff {
                surplus = tower[tower[root].subs[0]].totalWeight-tower[tower[root].subs[1]].totalWeight
        } else {
                surplus = tower[diff].totalWeight-tower[tower[root].subs[0]].totalWeight
        }
        fmt.Println(tower[odd].weight-surplus)
}

func findDifferent(tower map[string]Node, children []string) string {
        if len(children) > 2 {
                a, b, c := tower[children[0]], tower[children[1]], tower[children[2]]
                if a.totalWeight == b.totalWeight {
                        for _, v := range children[2:] {
                                if tower[v].totalWeight != a.totalWeight {
                                        return v
                                }
                        }
                        return NULL
                } else if a.totalWeight == c.totalWeight {
                        if b.totalWeight == a.totalWeight {
                                return NULL
                        } else {
                                return children[1]
                        }
                } else if b.totalWeight == c.totalWeight {
                        if b.totalWeight == a.totalWeight {
                                return NULL
                        } else {
                                return children[0]
                        }
                } else {
                        return NULL
                }
        }
        return NULL
}

func findOdd(tower map[string]Node, node string) string {
        if len(tower[node].subs) > 0 {
                diff := findDifferent(tower, tower[node].subs)
                if diff != NULL {
                        return findOdd(tower, diff)
                } else {
                        return node
                }
        }
        return node
}

func extract(words []string) []string {
        var subs []string
        for _, v := range words[:len(words)-1] {
                subs = append(subs, v[:len(v)-1])
        }
        subs = append(subs, words[len(words)-1])
        return subs
}

func findParent(tower map[string]Node, node string) string {
        for k, v := range tower {
                for _, el := range v.subs {
                        if el == node {
                                return k
                        }
                }
        }
        return NULL
}

func fillWeights(tower map[string]Node, node string) {
        children := tower[node].subs
        weight := tower[node].weight
        for _, child := range children {
                fillWeights(tower, child)
        }
        for _, child := range children {
                weight += tower[child].totalWeight
        }
        currentNode := tower[node]
        currentNode.totalWeight = weight
        tower[node] = currentNode
}
