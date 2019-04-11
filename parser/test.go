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

type Arguments struct {
    args []Void
}

func callArguments(a Arguments) {
    sum := 0
    for arg := range a.args {
        if reflect.TypeOf(sum) == reflect.TypeOf(arg) {
            sum += arg
        }
    }
    fmt.Println(sum)
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

    var arr []Void
    for i := 1; i < 10; i++ {
        arr = append(arr, i)
    }
    fmt.Println(arr)

    args := Arguments{arr}
    fmt.Println(args)
    callArguments(args)
}
