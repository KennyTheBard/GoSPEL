package operation

import (
    "image"
    "reflect"
    "strconv"
    lib "../../lib"
    generics "../generics"
    error "../error"
)

func RotateHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(2, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool
    pos := 0

    _, ok = args[pos].(image.Image)
    err = error.AssertArgumentType(!ok, pos + 1, "image.Image",
        reflect.TypeOf(args[0]).Name())
    if err.Code != error.NoError {
        return nil, err
    }
    pos += 1

    _, ok = args[pos].(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string",
        reflect.TypeOf(args[0]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    arg0, _ := args[0].(image.Image)
    aux1, _ := args[1].(string)
    arg1, _ := strconv.ParseFloat(aux1, 64)
    return lib.Rotate(arg0, arg1), error.CreateNoError()
}
