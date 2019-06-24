package operation

import (
    "strconv"
    "reflect"
    lib "../../lib"
    modifiers "../../lib/generators/modifiers"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and return the required Grayscale modifier.
 *  Usage: gen modif grayscale <red_ratio> <green_ratio> <blue_ratio>
 */
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
        return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
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
        return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
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
        return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
    }
    blue, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "float", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return modifiers.Grayscale(red, green, blue), error.CreateNoError()
}

/**
 *  Handle the arguments and return the required custom modifier.
 *  Usage: gen modif custom <4*5_float_elements>
 *      the elements represent groups of 4 ratios for channels
  *     + 1 ratio for max channel value; one group for each channel
 */
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

    // extract all the members of the modifier
    for pos := 0; pos < size * size + const_num; pos++ {
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

        idx := pos + 1
        if idx % 5 == 0 {
            constants[pos / 5] = member
        } else {
            mat[pos / 5][pos % 5] = member
        }
    }

    // return the modifier
    return lib.Modifier{mat, constants}, error.CreateNoError()
}
