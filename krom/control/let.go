package control

import (
    "reflect"
    generics "../generics"
    error "../error"
)

func LetHandle(raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 4
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

    // extract the variable name
    args[pos] = func_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    name, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the variable value
    args[pos] = func_args[pos].(generics.InterpreterTree)
    value, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    // continue interpretation
    args[pos] = func_args[pos].(generics.InterpreterTree)
    return args[pos].Interpret(scope.Extend(name, value))
}
