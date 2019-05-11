package krom

import (
    generics "./generics"
)

/**
 *	Holds a pair name-expression of a defined macro.
 */
type Macro struct {
    Name string
    Expression generics.InterpreterTree
}

/**
 *	Evaluates if the macro is right for the given name.
 */
func (macro Macro) Match(name string) bool {
    return name == macro.Name
}

/**
 *	Holds all pairs of defined macros.
 */
type Defines struct {
    Macros []Macro
}

/**
 *	Finds a macro if it's defined, or returns nil if is not.
 */
func (def Defines) GetMacro(name string) generics.InterpreterTree {
    for _, macro := range def.Macros {
        if macro.Match(name) {
            return macro.Expression
        }
    }
    return nil
}

var Macros Defines
