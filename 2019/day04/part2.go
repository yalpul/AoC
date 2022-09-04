package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numRange := os.Args[1]
	numbers := strings.Split(numRange, "-")
	n1, n2 := numbers[0], numbers[1]
	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)
	count := 0
	for i := num1; i <= num2; i++ {
		if valid(i) {
			count++
		}
	}
	fmt.Println(count)
}

func valid(num int) bool {
	if num < 100000 || num > 999999 {
		return false
	}
	s := fmt.Sprintf("%d", num)

	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}

	numMap := make(map[rune]int)
	for _, b := range s {
		numMap[b]++
	}
	for _, v := range numMap {
		if v == 2 {
			return true
		}
	}
	return false
}
