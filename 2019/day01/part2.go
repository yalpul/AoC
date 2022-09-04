package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	reader := bufio.NewScanner(f)
	sum := 0

	for reader.Scan() {
		t := reader.Text()
		mass, _ := strconv.Atoi(t)
		sum += fuelRequirement(mass)
	}
	fmt.Println(sum)
}

func fuelRequirement(mass int) int {
	fuel := mass/3 - 2
	current := fuel
	sum := current

	for {
		current = current/3 - 2
		if current > 0 {
			sum += current
		} else {
			break
		}
	}
	return sum
}
