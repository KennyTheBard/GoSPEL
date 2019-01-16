package main

import (
    "fmt"
    "os"
    "strings"
    parser "./parser"
)

func main() {
    args := strings.Join(os.Args[1:], " ")
    p := parser.ParseTree{nil, nil}

    p.Parse(args)
    p.Resolve(parser.Scope{nil, nil})

    fmt.Println("Done!")
}
