package handle

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func LoadHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(1, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool

    _, ok = args[0].(string)
    err = error.AssertArgumentType(!ok, 1, "string", reflect.TypeOf(args[0]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(string)
    return lib.DecodeImage(arg0), error.CreateNoError()
}

func SaveHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(3, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(image.Image)
    err = error.AssertArgumentType(!ok, pos + 1, "image.Image", reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    _, ok = args[pos].(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string", reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    _, ok = args[pos].(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string", reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(image.Image)
    arg1, _ := args[1].(string)
    arg2, _ := args[2].(string)
    return lib.EncodeImage(arg0, arg1, arg2), error.CreateNoError()
}
