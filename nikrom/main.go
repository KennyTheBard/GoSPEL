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
	str := "save (rotate (resize (load image.jpg) (rect (point 0 0) (point 1000 1000))) 90) result png"
	fmt.Println(str)

	tree := parser.BuildTree(str)

	// printTree(">", tree)

	_, err := interpreter.Interpret(tree)

	fmt.Println(err)
}
