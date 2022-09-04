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

	codes[1] = 12
	codes[2] = 2
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

	fmt.Println(codes[0])
}
