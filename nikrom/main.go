package main

import (
	"fmt"
	"./parser"
	"./generics"
)

func printTree(space string, tree generics.Atom) {
	fmt.Println(space, tree.Process)
	space = "~" + space
	for _, branch := range tree.Subatoms {
		printTree(space, branch)
	}
}

func main() {

	str := "filter (copy (load image.jpg)) (gen_f grayscale ' :)')"

	fmt.Println(str)
	tree := parser.BuildTree(str)
	printTree(">", tree)

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
