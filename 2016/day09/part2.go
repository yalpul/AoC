package main

import (
        "fmt"
        "os"
        "io"
)

func main() {
        f, _ := os.Open(os.Args[1])

        data, _ := io.ReadAll(f)
        fmt.Println(decompress(data)-1)
}

func decompress(data []byte) int {
        c := 0
        totalLength := 0
        for c < len(data) {
                if data[c] == '(' {
                        d := c+1
                        for data[d] != ')' {
                                d++
                        }
                        var a, b int
                        fmt.Sscanf(string(data[c:d]), "(%dx%d", &a, &b)
                        c = d+1
                        length := decompress(data[c:c+a])
                        totalLength += length*b
                        c = c+a
                        continue
                }
                c++
                totalLength++
        }
        return totalLength
}
