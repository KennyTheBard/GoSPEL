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
    args = make(generics.Void, len(tree.Subatoms))
    for i, branch = range tree.Subatoms {
        (args[i], err) = Interpret(branch)
		if err.code != error.NoError {
			return (nil, err)
		}
    }

    if len(args) == 0 {
        return tree.Process
    } else {
        (handle, err) := handles.GetHandle(tree.Process)
		if err.code == error.NoError {
			return handle(args)
		}
    }
}
