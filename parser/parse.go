package main

import "fmt"

func Tokenize(str string) ([]string) {
	var tokens []string
	prev := 0
	nest := 0
	nested := false
	literal := false

	if len(str) == 0 {
		return tokens
	}

	if str[len(str) - 1:] != " " {
		str += " "
	}

	for pos, ch := range str {
		if ch == '\'' {
			literal = !literal
		}

		if !literal {
			if ch == '(' {
				nest += 1
				nested = true
			} else if ch == ')' {
				nest += -1
			} else if ch == ' ' && nest == 0 {
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

func main() {

	str := "filter (load image.jpg) (gen_f grayscale ' :)')"

	tokens := Tokenize(str)
	fmt.Println(tokens)

	for _, t := range tokens {
		fmt.Print(t)
		fmt.Print(" -> ")
		fmt.Println()

		aux := Tokenize(t)
		for _, a := range aux {
			fmt.Println("\t", a)
		}
		fmt.Println()
	}
}
