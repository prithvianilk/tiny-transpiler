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
	EQUAL_TOKEN        string = "EQUAL"
	PLUS_TOKEN         string = "PLUS"
	MINUS_TOKEN        string = "MINUS"
	DOUBLE_QUOTE_TOKEN string = "DOUBLE_QUOTE"
)

var keywords = map[string]bool{
	"PRINT": true,
	"READ":  true,
}

type Lexer struct {
	text         string
	tokens       []string
	currentIndex int
}

func CreateLexer(text string) Lexer {
	return Lexer{text: text, currentIndex: 0, tokens: []string{}}
}

func (lexer *Lexer) getFullToken(continueCondition func(rune) bool) {
	endIndex := lexer.currentIndex
	for next := lexer.currentIndex + 1; next < len(lexer.text); next++ {
		nextChar := lexer.text[next]
		if !continueCondition(rune(nextChar)) {
			break
		}
		endIndex = next
	}
	identifier := lexer.text[lexer.currentIndex : endIndex+1]
	lexer.tokens = append(lexer.tokens, identifier)
	lexer.currentIndex = endIndex
}

func (lexer *Lexer) Lex() {
	for lexer.currentIndex < len(lexer.text) {
		char := rune(lexer.text[lexer.currentIndex])
		isSkippable := char == NEW_LINE_CHARACTER || char == WHITESPACE_CHARACTER
		if !isSkippable {
			if char == EQUAL_SIGN {
				lexer.tokens = append(lexer.tokens, EQUAL_TOKEN)
			} else if char == PLUS_SIGN {
				lexer.tokens = append(lexer.tokens, PLUS_TOKEN)
			} else if char == MINUS_SIGN {
				lexer.tokens = append(lexer.tokens, MINUS_TOKEN)
			} else if char == DOUBLE_QUOTE_CHARACTER {
				lexer.tokens = append(lexer.tokens, DOUBLE_QUOTE_TOKEN)
			} else if isNumeric(char) {
				lexer.getFullToken(isNumeric)
			} else if isAlphaNumeric(char) {
				lexer.getFullToken(isAlphaNumeric)
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
