package handles

import (
    "image"
    "reflect"
    generics "../generics"
    error "../error"
)

func SizeofHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(1, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(image.Image)
    err = error.AssertArgumentType(!ok, pos + 1, "image.Image",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(image.Image)
    return arg0.Bounds(), error.CreateNoError()
}
