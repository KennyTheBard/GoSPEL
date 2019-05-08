package krom

import (
	error "./error"
	generics "./generics"
)

type Evaluation func(generics.Namespace, generics.Handle,
	 []generics.Void) (generics.Void, error.Error)

/**
 *	Interprets the given tree in the normal order
 *	allowing efficient control structures to be implemented.
 */
func normalEvaluation(namespace generics.Namespace, handle generics.Handle,
		args []generics.Void) (generics.Void, error.Error) {

    // call the handle to interpret the arguments
    return handle(append([]generics.Void{namespace}, args...))
}

/**
 *	Interprets the given tree in the reverse order
 *	allowing to interpret the arguments first.
 */
func reverseEvaluation(namespace generics.Namespace, handle generics.Handle,
		args []generics.Void) (generics.Void, error.Error) {
	// evaluate the nested atoms
    for pos, arg := range args {
		if atom, ok := arg.(Atom); ok {
			ret, err := atom.Interpret(namespace.Clone())
			if err.Code != error.NoError {
				return nil, err
			}
			args[pos] = ret
		}
    }

    // call the handle to interpret the arguments
    return handle(args)
}
