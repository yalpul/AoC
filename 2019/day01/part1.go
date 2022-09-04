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
		fuel := mass/3 - 2
		sum += fuel
	}
	fmt.Println(sum)
}
