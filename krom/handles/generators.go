package handles

import (
    "reflect"
    generics "../generics"
    error "../error"
)

func GeneratorHandle(args []generics.Void) (generics.Void, error.Error) {
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

    var handler generics.Handle
    switch arg0 {
    case "filter":
        handler, err = FilterHandle([]generics.Void{args[1]})
    case "modif":
        handler, err = ModifierHandle([]generics.Void{args[1]})
    case "transf":
        handler, err = TransformationHandle([]generics.Void{args[1]})
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle namefor generators \"" + arg0 + "\"!")
    }

    if err.Code != error.NoError {
        return nil, err
    }
    return handler(args[2:])
}

func FilterHandle(args []generics.Void) (generics.Handle, error.Error) {
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

    switch arg0 {
    case "blur":
        return BoxBlurHandle, error.CreateNoError()
    case "custom":
        return CustomFilterHandle, error.CreateNoError()
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown filter name \"" + arg0 + "\"!")
    }
}

func ModifierHandle(args []generics.Void) (generics.Handle, error.Error) {
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

    switch arg0 {
    case "grayscale":
        return GrayscaleHandle, error.CreateNoError()
    case "custom":
        return CustomModifierHandle, error.CreateNoError()
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown modifier name \"" + arg0 + "\"!")
    }
}

func TransformationHandle(args []generics.Void) (generics.Handle, error.Error) {
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

    switch arg0 {
    // case "mirror":
    //     // TODO
    // case "custom":
    //     // TODO
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown transformation name \"" + arg0 + "\"!")
    }
}
