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

func AssertArgumentType(check bool, pos int, expected, received string) (Error) {
    if check {
        return ArgumentTypeError(pos, expected, received)
    }
    return CreateNoError()
}
