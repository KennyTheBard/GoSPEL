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

    err = err = error.AssertNumberArgument(3, len(args))
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
    // TODO finish!!

    arg0, _ := args[0].(string)
}
