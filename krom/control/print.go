package control

import (
    "fmt"
    "reflect"
    generics "../generics"
    error "../error"
)

func PrintHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the mesage
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    mesage, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // execute the body
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    ret, err := args[pos].Interpret(scope.Clone())

    // print
    fmt.Println(mesage)
    return ret, err
}
