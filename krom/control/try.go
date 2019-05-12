package control

import (
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and try an operation, if it
 *  returns any error, it will execute the second operation
 *  returning its results.
 *  Usage: try (<risky_code>) (<fallback_code>)
 */
func TryHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // execute the risky branch
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    ret, err := args[pos].Interpret(scope.Clone())
    if err.Code == error.NoError {
        return ret, err
    }
    pos += 1

    // execute the fall-back branch
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    return args[pos].Interpret(scope.Clone())
}
