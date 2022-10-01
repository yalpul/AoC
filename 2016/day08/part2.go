package main

import (
        "fmt"
        "os"
        "strings"
        "strconv"
        "bufio"
)

const (
        width = 50
        height = 6
        rect = "rect"
        rotate = "rotate"
)

type grid struct {
        grid [height][width]bool
}

func main() {
        f, _ := os.Open(os.Args[1])
        reader := bufio.NewScanner(f)
        var field grid

        for reader.Scan() {
                line := reader.Text()
                words := strings.Split(line, " ")
                if cmd := words[0]; cmd == rect {
                        var a,b int
                        fmt.Sscanf(words[1], "%dx%d", &a, &b)
                        field.fill(a, b)
                } else if cmd == rotate {
                        amount, _ := strconv.Atoi(words[4])
                        var axis rune
                        var num int
                        fmt.Sscanf(words[2], "%c=%d", &axis, &num)

                        if axis == 'x' {
                                field.shiftDown(num, amount)
                        } else if axis == 'y' {
                                field.shiftRight(num, amount)
                        }
                }
        }
        field.print()
}

func (f *grid) numOns() int {
        count := 0
        for i := 0; i < height; i++ {
                for j := 0; j < width; j++ {
                        if f.grid[i][j] {
                                count++
                        }
                }
        }
        return count
}


func (f *grid) fill(a, b int) {
        for i := 0; i < b; i++ {
                for j := 0; j < a; j++ {
                        f.grid[i][j] = true
                }
        }
}

func (f *grid) shiftDown(num, amount int) {
        list := make([]bool, height)
        for i := 0; i < height; i++ {
                list[i] = f.grid[i][num]
        }
        list = append(list[height-amount:], list[:height-amount]...)
        for i := 0; i < height; i++ {
                f.grid[i][num] = list[i]
        }
}

func (f *grid) shiftRight(num, amount int) {
        list := make([]bool, width)
        for i := 0; i < width; i++ {
                list[i] = f.grid[num][i]
        }
        list = append(list[width-amount:], list[:width-amount]...)
        for i := 0; i < width; i++ {
                f.grid[num][i] = list[i]
        }
}

func (f *grid) print() {
        for i := 0; i < height; i++ {
                for j := 0; j < width; j++ {
                        if f.grid[i][j] {
                                fmt.Print("#")
                        } else {
                                fmt.Print(".")
                        }
                }
                fmt.Println("")
        }
}
