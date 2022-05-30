package main

import (
	"errors"
	"fmt"
)

type Parser struct {
	tokens       []Token
	currentIndex int
	generator    Generator
}

func CreateParser(tokens []Token) Parser {
	return Parser{tokens: tokens, currentIndex: 0, generator: Generator{buffer: ""}}
}

func (parser *Parser) Parse() error {
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
	token := parser.getCurrentToken()
	isPrint := token.Value == PRINT_KEYWORD
	isRead := token.Value == READ_KEYWORD
	isKeyword := token.Type == KEYWORD_TOKEN_TYPE
	if isKeyword {
		if isPrint {
			parser.generator.WriteString("print(")
			parser.currentIndex++
			if parser.value() {
				parser.generator.WriteString(")\n")
				return true
			}
			return false
		} else if isRead {
			parser.currentIndex++
			token = parser.getCurrentToken()
			parser.currentIndex++
			nextCharIsVariable := token.Type == IDENTIFIER_TOKEN_TYPE
			code := fmt.Sprintf("%s=input()\n", token.Value)
			parser.generator.WriteString(code)
			return nextCharIsVariable
		}
	}
	return parser.assignment()
}

func (parser *Parser) assignment() bool {
	token := parser.getCurrentToken()
	isVariable := token.Type == IDENTIFIER_TOKEN_TYPE
	if isVariable {
		parser.generator.WriteString(token.Value)
		parser.currentIndex++
		token = parser.getCurrentToken()
		if token.Type == ASSIGNMENT_TOKEN_TYPE {
			parser.generator.WriteString("=")
			parser.currentIndex++
			if parser.value() {
				parser.generator.WriteString("\n")
				return true
			}
		}
	}
	return false
}

func (parser *Parser) value() bool {
	token := parser.getCurrentToken()
	isVariable := token.Type == IDENTIFIER_TOKEN_TYPE
	isStringValue := token.Type == STRING_TOKEN_TYPE
	if isVariable || isStringValue {
		parser.generator.WriteString(token.Value)
		parser.currentIndex++
		return true
	}
	return parser.numericBinaryOperation()
}

func (parser *Parser) numericBinaryOperation() bool {
	token := parser.getCurrentToken()
	isNumeric := token.Type == NUMBER_TOKEN_TYPE
	if isNumeric {
		parser.generator.WriteString(token.Value)
		parser.currentIndex++
		token = parser.getCurrentToken()
		isOperation := token.Type == OPERATOR_TOKEN_TYPE
		if isOperation {
			parser.generator.WriteString(token.Value)
			parser.currentIndex++
			token = parser.getCurrentToken()
			isNumeric = token.Type == NUMBER_TOKEN_TYPE
			parser.generator.WriteString(token.Value)
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
