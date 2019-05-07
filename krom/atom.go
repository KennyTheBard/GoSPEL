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
 *	Interprets the given tree starting with arguments.
 */
func (tree Atom) Interpret(namespace generics.Namespace) (generics.Void, error.Error) {
	// check if this is a leaf
	if len(tree.Subatoms) == 0 {
        return tree.Process, error.CreateNoError()
    }

	// obtain the right handle
	handle, err := GetHandle(tree.Process)
	if err.Code != error.NoError {
		return nil, err
	}

	// process the arguments
	var args []generics.Void
    for _, branch := range tree.Subatoms {
		// replace the variable name with value from the current scope
		if name, err := isVariable(branch); err.Code == error.NoError && name != ""{
			value := namespace.Get(name)

			if value != nil {
				args = append(args, value)
			}
		} else {
			// interpret subtrees
			arg, err := branch.Interpret(namespace.Clone())
			if err.Code != error.NoError {
				return nil, err
			}
			args = append(args, arg)
		}
    }

	// call the handle with the arguments
	ret, err := handle(args)
	if err.Code != error.NoError {
		return nil, err
	} else {
		return ret, err
	}
}

func isVariable(elem generics.Void) (string, error.Error) {
	str, ok := elem.(string)
    if !ok {
		return "", error.CreateNoError()
	}
    if len(str) < 1 {
        return "", error.CreateNoError()
    }
	if str[0] != '$' {
        return "", error.CreateNoError()
    }
	name := str[1:]
	if len(name) < 1 {
		return "", error.CreateError(error.MissingIdentifier, "No string found after $!")
	}
	// if _, err := strconv.Atoi(v); err == nil {
	//     fmt.Printf("%q looks like a number.\n", v)
	// }
	return name, error.CreateNoError()
}
