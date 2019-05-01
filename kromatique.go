package main

import (
    "os"
    "fmt"
    "io/ioutil"
    parser "./krom/parser"
    interpreter "./krom/interpreter"
    generics "./krom/generics"
)

func printTree(space string, tree generics.Atom) {
	fmt.Println(space, tree.Process)
	space = "~" + space
	for _, branch := range tree.Subatoms {
		printTree(space, branch)
	}
}

func main() {
    args := os.Args[1:]

    for _, arg := range args {
        file, err_file := os.Open(arg)
        if err_file != nil {
            fmt.Println("Could not open file", arg)
            continue;
        }
        defer file.Close()

        bs, _ := ioutil.ReadAll(file)
        script := string(bs)

        tree := parser.BuildTree(script)
        // printTree(">", tree)
        _, err := interpreter.Interpret(tree)
    	fmt.Println(err)
    }
}