package error

func AssertNumberArgument(expected, received int) (Error) {
    if expected != received {
        return NumberArgumentsError(expected, received)
    }
    return CreateNoError()
}

func AssertNumberArgumentAtLeast(expected, received int) (Error) {
    if expected > received {
        return NumberArgumentsErrorAtLeast(expected, received)
    }
    return CreateNoError()
}

func AssertNumberArgumentControl(expected, received int) (Error) {
    if expected > received {
        return NumberArgumentsError(expected - 1, received - 1)
    }
    return CreateNoError()
}

func AssertArgumentType(check bool, pos int, expected, received string) (Error) {
    if check {
        return ArgumentTypeError(pos, expected, received)
    }
    return CreateNoError()
}
