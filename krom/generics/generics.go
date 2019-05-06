package generics

import (
    error "../error"
)

/**
 *	A general interface to be passed up in the Interpreter Tree
 *	and given as arguments to internal functions.
 */
type Void interface {}

/**
 *	Takes place of optional arguments and provides support for
 *	generic data types passed as arguments.
 */
type Arguments struct {
	Args []Void
}

/**
 *	Serves as template for all image processing functions
 *	provided by the program.
 */
type ImageProcessing func(Arguments) Void


type InterpreterTree interface {
    Interpret() (Void, error.Error)
}
