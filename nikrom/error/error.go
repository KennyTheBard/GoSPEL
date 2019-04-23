package error


/**
 *	Enum with posible error codes
 */
const (
    NoError = 0
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

func NoError() (Error) {
    return Error{NoError, "No unusual behaviour"}
}
