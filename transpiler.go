package main

import (
	"os"
)

func readCode(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	code := string(content)
	return code, nil
}

func main() {
	filename := os.Args[1]
	codeFilename := os.Args[2]
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	text := string(content)
	lexer := CreateLexer(text)
	lexer.Tokenize()
	parser := CreateParser(lexer.tokens)
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	file, err := os.Create(codeFilename)
	if err != nil {
		panic(err)
	}
	file.WriteString(parser.generator.buffer)
	if file.Close() != nil {
		panic(err)
	}
}
