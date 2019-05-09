package main

import (
    "os"
    "fmt"
    "io/ioutil"
    krom "./krom"
)

func main() {
    file, err_file := os.Open(os.Args[1])
    if err_file != nil {
        fmt.Println("Could not open script file", os.Args[1])
    }
    defer file.Close()

    bs, _ := ioutil.ReadAll(file)
    script := string(bs)

    forrest := krom.BuildForrest(script)

    args := os.Args[2:]
    err := krom.ExecuteAll(forrest, krom.ConvertStringArguments(args))
    fmt.Println(err)
}
