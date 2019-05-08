package operation

import (
    "strconv"
    "reflect"
    lib "../../lib"
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

func CustomModifierHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error
    const mat_size = 4 * 4
    const constants = 4

    err = error.AssertNumberArgument(mat_size + constants, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool

    var mat [4][4]float64
    var consts [4]float64

    for i := 0; i < mat_size + constants; i++ {
        _, ok = args[i].(string)
        err = error.AssertArgumentType(!ok, i + 1, "string",
            reflect.TypeOf(args[i]).Name())
        if err.Code != error.NoError {
            return nil, err
        }

        aux, _ := args[i].(string)
        arg, _ := strconv.ParseFloat(aux, 64)

        idx := i + 1
        if idx % 5 == 0 {
            consts[i / 5] = arg
        } else {
            mat[i / 5][i % 5] = arg
        }
    }

    return lib.Modifier{mat, consts}, error.CreateNoError()
}
