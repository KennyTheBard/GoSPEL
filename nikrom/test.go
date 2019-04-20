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

func (s String) Print() {
    fmt.Println(s.s)
}

func (p Pair) Print() {
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

func (s String) FuncWithArgs(i int) {
	fmt.Println(i)
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

	s.Print()
	p := Pair{&String{"value"}, 10}
	p.Print()

	var arr []Void
	for i := 1; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)

	args := Arguments{arr}
	fmt.Println(args)
	callArguments(args)

	s = String{"this is a struct"}
	s.Print()
	// fmt.Println( reflect.TypeOf(String{}).Method(0) )
	fmt.Println( reflect.TypeOf(String{}).Method(1) )
	// fmt.Println( reflect.TypeOf(String{}).MethodByName("FuncWithArgs") )//.Type().In(0).Elem().Name()

	// nu are nevoie de argumente in call deoarece se foloseste ValueOf
	fmt.Println( reflect.ValueOf(s).Method(1).Call(nil))

	// are nevoie de valori ina rgumente deoarece se foloseste TypeOf
	i := 1
	inArgs := []reflect.Value{reflect.ValueOf(s), reflect.ValueOf(i)}
	fmt.Println( reflect.TypeOf(String{}).Method(0).Func.Call(inArgs))

	// ambele returneaza [] deoarece nu au return type, dar printeaza
	// din interiorul functiei

}
