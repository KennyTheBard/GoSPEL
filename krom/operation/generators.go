package operation

import (
    "reflect"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and call the required sub-handle.
 *  Usage: gen <sub_handle> ...
 */
func GeneratorHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 2
    received := len(raw_args)
    if expected >= received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    // prepare extraction of function arguments
    pos := 0

    // extract the sub-handle
    aux, err := raw_args[0].(generics.InterpreterTree).Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    sub_handle, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // obtain the correct handle
    var handler generics.Handle
    switch sub_handle {
    case "filter":
        handler, err = FilterHandle(scope, []generics.Void{raw_args[1]})
    case "modif":
        handler, err = ModifierHandle(scope, []generics.Void{raw_args[1]})
    case "transf":
        handler, err = TransformationHandle(scope, []generics.Void{raw_args[1]})
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown sub-handle name for generator \"" + sub_handle + "\"!")
    }

    // pass the rest fo the arguments to the handle
    if err.Code != error.NoError {
        return nil, err
    }
    return handler(scope, raw_args[2:])
}

/**
 *  Handle the arguments and return the required filter handle.
 *  Usage: gen filter...
 */
func FilterHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Handle, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the filter's name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    name, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // return the filter's handle
    switch name {
    case "blur":
        return BoxBlurHandle, error.CreateNoError()
    case "custom":
        return CustomFilterHandle, error.CreateNoError()
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown filter name \"" + name + "\"!")
    }
}

/**
 *  Handle the arguments and return the required modifier handle.
 *  Usage: gen modif...
 */
func ModifierHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Handle, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the modifier's name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    name, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // return the modifier's handle
    switch name {
    case "grayscale":
        return GrayscaleHandle, error.CreateNoError()
    case "custom":
        return CustomModifierHandle, error.CreateNoError()
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown modifier name \"" + name + "\"!")
    }
}

/**
 *  Handle the arguments and return the required transformation handle.
 *  Usage: gen transf...
 */
func TransformationHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Handle, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected != received {
        return nil, error.NumberArgumentsError(expected, received)
    }

    // prepare extraction of function arguments
    args := make([]generics.InterpreterTree, len(raw_args))
    pos := 0

    // extract the transformation's name
    args[pos] = raw_args[pos].(generics.InterpreterTree)
    aux, err := args[pos].Interpret(scope.Clone())
    if err.Code != error.NoError {
        return nil, err
    }
    name, ok := aux.(string)
    if !ok {
        return nil, error.ArgumentTypeError(pos, "string", reflect.TypeOf(aux).Name())
    }

    // return the transformation's handle
    switch name {
    // case "mirror":
    //     // TODO
    // case "custom":
    //     // TODO
    default:
        return nil, error.CreateError(error.UnknownHandle,
            "Unknown transformation name \"" + name + "\"!")
    }
}
