package main

import (
        "fmt"
        "os"
        "bufio"
        "strings"
)

const NULL = ""

func main() {
        f, _ := os.Open(os.Args[1])
        scanner := bufio.NewScanner(f)

        tower := make(map[string][]string)

        for scanner.Scan() {
                line := scanner.Text()
                words := strings.Fields(line)
                elem := words[0]
                if len(words) > 2 {
                        subs := extract(words[3:])
                        tower[elem] = subs
                }
        }

        var start string
        for k := range tower {
                start = k
                break
        }
        node := findParent(tower, start)
        var root string
        for node != NULL {
                root = node
                node = findParent(tower, node)
        }

        fmt.Println(root)
}

func extract(words []string) []string {
        var subs []string
        for _, v := range words[:len(words)-1] {
                subs = append(subs, v[:len(v)-1])
        }
        subs = append(subs, words[len(words)-1])
        return subs
}

func findParent(tower map[string][]string, node string) string {
        for k, v := range tower {
                for _, el := range v {
                        if el == node {
                                return k
                        }
                }
        }
        return NULL
}
