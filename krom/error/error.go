package error

import (
    "strconv"
)

func CreateError(code int, message string) (Error) {
    return Error{code, message}
}

func CreateNoError() (Error) {
    return Error{NoError, "No unusual behaviour"}
}

func NumberArgumentsError(expected, received int) (Error) {
    return Error{InvalidNumberOfArguments, "Expected " +
        strconv.Itoa(expected) + ", received " +
        strconv.Itoa(received) + "!"}
}

func NumberArgumentsErrorAtLeast(expected, received int) (Error) {
    return Error{InvalidNumberOfArguments, "Expected at least " +
        strconv.Itoa(expected) + ", received " +
        strconv.Itoa(received) + "!"}
}

func ArgumentTypeError(pos int, expected, received string) (Error) {
    return Error{InvalidArgumentType, "Expected argument " +
        strconv.Itoa(pos) + " of type " + expected + ", received " +
        received + "!"}
}

func UndeclaredIdentifierError(identifier string) (Error) {
    return Error{UndeclaredIdentifier, "The " +
        identifier + " identifier is undeclared in the current scope!"}
}
