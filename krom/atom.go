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
	Subatoms []Atom
}

/**
 *	Interprets the given tree.
 */
func (tree Atom) Interpret() (generics.Void, error.Error) {
	var args []generics.Void

    for _, branch := range tree.Subatoms {
        arg, err := branch.Interpret()
		if err.Code != error.NoError {
			return nil, err
		}
		args = append(args, arg)
    }

    if len(args) == 0 {
        return tree.Process, error.CreateNoError()
    } else {
        handle, err := handles.GetHandle(tree.Process)
		if err.Code != error.NoError {
			return nil, err
		} else {
			ret, err := handle(args)
			if err.Code != error.NoError {
				return nil, err
			} else {
				return ret, err
			}
		}
    }
}
