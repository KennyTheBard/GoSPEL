package krom

import (
	error "./error"
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
 *	Interprets the tree.
 */
func (tree Atom) Interpret(namespace generics.Namespace) (generics.Void, error.Error) {
	// check if this is a leaf
	if len(tree.Subatoms) == 0 {

		// leaf case of literar string
		if tree.Process[0] == LiteralMarking {
			return tree.Process[1:len(tree.Process) - 1], error.CreateNoError()

		// leaf case of NULL operator
		} else if tree.Process == "."{
			return nil, error.CreateNoError()

		// leaf case of variable
		} else if name, err := tree.IsVariable();
				err.Code == error.NoError && name != "" {
			value := namespace.Get(name)
			if value != nil {
				return value, error.CreateNoError()
			} else {
				return nil, error.CreateError(error.UndeclaredIdentifier,
					"The variable " + name + " was not declared!")
			}

		// } else if expression := macro

		// default leaf case
		} else {
			return tree.Process, error.CreateNoError()
		}
    }

	// prepare the arguments to be passed on
	args := make([]generics.Void, len(tree.Subatoms))
	for pos, branch := range tree.Subatoms {
		args[pos] = branch
	}

	// obtain the right handle and evaluate the interpretation atom
	ret := GetHandle(tree.Process)
	if handle, ok := isHandle(ret); ok {
		return handle(namespace.Clone(), args)
	} else if macro, ok := ret.(generics.InterpreterTree); ok {
		return Execute(macro, args)
	} else {
		return nil, error.CreateError(error.UnknownHandle,
            "Unknown handle name \"" + tree.Process + "\" !")
	}
}

func (tree Atom) IsVariable() (string, error.Error) {
	//check if it is a variable
	if tree.Process[0] != '$' {
        return "", error.CreateNoError()
    }

	// check if it's correctly defined
	name := tree.Process[1:]
	if len(name) < 1 {
		return "", error.CreateError(error.MissingIdentifier, "No string found after $!")
	}

	// if _, err := strconv.Atoi(v); err == nil {
	//     fmt.Printf("%q looks like a number.\n", v)
	// }
	return name, error.CreateNoError()
}

/**
 *	Used only to make a type assertion less crude.
 */
func isHandle(obj generics.Void) (generics.Handle, bool) {
	handle, ok := obj.(func(generics.Namespace, []generics.Void) (generics.Void, error.Error))

	return handle, ok
}
