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

	var youOrbit, santaOrbit string
	for reader.Scan() {
		line := reader.Text()
		planets := strings.Split(line, ")")
		node1, node2 := planets[0], planets[1]
		if node2 == "YOU" {
			youOrbit = node1
		} else if node2 == "SAN" {
			santaOrbit = node1
		}
		orbits[node1] = append(orbits[node1], node2)
		orbits[node2] = append(orbits[node2], node1)
	}
	n := bfs(orbits, youOrbit, santaOrbit)
	fmt.Println(n)
}

func bfs(orbits map[string][]string, you, santa string) int {
	queue := []struct {
		node string
		dist int
	}{{you, 0}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node.node] = true
		current := node.dist + 1
		adjList := orbits[node.node]
		for _, n := range adjList {
			if visited[n] {
				continue
			}
			if n == santa {
				return current
			}
			queue = append(queue, struct {
				node string
				dist int
			}{n, current})
		}
	}
	return -1
}
