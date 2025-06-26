package main

import "fmt"

type token struct {
	kind  string
	value string
}

func tokenizer(input string) []token {
	input += "\n"
	// Making a counter (like a cursor over text) to track what we have read
	current := 0
	tokens := []token{}
	for current < len([]rune(input)) {
		char := string([]rune(input)[current])
		if char == "(" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: "(",
			})
			// Incrementing count
			current++
			continue
		}
		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: ")",
			})
			current++
			continue
		}
		if char == " " {
			current++
			continue
		}
		if isNumber(char) {
			value := ""

			for isNumber(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}

			tokens = append(tokens, token{
				kind:  "num",
				value: value,
			})
			continue
		}
		if isLetter(char) {
			value := ""

			for isLetter(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}

			tokens = append(tokens, token{
				kind:  "name",
				value: value,
			})
			continue
		}
		break
	}
	return tokens
}

func isNumber(char string) bool {
	ch := []rune(char)[0]
	if ch >= '0' && ch <= '9' {
		return true
	} else {
		return false
	}
}

func isLetter(char string) bool {
	ch := []rune(char)[0]
	if ch >= 'a' && ch <= 'z' {
		return true
	} else {
		return false
	}
}

func main() {
	input := "()"
	tokens := tokenizer(input)
	fmt.Println(tokens)
}
