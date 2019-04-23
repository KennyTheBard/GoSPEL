package intepreter

import (
	generics "../generics"
    handles "../handles"
)

/**
 *	Interprets the given tree.
 */
func Interpret(tree generics.Atom) (generics.Void) {
    args = make(generics.Void, len(tree.Subatoms))
    for i, branch = range tree.Subatoms {
        args[i] = Interpret(branch)
    }

    if len(args) == 0 {
        return tree.Process
    } else {
        return handles.GetHandle(tree.Process)
    }
}
