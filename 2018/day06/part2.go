package main

import (
	"fmt"
	"os"
)


func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	var x, y int
	var points [][3]int
	n, _ := fmt.Fscanf(file, "%d, %d\n", &x, &y)
	id := 0
	for n > 0 {
		points = append(points, [3]int{x, y, id})
		n, _ = fmt.Fscanf(file, "%d, %d\n", &x, &y)
	}
	xmin, ymin, xmax, ymax := findBorder(points)
	count := 0
	for i := xmin; i <= xmax; i++ {
		for j := ymin; j <= ymax; j++ {
			totalDist := findTotalDist(points, i, j)
			if totalDist < 10000 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func findTotalDist(points [][3]int, x, y int) int {
	total := 0
	for _, p := range points {
		dist := findDist(x, y, p[0], p[1])
		total += dist
	}
	return total
}

func findBorder(points [][3]int) (int, int, int, int) {
	xmin := points[0][0]
	ymin := points[0][1]
	xmax := points[0][0]
	ymax := points[0][1]
	for _, p := range points[1:] {
		if p[0] < xmin {
			xmin = p[0]
		} else if p[0] > xmax {
			xmax = p[0]
		}
		if p[1] < ymin {
			ymin = p[1]
		} else if p[1] > ymax {
			ymax = p[1]
		}
	}
	return xmin, ymin, xmax, ymax
}

func findDist(x1, y1, x2, y2 int) int {
	x := x1-x2
	y := y1-y2
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

