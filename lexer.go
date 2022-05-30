package main

const (
	NEW_LINE_CHARACTER     rune = '\n'
	WHITESPACE_CHARACTER   rune = ' '
	EQUAL_SIGN             rune = '='
	PLUS_SIGN              rune = '+'
	MINUS_SIGN             rune = '-'
	DOUBLE_QUOTE_CHARACTER rune = '"'
)

const (
	PRINT_KEYWORD string = "PRINT"
	READ_KEYWORD  string = "READ"
)

var keywords = map[string]bool{
	PRINT_KEYWORD: true,
	READ_KEYWORD:  true,
}

type Lexer struct {
	text         string
	tokens       []Token
	currentIndex int
}

func CreateLexer(text string) Lexer {
	return Lexer{text: text, currentIndex: 0, tokens: []Token{}}
}

func (lexer *Lexer) getEndIndex(continueCondition func(rune) bool) int {
	endIndex := lexer.currentIndex
	for next := lexer.currentIndex + 1; next < len(lexer.text); next++ {
		nextChar := lexer.text[next]
		if !continueCondition(rune(nextChar)) {
			break
		}
		endIndex = next
	}
	return endIndex
}

func (lexer *Lexer) getFullToken() Token {
	var token Token
	var endIndex int
	startingToken := rune(lexer.text[lexer.currentIndex])
	isNumericToken := isNumeric(startingToken)
	if isNumericToken {
		endIndex = lexer.getEndIndex(isNumeric)
		token.Value = lexer.text[lexer.currentIndex : endIndex+1]
		token.Type = NUMBER_TOKEN_TYPE
	} else {
		endIndex = lexer.getEndIndex(isAlphaNumeric)
		token.Value = lexer.text[lexer.currentIndex : endIndex+1]
		_, isKeyword := keywords[token.Value]
		if isKeyword {
			token.Type = KEYWORD_TOKEN_TYPE
		} else {
			token.Type = IDENTIFIER_TOKEN_TYPE
		}
	}
	lexer.currentIndex = endIndex
	return token
}

func (lexer *Lexer) getFullStringToken() Token {
	endIndex := lexer.getEndIndex(func(r rune) bool {
		return r != DOUBLE_QUOTE_CHARACTER
	})
	value := lexer.text[lexer.currentIndex : endIndex+2]
	token := CreateToken(value, STRING_TOKEN_TYPE)
	lexer.currentIndex = endIndex + 1
	return token
}

func (lexer *Lexer) Tokenize() {
	for lexer.currentIndex < len(lexer.text) {
		var token Token
		char := rune(lexer.text[lexer.currentIndex])
		isSkippable := char == WHITESPACE_CHARACTER
		if !isSkippable {
			value := string(char)
			if char == NEW_LINE_CHARACTER {
				token = CreateToken(value, NEW_LINE_TOKEN_TYPE)
			} else if char == EQUAL_SIGN {
				token = CreateToken(value, ASSIGNMENT_TOKEN_TYPE)
			} else if char == PLUS_SIGN {
				token = CreateToken(value, OPERATOR_TOKEN_TYPE)
			} else if char == MINUS_SIGN {
				token = CreateToken(value, OPERATOR_TOKEN_TYPE)
			} else if char == DOUBLE_QUOTE_CHARACTER {
				token = lexer.getFullStringToken()
			} else if isNumeric(char) {
				token = lexer.getFullToken()
			} else if isAlphaNumeric(char) {
				token = lexer.getFullToken()
			}
			lexer.tokens = append(lexer.tokens, token)
		}
		lexer.currentIndex++
	}
}

func isNumeric(char rune) bool {
	isNumeric := (char <= '9' && char >= '0')
	return isNumeric
}

func isAlphaNumeric(char rune) bool {
	isAlpha := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
	return isAlpha || isNumeric(char)
}
