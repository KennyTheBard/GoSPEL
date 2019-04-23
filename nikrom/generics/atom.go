package generics

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
 *	Structure used for source code parsing.
 */
type Atom struct {
	Process string
	Subatoms []Atom
}

/**
 *	Evaluates the curret expression and calls a function on given
 *	arguments.
 */
func (atom Atom) Interpret() Void {
	// TODO: move this function to another file
	return nil
}
