package main

import (
        "fmt"
        "os"
        "io"
        "bytes"
)

func main() {
        f, _ := os.Open(os.Args[1])

        data, _ := io.ReadAll(f)
        buf := &bytes.Buffer{}

        c := 0

        for c < len(data) {
                if data[c] == '(' {
                        d := c+1
                        for data[d] != ')' {
                                d++
                        }
                        var a, b int
                        fmt.Sscanf(string(data[c:d]), "(%dx%d", &a, &b)
                        c = d+1
                        add := bytes.Repeat(data[c:c+a], b)
                        buf.Write(add)
                        c = c+a
                        continue
                }
                buf.WriteByte(data[c])
                c++
        }
        fmt.Println(buf.Len()-1)
}
