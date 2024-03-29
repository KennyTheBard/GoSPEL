package control

import (
    "reflect"
    generics "../generics"
    error "../error"
    macro "../macro"
)

/**
 *  Handle the arguments and define a macro with
 *  the given name and InterpreterTree.
 *  Usage: define <name> (<code>)
 */
func DefineHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the macro name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    name, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the macro
    args[pos] = raw_args[pos].(generics.InterpreterTree)

    // create a new macro entry
    (&macro.Macros).AddMacro(name, args[pos])

    // return
    return nil, error.CreateNoError()
}
