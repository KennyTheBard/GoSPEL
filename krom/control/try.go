package control

import (
    "reflect"
    generics "../generics"
    error "../error"
)

func TryHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(3, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(generics.Namespace)
    err = error.AssertArgumentType(!ok, pos + 1, "generics.Namespace",
        reflect.TypeOf(args[pos]).Name())
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
    pos += 1

    _, ok = args[pos].(generics.InterpreterTree)
    err = error.AssertArgumentType(!ok, pos + 1, "generics.InterpreterTree",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(generics.Namespace)
    arg1, _ := args[1].(generics.InterpreterTree)

    var ret generics.Void
    ret, err = arg1.Interpret(arg0)
    if err.Code == error.NoError {
        return ret, err
    } else {
        arg2, _ := args[2].(generics.InterpreterTree)
        return arg2.Interpret(arg0)
    }
}
