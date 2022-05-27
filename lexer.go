package main

const (
	NEW_LINE_CHARACTER     rune = '\n'
	WHITESPACE_CHARACTER   rune = ' '
	EQUAL_SIGN             rune = '='
	PLUS_SIGN              rune = '+'
	MINUS_SIGN             rune = '-'
	DOUBLE_QUOTE_CHARACTER rune = '"'
)

var keywords = map[string]bool{
	"PRINT": true,
	"READ":  true,
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

func (lexer *Lexer) getFullToken() {
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
	lexer.tokens = append(lexer.tokens, token)
	lexer.currentIndex = endIndex
}

func (lexer *Lexer) getFullStringToken() {
	endIndex := lexer.getEndIndex(func(r rune) bool {
		return r != DOUBLE_QUOTE_CHARACTER
	})
	value := lexer.text[lexer.currentIndex : endIndex+2]
	token := CreateToken(value, STRING_TOKEN_TYPE)
	lexer.tokens = append(lexer.tokens, token)
	lexer.currentIndex = endIndex + 2
}

func (lexer *Lexer) Tokenize() {
	for lexer.currentIndex < len(lexer.text) {
		char := rune(lexer.text[lexer.currentIndex])
		isSkippable := char == NEW_LINE_CHARACTER || char == WHITESPACE_CHARACTER
		if !isSkippable {
			if char == EQUAL_SIGN {
				token := CreateToken(string(char), ASSIGNMENT_TOKEN_TYPE)
				lexer.tokens = append(lexer.tokens, token)
			} else if char == PLUS_SIGN {
				token := CreateToken(string(char), OPERATOR_TOKEN_TYPE)
				lexer.tokens = append(lexer.tokens, token)
			} else if char == MINUS_SIGN {
				token := CreateToken(string(char), OPERATOR_TOKEN_TYPE)
				lexer.tokens = append(lexer.tokens, token)
			} else if char == DOUBLE_QUOTE_CHARACTER {
				lexer.getFullStringToken()
			} else if isNumeric(char) {
				lexer.getFullToken()
			} else if isAlphaNumeric(char) {
				lexer.getFullToken()
			}
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
