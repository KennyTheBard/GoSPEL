package krom

import (
	error "./error"
	handles "./handles"
	generics "./generics"
)

/**
 *	Structure used for source code parsing.
 */
type Atom struct {
	Process string
	Subatoms []generics.InterpreterTree
}

/**
 *	Interprets the given tree starting with arguments.
 */
func (tree Atom) Interpret(generics.Namespace) (generics.Void, error.Error) {
	// check if this is a leaf
	if len(tree.Subatoms) == 0 {
        return tree.Process, error.CreateNoError()
    }

	// obtain the right handle
	handle, err := handles.GetHandle(tree.Process)
	if err.Code != error.NoError {
		return nil, err
	}

	// process the arguments
	var args []generics.Void
    for _, branch := range tree.Subatoms {
        arg, err := branch.Interpret()
		if err.Code != error.NoError {
			return nil, err
		}
		args = append(args, arg)
    }

	// call the handle with the arguments
	ret, err := handle(args)
	if err.Code != error.NoError {
		return nil, err
	} else {
		return ret, err
	}
}
