package lexer

import (
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

func CreateTokenCouple(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, bool) {
	matchedString := ConstructChunk(lexer, string(lexer.Current), func(c rune) bool {
		return unicode.IsPunct(c) || unicode.IsSymbol(c)
	})
	couple, ok := lexer.Context.Tokens.Couple[matchedString]
	if ok {
		return lexer_token.CreateToken(lexer.Position, couple), matchedString, true
	}

	return lexer_token.AwooLexerToken{}, matchedString, false
}