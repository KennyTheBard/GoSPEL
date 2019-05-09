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

		// default leaf case
		} else {
			return tree.Process, error.CreateNoError()
		}
    }

	// obtain the right handle
	handle, evaluation, err := GetHandle(tree.Process)
	if err.Code != error.NoError {
		return handle, err
	}

	// process the arguments
	var args []generics.Void
    for _, branch := range tree.Subatoms {
		// replace the variable name with value from the current scope
		if name, err := branch.IsVariable(); err.Code == error.NoError && name != "" {
			value := namespace.Get(name)

			if value != nil {
				args = append(args, value)
			}
		} else {
			args = append(args, branch.(generics.Void))
		}
    }

	// evaluate the interpretation atom
	return evaluation(namespace, handle, args)
}

func (tree Atom) IsVariable() (string, error.Error) {
	// check if it's a leaf
	if len(tree.Subatoms) > 0 {
			return "", error.CreateNoError()
    }

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
