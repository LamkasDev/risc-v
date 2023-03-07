package lexer

import (
	"strings"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

func CreateTokenLetter(lexer *AwooLexer) (lexer_token.AwooLexerToken, string) {
	matchedString := ConstructChunkFast(lexer, string(lexer.Current), func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsNumber(c)
	})
	matchingKeyword, ok := lexer.Settings.Tokens.Keywords[strings.ToLower(matchedString)]
	if ok {
		return lexer_token.CreateToken(lexer.Position, matchingKeyword), matchedString
	}
	matchingType, ok := lexer_context.GetContextType(&lexer.Context, matchedString)
	if ok {
		return lexer_token.CreateTokenType(lexer.Position, matchingType.Id), matchedString
	}

	return lexer_token.CreateTokenIdentifier(lexer.Position, matchedString), matchedString
}
