package error

import (
    generics "../generics"
)

func Error(code int, message string) (Error) {
    return Error{code, message}
}

func NoError() (Error) {
    return Error{NoError, "No unusual behaviour"}
}

func NumberArgumentsError(expected, received int) (Error) {
    return Error{error.InvalidNumberOfArguments, "Expected " + expected + \
        ", received " + received + "!"}
}

func ArgumentTypeError(pos int, expected, received string) {
    return Error{error.InvalidArgumentType, "Expected argument " + pos + \
        " of type " + epected + ", received " + received + "!"}
}
