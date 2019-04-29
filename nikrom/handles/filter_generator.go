package handles

import (
    "strconv"
    "reflect"
    filters "../../lib/generators/filters"
    generics "../generics"
    error "../error"
)

func BoxBlurHandle(args []generics.Void) (generics.Void, error.Error) {
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

    aux0, _ := args[0].(string)
    arg0, _ := strconv.Atoi(aux0)
    return filters.BoxBlur(arg0), error.CreateNoError()
}
