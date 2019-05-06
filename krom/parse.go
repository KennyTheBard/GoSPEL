package krom

import (
	"strings"
)

const LiteralMarking = '\''
const OpenAtom = '('
const CloseAtom = ')'
const Separator = ' '

const Whitespaces = " \t\n"

/**
 *	Reduce the number of whitespaces in order to simplify the parsing process
 */
func Standardize(str string) (string) {
	whitespace := false
	literal := false
	ret := ""

	for _, ch := range str {
		// flag the current slice as literal
		if ch == LiteralMarking {
			literal = !literal
		}

		if !literal {
			// when fiding a non-whitespace character
			if strings.IndexByte(Whitespaces, byte(ch)) == -1 {

				// insert a separator before any OpenAtom
				if ch == OpenAtom && !whitespace {
					ret += string(Separator)
				}

				// insert a separater at the end of any whitespace
				// slice if it doesn't end in a CloseAtom
				if whitespace && ch != CloseAtom {
					ret += string(Separator)
				}

				// close whitespace slice and write the character
				whitespace = false
				ret += string(ch)

			// while in a whitespace slice
			} else {
				whitespace = true
			}
		}
	}

	// remove starting whitespaces
	if len(ret) > 0 && ret[0] == Separator {
		ret = ret[1:]
	}

	return ret
}

/**
 *	Break up the given string into tokens using
 *	the provided set of rules.
 */
func Tokenize(str string) ([]string) {
	var tokens []string
	prev := 0
	nest := 0
	nested := false
	literal := false

	// early escape
	if len(str) == 0 {
		return tokens
	}

	// safety
	if str[len(str) - 1:] != string(Separator) {
		str += string(Separator)
	}

	for pos, ch := range str {
		// flag the current slice as literal
		if ch == LiteralMarking {
			literal = !literal
		}

		if !literal {
			// atom opening
			if ch == OpenAtom {
				nest += 1
				nested = true
			// atom closing
			} else if ch == CloseAtom {
				nest += -1
			// atom slicing
			} else if ch == Separator && nest == 0 {
				if nested {
					tokens = append(tokens, str[prev + 1 : pos - 1])
				} else {
					tokens = append(tokens, str[prev:pos])
				}

				nested = false
				prev = pos + 1
			}
		}

	}

	if len(str[prev:]) > 0 {
		tokens = append(tokens, str[prev:])
	}

	return tokens
}

/**
 *	Parse a command into an atom tree ready
 *	to be interpreted. The leafs have only
 *	values in process field.
 */
func BuildTree(str string) (Atom) {
	if len(str) == 0 {
		aux := make([]Atom, 0)
		return Atom{"", aux}
	}

	cmd := Standardize(str)
	tokens := Tokenize(cmd)

	suba := make([]Atom, len(tokens) - 1)
	curr := Atom{tokens[0], suba}

	for i := 1; i < len(tokens); i++ {
		aux := BuildTree(tokens[i])
		if aux.Process != "" {
			curr.Subatoms[i - 1] = aux
		}
	}

	return curr
}