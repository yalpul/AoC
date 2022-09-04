package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	f, _ := os.Open(os.Args[1])
	reader := bufio.NewScanner(f)
	reader.Scan()
	line1 := reader.Text()
	reader.Scan()
	line2 := reader.Text()

	grid := make(map[point]int)

	var intersect []int

	var x, y, i int
	for _, inst := range strings.Split(line1, ",") {
		dir := inst[0]
		num, _ := strconv.Atoi(inst[1:])
		switch dir {
		case 'U':
			lim := y + num
			for y < lim {
				y++
				i++
				if _, ok := grid[point{x, y}]; !ok {
					grid[point{x, y}] = i
				}
			}
		case 'D':
			lim := y - num
			for y > lim {
				y--
				i++
				if _, ok := grid[point{x, y}]; !ok {
					grid[point{x, y}] = i
				}
			}
		case 'L':
			lim := x - num
			for x > lim {
				x--
				i++
				if _, ok := grid[point{x, y}]; !ok {
					grid[point{x, y}] = i
				}
			}
		case 'R':
			lim := x + num
			for x < lim {
				x++
				i++
				if _, ok := grid[point{x, y}]; !ok {
					grid[point{x, y}] = i
				}
			}
		}
	}
	x, y, i = 0, 0, 0

	for _, inst := range strings.Split(line2, ",") {
		dir := inst[0]
		num, _ := strconv.Atoi(inst[1:])
		switch dir {
		case 'U':
			lim := y + num
			for y < lim {
				y++
				i++
				length, ok := grid[point{x, y}]
				if ok {
					intersect = append(intersect, length+i)
				}
			}
		case 'D':
			lim := y - num
			for y > lim {
				y--
				i++
				length, ok := grid[point{x, y}]
				if ok {
					intersect = append(intersect, length+i)
				}
			}
		case 'L':
			lim := x - num
			for x > lim {
				x--
				i++
				length, ok := grid[point{x, y}]
				if ok {
					intersect = append(intersect, length+i)
				}
			}
		case 'R':
			lim := x + num
			for x < lim {
				x++
				i++
				length, ok := grid[point{x, y}]
				if ok {
					intersect = append(intersect, length+i)
				}
			}
		}
	}

	min := intersect[0]
	for i := 1; i < len(intersect); i++ {
		if intersect[i] < min {
			min = intersect[i]
		}
	}

	fmt.Println(min)
}
