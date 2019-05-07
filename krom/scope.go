package krom

import (
    error "./error"
    generics "./generics"
)

/**
 *  Namespace implementation designed to be passed along
 *  the interpreter tree.
 */
type Scope struct {
    Args map[string]Void
}

/**
 *  Returns a copy of the scope.
 */
func (s Scope) Clone() generics.Namespace {
    aux := make(map[string]int)
    for k, v := range s.Args {
        aux[k] = v
    }
    return Scope{aux}
}

/**
 *  Returns a copy of the scope with the given pair
 *  inserted into the internal map.
 */
func (s Scope) Extend(key string, value generics.Void) generics.Namespace {
    aux := s.Clone()
    aux.Args[key] = value
    return aux
}

<<<<<<< HEAD
/**
 *  Returns the value for the given key.
 */
=======
>>>>>>> e0b2bc5ba9fddeda87a179bfba8b78fa861fe556
func (s Scope) Get(key string) generics.Void {
    if value, ok := s.Args[key]; ok {
        return value
    } else {
        return nil
    }
}
