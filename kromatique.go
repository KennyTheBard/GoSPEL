package main

import (
    "os"
    "fmt"
    "io/ioutil"
    krom "./krom"
)

func main() {
    filepath := os.Args[1]
    file, err_file := os.Open(filepath)
    if err_file != nil {
        fmt.Println("Could not open script file", filepath)
    }
    defer file.Close()

    bs, _ := ioutil.ReadAll(file)
    script := string(bs)
    (&krom.ImportedFiles).Import(filepath)
    forrest := krom.BuildForrest(script)

    args := os.Args[2:]
    err := krom.ExecuteAll(forrest, krom.ConvertStringArguments(args))
    fmt.Println(err)
}
