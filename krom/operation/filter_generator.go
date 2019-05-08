package operation

import (
    "strconv"
    "reflect"
    lib "../../lib"
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

func CustomFilterHandle(args []generics.Void) (generics.Void, error.Error) {
    var err error.Error

    err = error.AssertNumberArgumentAtLeast(1, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    var ok bool

    _, ok = args[0].(string)
    err = error.AssertArgumentType(!ok, 1, "string",
        reflect.TypeOf(args[0]).Name())
    if err.Code != error.NoError {
        return nil, err
    }

    aux_size, _ := args[0].(string)
    size, _ := strconv.Atoi(aux_size)

    err = error.AssertNumberArgument(size * size + 1, len(args))
    if err.Code != error.NoError {
        return nil, err
    }

    mat := make([][]float64, size)
    for i := range mat {
        mat[i] = make([]float64, size)
    }

    for i := 1; i < size * size + 1; i++ {
        _, ok = args[i].(string)
        err = error.AssertArgumentType(!ok, i + 1, "string",
            reflect.TypeOf(args[i]).Name())
        if err.Code != error.NoError {
            return nil, err
        }

        aux, _ := args[i].(string)
        arg, _ := strconv.ParseFloat(aux, 64)

        idx := i - 1
        mat[idx / size][idx % size] = arg
    }

    return lib.Filter{mat}, error.CreateNoError()
}
