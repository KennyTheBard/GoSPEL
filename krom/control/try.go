package control

import (
    "reflect"
    generics "../generics"
    error "../error"
)

func TryHandle(args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 3
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected - 1, received - 1)
    }

    // extract function scope
    scope := raw_args[0].(generics.Namespace)

    // prepare extraction for function arguments
    func_args := raw_args[1:]
    args := make([]generics.InterpreterTree, len(func_args))
    pos := 0

    // execute the risky branch
    args[pos] = func_args[pos].(generics.InterpreterTree)
    ret, err := args[pos].Interpret(scope.Clone())
    if err.Code == error.NoError {
        return ret, err
    }
    pos += 1

    // execute the fall-back branch
    args[pos] = func_args[pos].(generics.InterpreterTree)
    return args[pos].Interpret(scope.Clone())
}
