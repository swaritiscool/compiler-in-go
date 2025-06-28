package main

import (
	"fmt"
	"log"
)

// tokenization (lexer)

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

// ------------ Parser -------------------

type node struct {
	kind       string
	name       string
	value      string
	callee     *node
	expression *node
	body       []node
	params     []node
	arguments  *[]node
	context    *[]node
}

// ast = Abstract Syntax Tree
type ast node

var pc int // parser counter

var pt []token // slice of tokens to parse

func parser(tokens []token) ast {
	pc = 0
	pt = tokens

	ast := ast{
		kind: "Program",
		body: []node{},
	}

	for pc < len(pt) {
		ast.body = append(ast.body, walk())
	}

	return ast
}

func walk() node {
	token := pt[pc]

	if token.kind == "number" {
		pc++

		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}
	}

	if token.kind == "paren" && token.value == "(" {
		pc++
		token = pt[pc]

		n := node{
			kind:   "CallExpression",
			name:   token.value,
			params: []node{},
		}

		pc++
		token = pt[pc]

		for token.kind != "param" || (token.kind == "param" && token.value == "(") {
			n.params = append(n.params, walk())
			token = pt[pc]
		}

		pc++

		return n
	}

	log.Fatal(token.kind)
	return node{}
}

func main() {
	input := "()"
	tokens := tokenizer(input)
	fmt.Println(tokens)
}
