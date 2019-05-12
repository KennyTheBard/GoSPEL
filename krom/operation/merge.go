package operation

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and call the Merge function from the lib.
 *  Usage: merge <target_image> <overlay_image> <offset>
 */
func MergeHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 3
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the image
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    trg, ok := aux.(image.Image)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Image", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the image
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    over, ok := aux.(image.Image)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Image", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the bounds
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    offset, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return lib.Merge(trg, over, offset), error.CreateNoError()
}
