package krom

import (
	error "./error"
	handles "./handles"
	generics "./generics"
)

/**
 *	Structure used for source code parsing.
 */
type Control struct {
	Process string
	Subatoms []generics.InterpreterTree
}

/**
 *	Interprets the given tree in the reverse order.
 */
func (tree Control) Interpret() (generics.Void, error.Error) {
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
    if len(args) == 0 {
        return tree.Process, error.CreateNoError()
    } else {
		ret, err := handle(args)
		if err.Code != error.NoError {
			return nil, err
		} else {
			return ret, err
		}
    }
}
