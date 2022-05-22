package main

import (
	"fmt"
	"bufio"
	"bytes"
	"strings"
	"os"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	min := 1 << 30
	for c := 'a'; c <= 'z'; c++ {
		newString := strings.ReplaceAll(line, string(c), "")
		newString = strings.ReplaceAll(newString, strings.ToUpper(string(c)), "")
		length := react(newString)
		if length < min {
			min = length
		}
	}
	fmt.Println(min)
}

func react(line string) int {
	lineBytes := []byte(line)
	lineBytesLower := bytes.ToLower(lineBytes)
	for {
		for i := 0; i < len(lineBytes)-1; i++ {
			f, s := lineBytes[i], lineBytes[i+1]
			lowerF, lowerS := lineBytesLower[i], lineBytesLower[i+1]
			if f != s && lowerF == lowerS {
				lineBytes[i] = ' '
				lineBytes[i+1] = ' '
				lineBytesLower[i] = ' '
				lineBytesLower[i+1] = ' '
				i++
			}
		}
		lineBytes2 := bytes.ReplaceAll(lineBytes, []byte{' '}, nil)
		lineBytesLower = bytes.ReplaceAll(lineBytesLower, []byte{' '}, nil)
		if bytes.Equal(lineBytes, lineBytes2) {
			break
		}
		lineBytes = lineBytes2
	}
	return len(lineBytes)
}
