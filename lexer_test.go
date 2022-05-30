package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLexerBig(t *testing.T) {
	expectedTokens := []Token{
		{Value: "PRINT", Type: KEYWORD_TOKEN_TYPE},
		{Value: "\"1\"", Type: STRING_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "A", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "=", Type: ASSIGNMENT_TOKEN_TYPE},
		{Value: "1", Type: NUMBER_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "B", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "=", Type: ASSIGNMENT_TOKEN_TYPE},
		{Value: "1", Type: NUMBER_TOKEN_TYPE},
		{Value: "+", Type: OPERATOR_TOKEN_TYPE},
		{Value: "2", Type: NUMBER_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "PRINT", Type: KEYWORD_TOKEN_TYPE},
		{Value: "B", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "C", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "=", Type: ASSIGNMENT_TOKEN_TYPE},
		{Value: "\"keke\"", Type: STRING_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "READ", Type: KEYWORD_TOKEN_TYPE},
		{Value: "D", Type: IDENTIFIER_TOKEN_TYPE},
	}
	execTest("test_lexer_big.txt", expectedTokens, t)
}

func TestLexerIO(t *testing.T) {
	expectedTokens := []Token{
		{Value: "READ", Type: KEYWORD_TOKEN_TYPE},
		{Value: "VARIABLE", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "PRINT", Type: KEYWORD_TOKEN_TYPE},
		{Value: "VARIABLE", Type: IDENTIFIER_TOKEN_TYPE},
	}
	execTest("test_lexer_io.txt", expectedTokens, t)
}

func TestLexerOperators(t *testing.T) {
	expectedTokens := []Token{
		{Value: "A", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "=", Type: ASSIGNMENT_TOKEN_TYPE},
		{Value: "1", Type: NUMBER_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "B", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "=", Type: ASSIGNMENT_TOKEN_TYPE},
		{Value: "2", Type: NUMBER_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "C", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "=", Type: ASSIGNMENT_TOKEN_TYPE},
		{Value: "B", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "+", Type: OPERATOR_TOKEN_TYPE},
		{Value: "A", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "D", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "=", Type: ASSIGNMENT_TOKEN_TYPE},
		{Value: "B", Type: IDENTIFIER_TOKEN_TYPE},
		{Value: "-", Type: OPERATOR_TOKEN_TYPE},
		{Value: "A", Type: IDENTIFIER_TOKEN_TYPE},
	}
	execTest("test_lexer_operator.txt", expectedTokens, t)
}

func TestLexerString(t *testing.T) {
	expectedTokens := []Token{
		{Value: "\"1\"", Type: STRING_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "\"two\"", Type: STRING_TOKEN_TYPE},
		{Value: string(NEW_LINE_CHARACTER), Type: NEW_LINE_TOKEN_TYPE},
		{Value: "\"san\"", Type: STRING_TOKEN_TYPE},
	}
	execTest("test_lexer_string.txt", expectedTokens, t)
}

func execTest(codeFilename string, expectedTokens []Token, t *testing.T) {
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
		if expectedToken.Value != token.Value || expectedToken.Type != token.Type {
			message := fmt.Sprintf("expected token number %d as %s, got %s", index+1, expectedToken, token)
			t.Fatal(message)
		}
	}
}
