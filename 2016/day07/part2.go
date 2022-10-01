package main

import (
        "fmt"
        "strings"
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
        nets, outs := extract_hypernets(line)
        for _, net := range nets {
                for _, aba := range extract_abas(net) {
                        for _, out := range outs {
                                if strings.Contains(out, aba[1:] + string(aba[1])) {
                                        return true
                                }
                        }
                }
        }
        return false
}

func extract_abas(line string) []string {
        var abas []string
        for i := 0; i <= len(line) - 3; i++ {
                if line[i] != line[i+1] && line[i] == line[i+2] {
                        abas = append(abas, line[i:i+3])
                }
        }
        return abas
}

func extract_hypernets(line string) ([]string, []string) {
        var nets []string
        var outs []string
        i := 0
        var start int
        for i < len(line) {
                if line[i] == '[' {
                        outs = append(outs, line[start:i])
                        start = i
                } else if line[i] == ']' {
                        nets = append(nets, line[start+1:i])
                        start = i+1
                }
                i++
        }
        outs = append(outs, line[start:])
        return nets, outs
}
