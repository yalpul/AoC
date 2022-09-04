package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	reader := bufio.NewScanner(f)

	reader.Scan()
	line := reader.Text()

	fields := strings.Split(line, ",")

	codes := make([]int, 0, len(fields))
	for _, f := range fields {
		i, _ := strconv.Atoi(f)
		codes = append(codes, i)
	}

loop:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result := check(codes, i, j)
			if result == 19690720 {
				fmt.Println(100*i + j)
				break loop
			}
		}
	}
}

func check(input []int, a, b int) int {
	codes := make([]int, len(input))
	for i, c := range input {
		codes[i] = c
	}
	codes[1] = a
	codes[2] = b
	current := codes[0]
	i := 0

	for current != 99 {
		switch current {
		case 1:
			codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
			i += 4
			current = codes[i]
		case 2:
			codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
			i += 4
			current = codes[i]
		}
	}

	return codes[0]
}
