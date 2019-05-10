package operation

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func LoadHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the file's name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    filename, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // decode image
    img := lib.DecodeImage(filename)
    if img == nil {
        return nil, error.CreateError(error.FileError, "Could not open \"" + filename + "\"")
    }
    return img, error.CreateNoError()
}

func SaveHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
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
    img, ok := aux.(image.Image)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Image", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the file's name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    filename, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the file's format
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    format, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // encode image
    img = lib.EncodeImage(img, filename, format)
    if img == nil {
        return nil, error.CreateError(error.FileError, "Could not create \"" + filename + "\"")
    }
    return img, error.CreateNoError()
}
