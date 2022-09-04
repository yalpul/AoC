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

	grid := make(map[point]bool)

	var intersect []point

	var x, y int
	for _, inst := range strings.Split(line1, ",") {
		dir := inst[0]
		num, _ := strconv.Atoi(inst[1:])
		switch dir {
		case 'U':
			lim := y + num
			for y < lim {
				y++
				grid[point{x, y}] = true
			}
		case 'D':
			lim := y - num
			for y > lim {
				y--
				grid[point{x, y}] = true
			}
		case 'L':
			lim := x - num
			for x > lim {
				x--
				grid[point{x, y}] = true
			}
		case 'R':
			lim := x + num
			for x < lim {
				x++
				grid[point{x, y}] = true
			}
		}
	}
	x, y = 0, 0

	for _, inst := range strings.Split(line2, ",") {
		dir := inst[0]
		num, _ := strconv.Atoi(inst[1:])
		switch dir {
		case 'U':
			lim := y + num
			for y < lim {
				y++
				full, ok := grid[point{x, y}]
				if ok && full {
					intersect = append(intersect, point{x, y})
				} else {
					grid[point{x, y}] = false
				}
			}
		case 'D':
			lim := y - num
			for y > lim {
				y--
				full, ok := grid[point{x, y}]
				if ok && full {
					intersect = append(intersect, point{x, y})
				} else {
					grid[point{x, y}] = false
				}
			}
		case 'L':
			lim := x - num
			for x > lim {
				x--
				full, ok := grid[point{x, y}]
				if ok && full {
					intersect = append(intersect, point{x, y})
				} else {
					grid[point{x, y}] = false
				}
			}
		case 'R':
			lim := x + num
			for x < lim {
				x++
				full, ok := grid[point{x, y}]
				if ok && full {
					intersect = append(intersect, point{x, y})
				} else {
					grid[point{x, y}] = false
				}
			}
		}
	}

	first := intersect[0]
	min := abs(first.x) + abs(first.y)
	for i := 1; i < len(intersect); i++ {
		current := intersect[i]
		x, y := abs(current.x), abs(current.y)
		if x+y < min {
			min = x + y
		}
	}

	fmt.Println(min)
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}
