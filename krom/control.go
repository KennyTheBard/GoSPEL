package krom

import (
	error "./error"
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
    // check if this is a leaf
    if len(tree.Subatoms) == 0 {
        return tree.Process, error.CreateNoError()
    }

    // obtain the right handle
    handle, err := GetHandle(tree.Process)
    if err.Code != error.NoError {
        return nil, err
    }

	var args []generics.Void
    for _, branch := range tree.Subatoms {
		args = append(args, branch)
    }

    // call the handle to interpret the arguments
    ret, err := handle(args)
    if err.Code != error.NoError {
        return nil, err
    } else {
        return ret, err
    }
}
