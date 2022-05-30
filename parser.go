package main

import (
	"errors"
	"fmt"
)

type Parser struct {
	tokens       []Token
	currentIndex int
}

func CreateParser(tokens []Token) Parser {
	return Parser{tokens: tokens, currentIndex: 0}
}

func (parser *Parser) Parse() error {
	fmt.Println("S")
	if parser.isLambda() {
		return nil
	}
	if parser.lines() {
		return parser.Parse()
	} else if parser.statement() {
		return parser.Parse()
	}
	return errors.New("error while parsing tokens")
}

func (parser *Parser) lines() bool {
	fmt.Println("LINES")
	if parser.isLambda() {
		return true
	}
	token := parser.getCurrentToken()
	if token.Type == NEW_LINE_TOKEN_TYPE {
		parser.currentIndex++
		return parser.lines()
	}
	return false
}

func (parser *Parser) statement() bool {
	fmt.Println("STMT")
	token := parser.getCurrentToken()
	isPrint := token.Value == PRINT_KEYWORD
	isRead := token.Value == READ_KEYWORD
	isKeyword := token.Type == KEYWORD_TOKEN_TYPE
	if isKeyword {
		if isPrint {
			parser.currentIndex++
			return parser.value()
		} else if isRead {
			parser.currentIndex++
			token = parser.getCurrentToken()
			parser.currentIndex++
			nextCharIsVariable := token.Type == IDENTIFIER_TOKEN_TYPE
			return nextCharIsVariable
		}
	}
	return parser.assignment()
}

func (parser *Parser) assignment() bool {
	fmt.Println("ASS")
	token := parser.getCurrentToken()
	isVariable := token.Type == IDENTIFIER_TOKEN_TYPE
	if isVariable {
		parser.currentIndex++
		token = parser.getCurrentToken()
		if token.Type == ASSIGNMENT_TOKEN_TYPE {
			parser.currentIndex++
			return parser.value()
		}
	}
	return false
}

func (parser *Parser) value() bool {
	fmt.Println("VALUE")
	token := parser.getCurrentToken()
	isVariable := token.Type == IDENTIFIER_TOKEN_TYPE
	isStringValue := token.Type == STRING_TOKEN_TYPE
	if isVariable || isStringValue {
		parser.currentIndex++
		return true
	}
	return parser.numericBinaryOperation()
}

func (parser *Parser) numericBinaryOperation() bool {
	fmt.Println("NUMERIC BINARY OP")
	token := parser.getCurrentToken()
	isNumeric := token.Type == NUMBER_TOKEN_TYPE
	if isNumeric {
		parser.currentIndex++
		token = parser.getCurrentToken()
		isOperation := token.Type == OPERATOR_TOKEN_TYPE
		if isOperation {
			parser.currentIndex++
			token = parser.getCurrentToken()
			isNumeric = token.Type == NUMBER_TOKEN_TYPE
			parser.currentIndex++
			return isNumeric
		}
		return true
	}
	return false
}

func (parser *Parser) getCurrentToken() Token {
	return parser.tokens[parser.currentIndex]
}

func (parser *Parser) isLambda() bool {
	isLambda := parser.currentIndex == len(parser.tokens)
	return isLambda
}
