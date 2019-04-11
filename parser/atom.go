
/**
 *	A general interface to be passed up in the Interpreter Tree
 *	and given as arguments to internal functions.
 */
type Void interface {}

/**
 *	Structure used for source code parsing.
 */
type Atom struct {
	cmd string
	subatoms []Atoms
}

/**
 *	Evaluates the curret expression and calls a function on given
 *	arguments.
 */
func (atom *Atom) Interpret() Void {
	
}
