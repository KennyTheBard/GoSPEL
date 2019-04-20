package main

import (
	"fmt"
	"reflect"
)

type String struct {
    s string
}

func (f String) Bar() string {
	return f.s
}

func (f String) Print() {
    fmt.Println(f.s)
}

func (f String) Baz() {
}

func main() {
	fmt.Println(reflect.TypeOf(String{}).NumMethod())
}
