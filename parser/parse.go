package main

import "fmt"

func Tokenize(str string) ([]string) {
	var tokens []string
	prev := 0
	nested := 0

	for pos, ch := range str {
		if ch == '(' {
			nested += 1
		} else if ch == ')' {
			nested += -1
		} else if ch == ' ' && nested == 0{
			tokens = append(tokens, str[prev:pos])
			prev = pos + 1
		}

		if pos == len(str) - 1 {
			tokens = append(tokens, str[prev:])
		}
	}

	return tokens
}

func main() {

	str := "filter (load image.jpg) (gen_f grayscale)"

	tokens := Tokenize(str)

	for _, t := range tokens {
		fmt.Println(t)
	}
}
