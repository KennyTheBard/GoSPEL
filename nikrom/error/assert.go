package error

func AssertNumberArgument(expected, received int) (Error) {
    if expected != received {
        return NumberArgumentsError(expected, received)
    }
    return NoError()
}

func AssertArgumentType(check bool, pos int, expected, received string) (Error) {
    if check {
        return ArgumentTypeError(pos, expected, received)
    }
    return NoError()
}
