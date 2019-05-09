package main

import (
    "os"
    "fmt"
    "strconv"
    "io/ioutil"
    krom "./krom"
)

func printTree(space string, tree krom.Atom) {
	fmt.Println(space, tree.Process)
	space = "~" + space
	for _, branch := range tree.Subatoms {
		printTree(space, branch.(krom.Atom))
	}
}

func main() {
    file, err_file := os.Open(os.Args[1])
    if err_file != nil {
        fmt.Println("Could not open script file", os.Args[1])
    }
    defer file.Close()

    bs, _ := ioutil.ReadAll(file)
    script := string(bs)

    tree := krom.BuildTree(script)
    // printTree(">", tree)

    args := os.Args[2:]
    scope := krom.NewScope()

    for i, arg := range args {
        scope = scope.Extend(strconv.Itoa(i), arg).(krom.Scope)
    }

    _, err := tree.Interpret(scope)
    fmt.Println(err)
}
