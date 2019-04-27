package main

import (
	"fmt"
	"./parser"
	"./generics"
	interpreter "./interpreter"
)

func printTree(space string, tree generics.Atom) {
	fmt.Println(space, tree.Process)
	space = "~" + space
	for _, branch := range tree.Subatoms {
		printTree(space, branch)
	}
}

func main() {

	str := "save (copy (load image.jpg)) result png"

	fmt.Println(str)
	tree := parser.BuildTree(str)
	printTree(">", tree)

	_, err := interpreter.Interpret(tree)
	fmt.Println(err)

	// for _, t := range tokens {
	// 	fmt.Print(t)
	// 	fmt.Print(" -> ")
	// 	fmt.Println()
	//
	// 	aux := Tokenize(t)
	// 	for _, a := range aux {
	// 		fmt.Println("\t", a)
	// 	}
	// 	fmt.Println()
	// }
}
