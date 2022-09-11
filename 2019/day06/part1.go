package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	reader := bufio.NewScanner(f)

	orbits := make(map[string][]string)
	counts := make(map[string]int)

	for reader.Scan() {
		line := reader.Text()
		planets := strings.Split(line, ")")
		node1, node2 := planets[0], planets[1]
		orbits[node1] = append(orbits[node1], node2)
	}
	dfs(orbits, counts, "COM", 0)
	sum := 0
	for _, v := range counts {
		sum += v
	}
	fmt.Println(sum)
}

func dfs(orbits map[string][]string, counts map[string]int, node string, current int) {
	adjList := orbits[node]
	counts[node] = current

	for _, n := range adjList {
		dfs(orbits, counts, n, current+1)
	}
}
