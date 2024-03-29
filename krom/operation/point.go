package operation

import (
    "math"
    "image"
    "reflect"
    "strconv"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and call the required sub-handle.
 *  Usage: point <sub_handle> ...
 */
func PointHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected >= received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    // prepare extraction of function arguments
    pos := 0

    // extract the sub-handle
    aux, err := raw_args[pos].(generics.InterpreterTree).Interpret(scope.Clone())
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
        return NewPointHandle(scope, raw_args[pos:])
    case "add":
        return AddPointHandle(scope, raw_args[pos:])
    case "mul":
        return MultiplyPointHandle(scope, raw_args[pos:])
    case "div":
        return DividePointHandle(scope, raw_args[pos:])
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle name for point \"" + sub_handle + "\"!")
    }
}

/**
 *  Handle the arguments and create a new Point.
 *  Usage: point new <int_x> <int_y>
 */
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
        return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
    }
    x, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the y coordinate
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    str, ok = aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "number as a string", reflect.TypeOf(aux).Name())
    }
    y, conv_err := strconv.ParseFloat(str, 64)
    if conv_err != nil {
        return nil, error.ArgumentTypeError(pos, "integer", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return image.Point{int(math.Round(x)), int(math.Round(y))}, error.CreateNoError()
}

/**
 *  Handle the arguments and return the sum of the two points.
 *  Usage: point add <point_A> <point_B>
 */
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
    aux, err = args[pos].Interpret(scope.Clone())
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

/**
 *  Handle the arguments and return the product of the two points.
 *  Usage: point mul <point_A> <point_B>
 */
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
    aux, err = args[pos].Interpret(scope.Clone())
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

/**
 *  Handle the arguments and return the quotient of the two points.
 *  Usage: point div <point_A> <point_B>
 */
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
    aux, err = args[pos].Interpret(scope.Clone())
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
