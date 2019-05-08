package operation

import (
    "image"
    "reflect"
    generics "../generics"
    error "../error"
)

func RectangleHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgumentAtLeast(1, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }


    arg0, _ := args[0].(string)

    switch arg0 {
    case "new":
        return NewRectangleHandle(args[1:])
    case "first":
        return FirstHandle(args[1:])
    case "last":
        return LastHandle(args[1:])
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle name for rectangle \"" + arg0 + "\"!")
    }
}

func NewRectangleHandle(args []generics.Void) (generics.Void, error.Error) {
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
