package parser

import (
	generics "../generics"
)

const LiteralMarking = '\''
const OpenAtom = '('
const CloseAtom = ')'
const Separator = ' '

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

func BuildTree(str string) (generics.Atom) {
	if len(str) == 0 {
		aux := make([]generics.Atom, 0)
		return generics.Atom{"", aux}
	}

	tokens := Tokenize(str)

	suba := make([]generics.Atom, len(tokens) - 1)
	curr := generics.Atom{tokens[0], suba}

	for i := 1; i < len(tokens); i++ {
		aux := BuildTree(tokens[i])
		if aux.Proccess != "" {
			curr.Subatoms[i - 1] = aux
		}
	}

	return curr
}
