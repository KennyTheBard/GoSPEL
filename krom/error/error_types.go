package error

/**
 *	Enum with posible error codes
 */
const NoError                   = 0
const InvalidNumberOfArguments  = 1 << 0
const InvalidArgumentType       = 1 << 1
const UnknownHandle             = 1 << 2
const FileError                 = 1 << 3
const UndeclaredIdentifier      = 1 << 4
const MissingIdentifier         = 1 << 5


/**
 *	Encapsulates the data related to an error raised while running.
 */
type Error struct {
    Code int
    Message string
}
