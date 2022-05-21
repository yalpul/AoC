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
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	outer:
	for i, line := range lines {
		for j := i+1; j < len(lines); j++ {
			line2 := lines[j]
			different := 0
			for k := 0; k < len(line); k++ {
				if line[k] != line2[k] {
					different++
				}
			}
			if different == 1 {
				for k := 0; k < len(line); k++ {
					if line[k] != line2[k] {
						fmt.Println(line[:k]+line[k+1:])
						break outer
					}
				}
			}
		}
	}
}
