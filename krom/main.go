package main

import (
	"fmt"
	parser "./parser"
	generics "./generics"
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
	str := `save
				(modify
					(merge
						(crop
							(load image.jpg)
							(rect new
								(point new 0 0)
								(point new 300 300)))
						(filter
							(rotate
								(resize
									(load image.jpg)
									(rect new
										(point new 0 0)
										(point new 1000 1000)))
								90)
							(gen filter custom 7
								0.02 0.02 0.02 0.02 0.02 0.02 0.02
								0.02 0.02 0.02 0.02 0.02 0.02 0.02
								0.02 0.02 0.02 0.02 0.02 0.02 0.02
								0.02 0.02 0.02 0.02 0.02 0.02 0.02
								0.02 0.02 0.02 0.02 0.02 0.02 0.02
								0.02 0.02 0.02 0.02 0.02 0.02 0.02
								0.02 0.02 0.02 0.02 0.02 0.02 0.02 ))
						(point new 150 0))
					(gen modif grayscale 0.33 0.33 0.33))
				result
				png`
	fmt.Println(str)

	tree := parser.BuildTree(str)

	printTree(">", tree)

	_, err := interpreter.Interpret(tree)

	fmt.Println(err)
}
