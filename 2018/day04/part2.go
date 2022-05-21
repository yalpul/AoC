package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Strings(lines)

	var currentId, start int
	slept := make(map[int]int)
	sleptMinutes := make(map[int]*[60]int)
	for _, line := range lines {
		var id, minute int
		var op string
		fmt.Sscanf(line, "[%d-%d-%d %d:%d] %s #%d", new(int), new(int), new(int), new(int), &minute, &op, &id)
		switch op {
			case "Guard":
				currentId = id
				if _, ok := sleptMinutes[currentId]; !ok {
					sleptMinutes[currentId] = new([60]int)
				}
			case "falls":
				start = minute
			case "wakes":
				slept[currentId] += minute-start
				for i := start; i < minute; i++ {
					sleptMinutes[currentId][i]++
				}
		}
	}
	max, maxI, maxId := 0, 0, 0
	for id, minutes := range sleptMinutes {
		for m, freq := range minutes {
			if freq > max {
				max = freq
				maxI = m
				maxId = id
			}
		}
	}
	fmt.Println(maxI*maxId)
}
