package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HALT   = 99
	INPUT  = 3
	OUTPUT = 4
)

func main() {
	f, _ := os.Open(os.Args[1])
	reader := bufio.NewScanner(f)
	reader.Scan()
	line := reader.Text()
	tokens := strings.Split(line, ",")
	codes := make([]int, len(tokens))
	for i, v := range tokens {
		val, _ := strconv.Atoi(v)
		codes[i] = val
	}

	halt := false
	pc := 0

	for !halt {
		var newPc int
		newPc, halt = execute(codes, pc)
		pc = newPc
	}
}

func contains(l []int, p int) bool {
	for _, v := range l {
		if v == p {
			return true
		}
	}
	return false
}

func execute(codes []int, pc int) (int, bool) {
	switch code := codes[pc]; code {
	case HALT:
		return 0, true
	case INPUT:
		codes[codes[pc+1]] = 5
		return pc + 2, false
	case OUTPUT:
		fmt.Println(codes[codes[pc+1]])
		return pc + 2, false
	default:
		opcode := code % 100
		mode1 := (code / 100) % 10
		mode2 := (code / 1000) % 10
		var param1, param2 int
		if mode1 == 0 {
			param1 = codes[codes[pc+1]]
		} else {
			param1 = codes[pc+1]
		}
		if mode2 == 0 && contains([]int{1, 2, 7, 8}, opcode) {
			param2 = codes[codes[pc+2]]
		} else {
			param2 = codes[pc+2]
		}
		if opcode == 1 {
			codes[codes[pc+3]] = param1 + param2
		} else if opcode == 2 {
			codes[codes[pc+3]] = param1 * param2
		} else if opcode == 4 {
			fmt.Println(param1)
			return pc + 2, false
		} else if opcode == 5 {
			if param1 != 0 {
				if mode2 == 0 {
					return codes[codes[pc+2]], false
				} else {
					return codes[pc+2], false
				}
			}
			return pc + 3, false
		} else if opcode == 6 {
			if param1 == 0 {
				if mode2 == 0 {
					return codes[codes[pc+2]], false
				} else {
					return codes[pc+2], false
				}
			}
			return pc + 3, false
		} else if opcode == 7 {
			if param1 < param2 {
				codes[codes[pc+3]] = 1
			} else {
				codes[codes[pc+3]] = 0
			}
			return pc + 4, false
		} else if opcode == 8 {
			if param1 == param2 {
				codes[codes[pc+3]] = 1
			} else {
				codes[codes[pc+3]] = 0
			}
			return pc + 4, false
		} else {
			fmt.Println("unknown opcode:", opcode)
		}
		return pc + 4, false
	}
}
