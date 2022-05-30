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
	for index, token := range lexer.tokens {
		fmt.Printf("token %d: %s\n", index+1, token)
	}
	fmt.Println()
	parser := CreateParser(lexer.tokens)
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
}
