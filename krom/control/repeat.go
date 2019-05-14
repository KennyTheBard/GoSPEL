package control

import (
    "strconv"
    "reflect"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and repeat the given operation a
 *  predefined number of times, using a intern variable.
 *  The internal variable should have the same type as the
 *  return type of the repeated operation, as this variable
 *  will be fed back into the loop
 *  Usage: repeat <integer_repeats> <name> <value> (<operation>)
 */
func RepeatHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 4
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    pos := 0

    // extract the number of repetionions
    aux, err :=raw_args[pos].(generics.InterpreterTree).Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    repeat, conv_err := strconv.Atoi(str)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the variable name
    aux, err = raw_args[pos].(generics.InterpreterTree).Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    name, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the variable value
    value, err := raw_args[pos].(generics.InterpreterTree).Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    // extract the repeating code
    code := raw_args[pos].(generics.InterpreterTree)
    for i := 0; i < repeat; i++ {
        ret, err := code.Interpret(scope.Extend(name, value))
        if err.Code != error.NoError {
            return nil, err
        }
        // if reflect.TypeOf(ret) != reflect.TypeOf(value) {
        //     return nil, error.CreateError(error.InvalidArgumentType, "The repeating code produces elements of type " +
        //             reflect.TypeOf(ret).Name() + ", but it expects " + reflect.TypeOf(value).Name() + ". It needs the same type!")
        // }
        value = ret
    }

    return value, error.CreateNoError()
}
