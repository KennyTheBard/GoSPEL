package error


/**
 *	Enum with posible error codes
 */
const (
    InvalidNumberOfArguments = 1 << iota
    InvalidArgumentType = 1 << iota
)

/**
 *	Encapsulates the data related to an error raised while running.
 */
type Error struct {
    code int
    message string
}

func Error(code int, message string) (Error) {
    return Error{code, message}
}
