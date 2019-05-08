package control

import (
    "fmt"
    "reflect"
    generics "../generics"
    error "../error"
)

func PrintHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(3, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var aux generics.InterpreterTree
    var arg0 generics.Namespace
    var msg generics.Void
    var ok bool
    pos := 0

    arg0, ok = args[pos].(generics.Namespace)
    err = error.AssertArgumentType(!ok, pos + 1, "generics.Namespace",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    aux, ok = args[pos].(generics.InterpreterTree)
    err = error.AssertArgumentType(!ok, pos + 1, "generics.InterpreterTree",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    msg, err = aux.Interpret(arg0)
    _, ok = msg.(string)
    err = error.AssertArgumentType(err.Code != error.NoError, pos + 1, "string",
        reflect.TypeOf(aux).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    _, ok = args[pos].(generics.InterpreterTree)
    err = error.AssertArgumentType(!ok, pos + 1, "generics.InterpreterTree",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg1, _ := msg.(string)
    arg2, _ := args[2].(generics.InterpreterTree)

    var ret generics.Void
    ret, err = arg2.Interpret(arg0)

    fmt.Println(arg1)

    return ret, err
}
