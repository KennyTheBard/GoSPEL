package aritmethic

import (
    "reflect"
    generics "../generics"
    error "../error"
)

/**
 *  Handle the arguments and evaluate a Lisp-style
 *  aritmethic expresion.
 *  Usage: define <name> (<code>)
 */
func AritmethicHandle(scope generics.Namespace, raw_args []generics.Void) (generics.Void, error.Error) {
    // check the number of arguments
    expected := 1
    received := len(raw_args)
    if expected >= received {
        return nil, error.NumberArgumentsErrorAtLeast(expected, received)
    }

    
}
