package control

import (
    "reflect"
    generics "../generics"
    error "../error"
)

func LetHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgumentControl(4, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var arg0 generics.Namespace
    var arg1 generics.InterpreterTree
    var name string
    var arg3 generics.InterpreterTree
    var value generics.Void
    var aux generics.Void
    var ok bool
    pos := 0

    arg0 = args[pos].(generics.Namespace)
    pos += 1

    arg1 = args[pos].(generics.InterpreterTree)
    pos += 1

    aux, err = args[pos].(generics.InterpreterTree).Interpret(arg0)
    name, ok = aux.(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string",
        reflect.TypeOf(aux).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    arg3 = args[pos].(generics.InterpreterTree)

    value, err = arg1.Interpret(arg0)
    if err.Code != error.NoError {
        return nil, err
    }

    return arg3.Interpret(arg0.Extend(name, value))
}
