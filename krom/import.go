package krom

import (
    "os"
    "io/ioutil"
    "reflect"
    generics "./generics"
    error "./error"
)

func ImportHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the file name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    name, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // attempt to open the file
    file, err_open := os.Open(name + ".krom")
    if err_open != nil {
        return nil, error.CreateError(error.FileError,
            "The file " + name + "could not be imported!")
    }
    defer file.Close()

    // parse and interpret
    bs, _ := ioutil.ReadAll(file)
    script := string(bs)
    forrest := BuildForrest(script)
    return nil, ExecuteAll(forrest, []generics.Void{})
}
