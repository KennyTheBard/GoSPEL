package handles

import (
    "image"
    "reflect"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func GeneratorHandle(args []genercis.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgumentAtLeast(2, len(args))
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

    var handler handle
    switch arg0 {
    case "filter":
        handler, _ = FilterHandle([]generics.Void{args[1]})
    case "modif":
        handler, _ = ModifierHandle([]generics.Void{args[1]})
    case "transf":
        handler, _ = TransformationHandle([]generics.Void{args[1]})
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle name \"" + arg0 + "\"!")
    }

    return handler([]generics.Void{args[2:]})
}

func FilterHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(1, len(args))
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

    var handler handle
    switch arg0 {
    case "blur":
        // TODO
    case "custom":
        // TODO
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown filter name \"" + arg0 + "\"!")
    }

    return handle, error.CreateNoError()
}

func ModifierHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(1, len(args))
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

    var handler handle
    switch arg0 {
    case "grayscale":
        // TODO
    case "custom":
        // TODO
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown modifier name \"" + arg0 + "\"!")
    }

    return handle, error.CreateNoError()
}

func TransformationHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(1, len(args))
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

    var handler handle
    switch arg0 {
    case "mirror":
        // TODO
    case "custom":
        // TODO
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown transformation name \"" + arg0 + "\"!")
    }

    return handle, error.CreateNoError()
}
