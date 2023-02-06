package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementAssignment(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (statement.AwooParserStatement, error) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	identifierVariable, ok := parser_context.GetContextVariable(&cparser.Context, identifier)
	if !ok {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(identifier))
	}
	n := node.CreateNodeIdentifier(t)
	asStatement := statement.CreateStatementAssignment(n.Node)
	_, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	n, err = ConstructExpressionStart(cparser, &ConstructExpressionDetails{Type: cparser.Context.Lexer.Types.All[identifierVariable.Type]})
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	statement.SetStatementAssignmentValue(&asStatement, n.Node)

	return asStatement, nil
}
