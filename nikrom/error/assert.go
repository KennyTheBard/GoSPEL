package error

import (
    "reflect"
)

func AssertNumberArgument(expected, received int) (Error) {
    if expected != received {
        return NumberArgumentsError(expected, received)
    }
    return NoError()
}

func AssertArgumentType(pos int, expected, received generics.Void) (Error) {
    // TODO
}
