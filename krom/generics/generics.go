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
 *	Serves as template for all image processing functions
 *	provided by the program.
 */
type ImageProcessing func(Arguments) Void

/**
 *	Interface for elements of an interpreter tree.
 */
type InterpreterTree interface {
    Interpret(Namespace) (Void, error.Error)
}

/**
 *  Interface for quick cloning of namespaces.
 */
type Namespace interface {
    Clone() Namespace
    Extend(string, Void) Namespace
    Get(string) Void
}
