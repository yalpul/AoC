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
	countMap := make([]int, len(points))
	for i := xmin; i <= xmax; i++ {
		for j := ymin; j <= ymax; j++ {
			closest, equal := findClosest(points, i, j)
			if !equal {
				countMap[closest]++
			}
		}
	}
	eliminated := eliminate(points, xmin, xmax, ymin, ymax)

	max := countMap[0]
	for i, v := range countMap {
		if v > max && !eliminated[i] {
			max = v
		}
	}
	fmt.Println(max)
}

func findClosest(points [][3]int, x, y int) (int, bool) {
	closest, distance := 0, 1 << 30
	var equal bool
	for id, p := range points {
		dist := findDist(x, y, p[0], p[1])
		if dist < distance {
			distance = dist
			closest = id
			equal = false
		} else if dist == distance {
			equal = true
		}
	}
	return closest, equal
}

func eliminate(points [][3]int, xmin, xmax, ymin, ymax int) map[int]bool {
	eliminated := make(map[int]bool)
	for i := xmin; i <= xmax; i++ {
		closest, _ := findClosest(points, i, ymin)
		eliminated[closest] = true
		closest, _ = findClosest(points, i, ymax)
		eliminated[closest] = true
	}
	for j := ymin; j <= ymax; j++ {
		closest, _ := findClosest(points, xmin, j)
		eliminated[closest] = true
		closest, _ = findClosest(points, xmax, j)
		eliminated[closest] = true
	}
	return eliminated
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

