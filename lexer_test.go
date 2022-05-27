package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLexerBig(t *testing.T) {
	expectedTokens := []string{
		"PRINT",
		"\"1\"",
		"A",
		"EQUAL",
		"1",
		"B",
		"EQUAL",
		"2",
		"C",
		"EQUAL",
		"1",
		"PLUS",
		"2",
		"PRINT",
		"C",
		"D",
		"EQUAL",
		"\"keke\"",
		"PRINT",
		"D",
	}
	execTest("test_lexer_big.txt", expectedTokens, t)
}

func TestLexerIO(t *testing.T) {
	expectedTokens := []string{"READ", "VARIABLE", "PRINT", "VARIABLE"}
	execTest("test_lexer_io.txt", expectedTokens, t)
}

func TestLexerString(t *testing.T) {
	expectedTokens := []string{
		"\"1\"",
		"\"two\"",
		"\"san\"",
	}
	execTest("test_lexer_string.txt", expectedTokens, t)
}

func execTest(codeFilename string, expectedTokens []string, t *testing.T) {
	content, err := os.ReadFile(codeFilename)
	if err != nil {
		t.Fatal(err.Error())
	}
	code := string(content)
	lexer := CreateLexer(code)
	lexer.Tokenize()
	tokens := lexer.tokens
	numberOfTokens := len(tokens)
	numberOfExpectedTokens := len(expectedTokens)
	if numberOfExpectedTokens != numberOfTokens {
		message := fmt.Sprintf("expected %d tokens, got %d tokens", numberOfExpectedTokens, numberOfTokens)
		t.Fatal(message)
	}
	for index, token := range tokens {
		expectedToken := expectedTokens[index]
		if expectedToken != token {
			message := fmt.Sprintf("expected token number %d as %s, got %s", index+1, expectedToken, token)
			t.Fatal(message)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
