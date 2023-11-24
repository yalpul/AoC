package main

import (
        "fmt"
        "os"
        "bufio"
        "strings"
)

type node struct {
        name string
        incoming []string
        outgoing []string
}

func main() {
        f, _ := os.Open(os.Args[1])
        scanner := bufio.NewScanner(f)
        nodes := make(map[string]*node)
        for scanner.Scan() {
                text := scanner.Text()
                var l1, l2 string
                fmt.Sscanf(text, "Step %s must be finished before step %s can begin.", &l1, &l2)
                node1 := nodes[l1]
                if node1 == nil {
                        node1 = &node{l1, nil, nil}
                        nodes[l1] = node1
                }
                node2 := nodes[l2]
                if node2 == nil {
                        node2 = &node{l2, nil, nil}
                        nodes[l2] = node2
                }
                node1.outgoing = append(node1.outgoing, l2)
                node2.incoming = append(node2.incoming, l1)
        }

        visitables := make(map[string]*node)
        visited := make(map[string]bool)
        for k, v := range nodes {
                if len(v.incoming) == 0 {
                        visitables[k] = v
                }
        }
        var list []string

        for len(visitables) > 0 {
                smallest := findSmallest(visitables)
                delete(visitables, smallest)
                visited[smallest] = true
                list = append(list, smallest)
                visit(smallest, visitables, visited, nodes)
        }
        fmt.Println(strings.Join(list, ""))
}

func visit(smallest string, visitables map[string]*node, visited map[string]bool, nodes map[string]*node) {
        for _, node := range nodes[smallest].outgoing {
                if !visited[node] {
                        allVisited := true
                        for _, ins := range nodes[node].incoming {
                                if !visited[ins] {
                                        allVisited = false
                                        break
                                }
                        }
                        if allVisited {
                                visitables[node] = nodes[node]
                        }
                }
        }
}

func findSmallest(visitables map[string]*node) string {
        var current string
        for k := range visitables {
                current = k
                break
        }
        for k := range visitables {
                if k < current {
                        current = k
                }
        }
        return current
}
