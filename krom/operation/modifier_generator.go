package operation

import (
    "strconv"
    "reflect"
    lib "../../lib"
    modifiers "../../lib/generators/modifiers"
    generics "../generics"
    error "../error"
)

func GrayscaleHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 3
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the red channel
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
    }
    red, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the green channel
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok = aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
    }
    green, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the blue channel
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok = aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
    }
    blue, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return modifiers.Grayscale(red, green, blue), error.CreateNoError()
}

func CustomModifierHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // create a matrix and an array
    const size = 4
    var mat [size][size]float64
    const const_num = 4
    var constants [const_num]float64

    // check the number of arguments
    expected := size * size + const_num
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract all the members of the modifier
    for i := 0; i < size * size + const_num; i++ {
        args[pos] = raw_args[pos].(generics.InterpreterTree)
        aux, err := args[pos].Interpret(scope.Clone())
        if err.Code != error.NoError {
            return nil, err
        }
        str, ok := aux.(string)
        if !ok {
            return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
        }
        member, conv_err := strconv.ParseFloat(str, 64)
        if conv_err != nil {
            return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
        }

        idx := i + 1
        if idx % 5 == 0 {
            constants[i / 5] = member
        } else {
            mat[i / 5][i % 5] = member
        }
    }

    // return the modifier
    return lib.Modifier{mat, constants}, error.CreateNoError()
}
