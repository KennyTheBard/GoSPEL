package generics

import (
    error "../error"
)

/**
 *  Namespace implementation designed to be passed along
 *  the interpreter tree.
 */
type Scope struct {
    Args map[string]Void
}

func (s Scope) Clone() Namespace {
    aux := make(map[string]int)

    for k, v := range s.Args {
        aux[k] = v
    }

    return Scope{aux}
}

func (s Scope) Extend(key string, value Void) Namespace {
    aux := make(map[string]int)

    for k, v := range s.Args {
        aux[k] = v
    }

    aux[key] = value

    return Scope{aux}
}
