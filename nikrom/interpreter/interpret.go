package intepreter

import (
	generics "../generics"
    handles "../handles"
	error "../error"
)

/**
 *	Interprets the given tree.
 */
func Interpret(tree generics.Atom) (generics.Void, error.Error) {
	var args []generics.Void

    for _, branch := range tree.Subatoms {
        arg, err := Interpret(branch)
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
