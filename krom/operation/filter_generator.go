package operation

import (
    "strconv"
    "reflect"
    lib "../../lib"
    filters "../../lib/generators/filters"
    generics "../generics"
    error "../error"
)

func BoxBlurHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the diameter coordinate
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux_x, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    diam, ok := aux_x.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux_x).Name())
    }
    x, ok := strconv.Atoi(aux_x)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux_x).Name())
    }

    // call the operation
    return filters.BoxBlur(diam), error.CreateNoError()
}

func CustomFilterHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected >= received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the size coordinate
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    size, ok := strconv.Atoi(str)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }

    // check the number of arguments depending on the size
    expected = size * size + 1
    if expected != received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    // create a matrix
    var mat [size][size]float64

    // extract all the members of the filter
    for i := 1; i < size * size + 1; i++ {
        args[pos] = raw_args[pos].(generics.InterpreterTree)
        aux, err := args[pos].Interpret(scope.Clone())
        if err.Code != error.NoError {
            return nil, err
        }
        str, ok := aux.(string)
        if !ok {
            return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
        }
        member, ok := strconv.ParseFloat(aux, 64)
        if !ok {
            return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
        }

        idx := i - 1
        mat[idx / size][idx % size] = member
    }

    // create the filter
    return lib.Filter{mat}, error.CreateNoError()
}
