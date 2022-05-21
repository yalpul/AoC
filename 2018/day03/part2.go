package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	var space [1000][1000]int
	overlap := make(map[int]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var id, left, top, width, height int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)
		for i := top; i < top+height; i++ {
			for j := left; j < left+width; j++ {
				if space[i][j] == 0 {
					if _, ok := overlap[id]; !ok {
						overlap[id] = false
					}
					space[i][j] = id
				} else if space[i][j] > 0 {
					overlap[space[i][j]] = true
					overlap[id] = true
					space[i][j] = id
				}
			}
		}
	}
	for k, v := range overlap {
		if !v {
			fmt.Println(k)
			os.Exit(0)
		}
	}
}
