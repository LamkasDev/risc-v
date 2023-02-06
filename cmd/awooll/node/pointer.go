package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

func CreateNodePointer(t lexer_token.AwooLexerToken, value AwooParserNode) AwooParserNodeResult {
	return CreateNodeSingle(ParserNodeTypePointer, t, value)
}
