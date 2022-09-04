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

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}
