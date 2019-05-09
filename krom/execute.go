package krom

import (
    "strconv"
    error "./error"
	generics "./generics"
)

func Execute(tree generics.InterpreterTree, args []generics.Void) (generics.Void, error.Error) {
    scope := NewScope()
    for i, arg := range args {
        scope = scope.Extend(strconv.Itoa(i), arg).(Scope)
    }

    return tree.Interpret(scope)
}

func ExecuteAll(forrest []generics.InterpreterTree, args []generics.Void) (error.Error) {
    scope := NewScope()
    for i, arg := range args {
        scope = scope.Extend(strconv.Itoa(i), arg).(Scope)
    }

    for _, tree := range forrest {
        ret, err := tree.Interpret(scope)
        if err.Code != error.NoError {
            return err
        }
        if new_scope, ok := ret.(generics.Namespace); ok {
            scope = new_scope.(Scope)
        }
    }

    return error.CreateNoError()
}

func ConvertStringArguments(args []string) []generics.Void {
    ret := make([]generics.Void, len(args))
    for i, arg := range args { ret[i] = arg }
    return ret
}
