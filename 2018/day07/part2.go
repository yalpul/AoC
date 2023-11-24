package main

import (
        "fmt"
        "os"
        "bufio"
)

type node struct {
        name string
        incoming []string
        outgoing []string
}

type worker struct {
        idle bool
        node string
        remaining int
}


func main() {
        f, _ := os.Open(os.Args[1])
        scanner := bufio.NewScanner(f)
        nodes := make(map[string]*node)

        workers := make([]worker, 5)
        for i := range workers {
                workers[i].idle = true
        }

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

        for i := range workers {
                worker := &workers[i]
                smallest := findSmallest(visitables)
                if smallest == "" {
                        break
                }
                worker.idle = false
                worker.node = smallest
                worker.remaining = int(smallest[0] - 'A' + 1 + 60)
                delete(visitables, smallest)
                visited[smallest] = true
        }

        count := 0
        for anyWorkerNotIdle(workers) {
                count++
                for i := range workers {
                        worker := &workers[i]
                        if !worker.idle {
                                worker.remaining--
                                if worker.remaining == 0 {
                                        worker.idle = true
                                        visited[worker.node] = true
                                        visit(worker.node, visitables, visited, nodes)
                                }
                        }
                }
                for i := range workers {
                        worker := &workers[i]
                        smallest := findSmallest(visitables)
                        if smallest == "" {
                                break
                        }
                        if !worker.idle {
                                continue
                        }
                        worker.idle = false
                        worker.node = smallest
                        worker.remaining = int(smallest[0] - 'A' + 1 + 60)
                        delete(visitables, smallest)
                }
        }
        fmt.Println(count)
}

func anyWorkerNotIdle(workers []worker) bool {
        for _, v := range workers {
                if !v.idle {
                        return true
                }
        }
        return false
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
