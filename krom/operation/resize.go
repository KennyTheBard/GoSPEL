package operation

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func ResizeHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
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
    img, ok := aux.(image.Image)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Image", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the bounds
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    bounds, ok := aux.(image.Rectangle)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Rectangle", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return lib.Resize(img, bounds), error.CreateNoError()
}
