package main

import (
        "fmt"
        "os"
        "bufio"
)

func main() {
        f, _ := os.Open(os.Args[1])
        reader := bufio.NewScanner(f)

        count := 0
        for reader.Scan() {
                line := reader.Text()
                if valid(line) {
                        count++
                }
        }

        fmt.Println(count)
}

func valid(line string) bool {
        nets := extract_hypernets(line)
        for _, net := range nets {
                if has_abba(net) {
                        return false
                }
        }
        if has_abba(line) {
                return true
        }
        return false
}

func has_abba(line string) bool {
        for i := 0; i <= len(line) - 4; i++ {
                if line[i] != line[i+1] && line[i+1] == line[i+2] &&
                        line[i] == line[i+3] {
                        return true
                }
        }
        return false
}

func extract_hypernets(line string) []string {
        var nets []string
        i := 0
        var start int
        for i < len(line) {
                if line[i] == '[' {
                        start = i
                } else if line[i] == ']' {
                        nets = append(nets, line[start+1:i])
                }
                i++
        }
        return nets
}
