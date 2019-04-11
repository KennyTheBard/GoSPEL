package main

import (
    "fmt"
    "reflect"
)

type Void interface {}

type String struct {
    s string
}

type Pair struct {
    *String
    k int
}

func print(s Void) {
    fmt.Println(s)
}

func (s *String) print() {
    fmt.Println(s.s)
}

func (p *Pair) print() {
    fmt.Println(p.k, ":", p.s)
}

func call(v Void) {
    call(reflect.ValueOf(v))
}

func main() {
    var v Void
    print(v)
    v = "hello world"
    print(v)
    print(1)
    s := String{"this is a struct"}
    print(s)
    k := "this is a simple string"
    print(k)

    s.print()
    p := Pair{&String{"value"}, 10}
    p.print()
}
