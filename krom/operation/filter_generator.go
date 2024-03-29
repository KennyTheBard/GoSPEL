package operation

import (
    "math"
    "strconv"
    "reflect"
    lib "../../lib"
    filters "../../lib/generators/filters"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and return the required BoxBlur filter.
 *  Usage: gen filter blur <box_diameter>
 */
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

    // extract the diameter
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
    }
    diameter, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return filters.BoxBlur(int(math.Round(diameter))), error.CreateNoError()
}

/**
 *  Handle the arguments and return the required custom filter.
 *  Usage: gen filter custom <size> <size*size_float_elements>
 */
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
        return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
    }
    aux_size, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    size := int(math.Round(aux_size))

    // check the number of arguments depending on the size
    expected = size * size + 1
    if expected != received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    // create a matrix
    mat := make([][]float64, 1)
    for i := range mat {
        mat[i] = make([]float64, size)
    }

    // extract all the members of the filter
    for i := 1; i < size * size + 1; i++ {
        args[pos] = raw_args[pos].(generics.InterpreterTree)
        aux, err := args[pos].Interpret(scope.Clone())
        if err.Code != error.NoError {
            return nil, err
        }
        str, ok := aux.(string)
        if !ok {
            return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
        }
        member, conv_err := strconv.ParseFloat(str, 64)
        if conv_err != nil {
            return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
        }

        idx := i - 1
        mat[idx / size][idx % size] = member
    }

    // create the filter
    return lib.Filter{mat}, error.CreateNoError()
}
