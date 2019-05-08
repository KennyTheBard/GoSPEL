package operation

import (
    "image"
    "reflect"
    "strconv"
    generics "../generics"
    error "../error"
)

func PointHandle(args []generics.Void) (generics.Void, error.Error) {
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
        return NewPointHandle(args[1:])
    case "add":
        return AddPointHandle(args[1:])
    case "multiply":
        return MultiplyPointHandle(args[1:])
    case "divide":
        return DividePointHandle(args[1:])
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle name for point \"" + arg0 + "\"!")
    }
}

func NewPointHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(2, len(args))
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
    pos += 1

    _, ok = args[pos].(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    aux0, _ := args[0].(string)
    arg0, _ := strconv.Atoi(aux0)
    aux1, _ := args[1].(string)
    arg1, _ := strconv.Atoi(aux1)
    return image.Point{arg0, arg1}, error.CreateNoError()
}

func AddPointHandle(args []generics.Void) (generics.Void, error.Error) {
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
    return image.Point{arg0.X + arg1.X, arg0.Y + arg1.Y}, error.CreateNoError()
}

func MultiplyPointHandle(args []generics.Void) (generics.Void, error.Error) {
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
    return image.Point{arg0.X * arg1.X, arg0.Y * arg1.Y}, error.CreateNoError()
}

func DividePointHandle(args []generics.Void) (generics.Void, error.Error) {
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
    return image.Point{arg0.X / arg1.X, arg0.Y / arg1.Y}, error.CreateNoError()
}
