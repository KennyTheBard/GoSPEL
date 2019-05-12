package operation

import (
    "image"
    "reflect"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and call the required sub-handle.
 *  Usage: rect <sub_handle> ...
 */
func RectangleHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
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
        return NewRectangleHandle(scope, raw_args[pos:])
    case "first":
        return FirstHandle(scope, raw_args[pos:])
    case "last":
        return LastHandle(scope, raw_args[pos:])
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle name for rectangle \"" + sub_handle + "\"!")
    }
}

/**
 *  Handle the arguments and create a new Rectangle.
 *  Usage: rect new <point_A> <point_B>
 */
func NewRectangleHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the min point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    min, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }
    pos += 1

    // extract the max point
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err = args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    max, ok := aux.(image.Point)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Point", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return image.Rectangle{min, max}, error.CreateNoError()
}

/**
 *  Handle the arguments and returns the first point of the rectangle.
 *  Usage: rect new first <rectangle>
 */
func FirstHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
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
    rect, ok := aux.(image.Rectangle)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Rectangle", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return rect.Min, error.CreateNoError()
}

/**
 *  Handle the arguments and returns the last point of the rectangle.
 *  Usage: rect new last <rectangle>
 */
func LastHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
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
    rect, ok := aux.(image.Rectangle)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "image.Rectangle", reflect.TypeOf(aux).Name())
    }

    // call the operation
    return rect.Max, error.CreateNoError()
}
