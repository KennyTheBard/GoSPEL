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

/**
 *	Interface for elements of an interpreter tree.
 */
type InterpreterTree interface {
    Interpret(Namespace) (Void, error.Error)
    IsVariable() (string, error.Error)
}

/**
 *  Interface for quick cloning of namespaces.
 */
type Namespace interface {
    Clone() Namespace
    Extend(string, Void) Namespace
    Get(string) Void
}

/**
 *  Declare a type for all handles.
 */
type Handle func(Namespace, []Void) (Void, error.Error)
