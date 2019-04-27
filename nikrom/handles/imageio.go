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
    if err.code != error.NoError {
        return (nil, err)
    }

    var check bool

    _, check = args[0].(string)
    err = error.AssertArgumentType(check, 1, "string", reflect.TypeOf(args[0]).Name())
    if err.code != error.NoError {
        return (nil, err)
    }

    return (lib.DecodeImage(reflect.ValueOf(args[0])), error.NoError())
}

func SaveHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(3, len(args))
    if err.code != error.NoError {
        return (nil, err)
    }

    var check bool
    pos := 0

    _, check = args[pos].(image.Image)
    err = error.AssertArgumentType(check, pos + 1, "string", reflect.TypeOf(args[pos]).Name())
    if err.code != error.NoError {
        return (nil, err)
    }
    pos += 1

    _, check = args[pos].(string)
    err = error.AssertArgumentType(check, pos + 1, "string", reflect.TypeOf(args[pos]).Name())
    if err.code != error.NoError {
        return (nil, err)
    }

    _, check = args[pos].(string)
    err = error.AssertArgumentType(check, pos + 1, "string", reflect.TypeOf(args[pos]).Name())
    if err.code != error.NoError {
        return (nil, err)
    }
    pos += 1

    return (lib.EncodeImage(reflect.ValueOf(args[0])), error.NoError())
}
