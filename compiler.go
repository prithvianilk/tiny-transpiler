package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	text := string(content)
	lexer := CreateLexer(text)
	lexer.Tokenize()
	for n, r := range lexer.tokens {
		fmt.Printf("token %d: %s\n", n+1, string(r))
	}
}
