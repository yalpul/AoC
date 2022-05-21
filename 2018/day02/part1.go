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
	var twos, threes int

	for scanner.Scan() {
		counts := make(map[rune]int)
		line := scanner.Text()
		var two, three bool
		for _, char := range line {
			counts[char]++
		}
		for _, v := range counts {
			if v == 2 {
				two = true
			} else if v == 3 {
				three = true
			}
		}
		if two {
			twos++
		}
		if three {
			threes++
		}
	}
	fmt.Println(twos*threes)
}
