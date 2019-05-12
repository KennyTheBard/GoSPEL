package control

import (
    "reflect"
    generics "../generics"
    error "../error"
    macro "../macro"
)

/**
 *  Handle the arguments and undefine a macro with the given name.
 *  Usage: undefine <name>
 */
func UndefineHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
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

    // remove the macro from the global defined list
    (&macro.Macros).RemoveMacro(name)

    // return
    return nil, error.CreateNoError()
}
