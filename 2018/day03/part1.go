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
	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var left, top, width, height int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", new(int), &left, &top, &width, &height)
		for i := top; i < top+height; i++ {
			for j := left; j < left+width; j++ {
				space[i][j]++
			}
		}
	}
	for _, i := range space {
		for _, j := range i {
			if j > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
