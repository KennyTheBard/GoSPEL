package control

import (
    "fmt"
    "reflect"
    generics "../generics"
    error "../error"
)

func PrintHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgumentControl(3, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var arg0 generics.Namespace
    var arg1 string
    var arg2 generics.InterpreterTree
    var aux generics.Void
    var ok bool
    pos := 0

    arg0 = args[pos].(generics.Namespace)
    pos += 1

    aux, err = args[pos].(generics.InterpreterTree).Interpret(arg0)
    arg1, ok = aux.(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string",
        reflect.TypeOf(aux).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    arg2 = args[pos].(generics.InterpreterTree)

    var ret generics.Void
    ret, err = arg2.Interpret(arg0)

    fmt.Println(arg1)

    return ret, err
}
