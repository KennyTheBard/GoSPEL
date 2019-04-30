package handles

import (
    "strconv"
    "reflect"
    modifiers "../../lib/generators/modifiers"
    generics "../generics"
    error "../error"
)

func GrayscaleHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgument(3, len(args))
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
    pos += 1

    _, ok = args[pos].(string)
    err = error.AssertArgumentType(!ok, pos + 1, "string",
        reflect.TypeOf(args[pos]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    aux0, _ := args[0].(string)
    arg0, _ := strconv.ParseFloat(aux0, 64)
    aux1, _ := args[1].(string)
    arg1, _ := strconv.ParseFloat(aux1, 64)
    aux2, _ := args[2].(string)
    arg2, _ := strconv.ParseFloat(aux2, 64)
    return modifiers.Grayscale(arg0, arg1, arg2), error.CreateNoError()
}
