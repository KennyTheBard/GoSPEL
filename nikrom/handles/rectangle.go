package handles

import (
    "image"
    "reflect"
    "strconv"
    generics "../generics"
    error "../error"
)

func RectangleHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(2, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(image.Point)
    err = error.AssertArgumentType(!ok, pos + 1, "image.Point",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    _, ok = args[pos].(image.Point)
    err = error.AssertArgumentType(!ok, pos + 1, "image.Point",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(image.Point)
    arg1, _ := args[1].(image.Point)
    return image.Rectangle{arg0, arg1}, error.CreateNoError()
}

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

func FirstHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(1, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(image.Rectangle)
    err = error.AssertArgumentType(!ok, pos + 1, "image.Rectangle",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(image.Rectangle)
    return arg0.Min, error.CreateNoError()
}

func LastHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(1, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(image.Rectangle)
    err = error.AssertArgumentType(!ok, pos + 1, "image.Rectangle",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(image.Rectangle)
    return arg0.Max, error.CreateNoError()
}
