package error

/**
 *	Enum with posible error codes
 */
const NoError = 0
const InvalidNumberOfArguments = 1 << 0
const InvalidArgumentType = 1 << 1
const UnknownHandle = 1 << 2
const FailedOpenFile = 1 << 3
const UndeclaredIdentifier = 1 << 4

/**
 *	Encapsulates the data related to an error raised while running.
 */
type Error struct {
    Code int
    Message string
}
