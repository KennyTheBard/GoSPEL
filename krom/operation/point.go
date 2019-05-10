package operation

import (
    "image"
    "reflect"
    "strconv"
    generics "../generics"
    error "../error"
)

func PointHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected >= received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the sub-handle
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    sub_handle, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // execute the sub-handle
    switch sub_handle {
    case "new":
        return NewPointHandle(scope, args[pos:])
    case "add":
        return AddPointHandle(scope, args[pos:])
    case "multiply":
        return MultiplyPointHandle(scope, args[pos:])
    case "divide":
        return DividePointHandle(scope, args[pos:])
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle name for point \"" + sub_handle + "\"!")
    }
}

func NewPointHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the x coordinate
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    x, ok := strconv.Atoi(str)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the y coordinate
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    y, ok := strconv.Atoi(str)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return image.Point{x, y}, error.CreateNoError()
}

func AddPointHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the A point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    a, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the B point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    b, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return image.Point{a.X + b.X, a.Y + b.Y}, error.CreateNoError()
}

func MultiplyPointHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the A point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    a, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the B point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    b, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return image.Point{a.X * b.X, a.Y * b.Y}, error.CreateNoError()
}

func DividePointHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the A point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    a, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the B point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    b, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return image.Point{a.X / b.X, a.Y / b.Y}, error.CreateNoError()
}
