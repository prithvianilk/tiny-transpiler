package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestTranspilerBig(t *testing.T) {
	expectedPythonCode := "print(1)\nprint(2)\nA=input()\nprint(1+1)\nB=1+1\nC=2+3\nprint(4+2)\n"
	err := execTest("test_files/test_transpiler_big.txt", expectedPythonCode)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTranspilerNewLines(t *testing.T) {
	expectedPythonCode := ""
	err := execTest("test_files/test_transpiler_newlines.txt", expectedPythonCode)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTranspilerIO(t *testing.T) {
	expectedPythonCode := "A=input()\nprint(A)\n"
	err := execTest("test_files/test_transpiler_io.txt", expectedPythonCode)
	if err != nil {
		t.Fatal(err)
	}
}

func execTest(codeFilename, expectedPythonCode string) error {
	code, err := readCode(codeFilename)
	if err != nil {
		panic(err)
	}
	lexer := CreateLexer(code)
	lexer.Tokenize()
	parser := CreateParser(lexer.tokens)
	parser.Parse()
	if expectedPythonCode != parser.generator.buffer {
		message := fmt.Sprintf("incorrect transpilation:\nExpected:\n%s\n~Generated Code:\n%s", expectedPythonCode, parser.generator.buffer)
		return errors.New(message)
	}
	return nil
}
