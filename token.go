package main

type Token struct {
	Value string
	Type  string
}

const (
	KEYWORD_TOKEN_TYPE    string = "KEYWORD_TOKEN_TYPE"
	OPERATOR_TOKEN_TYPE   string = "OPERATOR_TOKEN_TYPE"
	NUMBER_TOKEN_TYPE     string = "NUMBER_TOKEN_TYPE"
	IDENTIFIER_TOKEN_TYPE string = "IDENTIFIER_TOKEN_TYPE"
	ASSIGNMENT_TOKEN_TYPE string = "ASSIGNMENT_TOKEN_TYPE"
	STRING_TOKEN_TYPE     string = "STRING_TOKEN_TYPE"
)

func CreateToken(value, type_ string) Token {
	return Token{Value: value, Type: type_}
}
