package error


/**
 *	Enum with posible error codes
 */
const (
    NoError = 0
    InvalidNumberOfArguments = 1 << iota
    InvalidArgumentType = 1 << iota
    UnknownHandle = 1 << iota
)

/**
 *	Encapsulates the data related to an error raised while running.
 */
type Error struct {
    code int
    message string
}
