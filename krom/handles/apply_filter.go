package handles

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func ApplyFilterHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(2, len(args))
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
    pos += 1

    _, ok = args[pos].(lib.Filter)
    err = error.AssertArgumentType(!ok, pos + 1, "Filter",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(image.Image)
    arg1, _ := args[1].(lib.Filter)
    return lib.ApplyFilter(arg0, arg1), error.CreateNoError()
}
