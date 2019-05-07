package krom

import (
    generics "./generics"
)

/**
 *  Namespace implementation designed to be passed along
 *  the interpreter tree.
 */
type Scope struct {
    Args map[string]generics.Void
}

/**
 *  Returns a copy of the scope.
 */
func (s Scope) Clone() generics.Namespace {
    ret := NewScope()
    for k, v := range s.Args {
        ret.Args[k] = v
    }
    return ret
}

/**
 *  Returns a copy of the scope with the given pair
 *  inserted into the internal map.
 */
func (s Scope) Extend(key string, value generics.Void) generics.Namespace {
    ret := s.Clone().(Scope)
    ret.Args[key] = value
    return ret
}

/**
 *  Returns the value for the given key.
 */
func (s Scope) Get(key string) generics.Void {
    if value, ok := s.Args[key]; ok {
        return value
    } else {
        return nil
    }
}

func NewScope() Scope {
    aux := make(map[string]generics.Void)
    return Scope{aux}
}
