package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementIfOuter(cparser *parser.AwooParser, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	n, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		EndToken: token.TokenTypeBracketCurlyRight,
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	ifStatement := statement.CreateStatementIf(n.Node)
	ifGroup, err := ConstructStatementGroup(cparser, details)
	if err != nil {
		return ifStatement, err
	}
	statement.SetStatementIfBody(&ifStatement, ifGroup)

	return ifStatement, nil
}

func ConstructStatementIf(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	ifStatement, err := ConstructStatementIfOuter(cparser, details)
	if err != nil {
		return ifStatement, err
	}
	for elseToken, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeElse); elseToken != nil; elseToken, _ = parser.ExpectTokenOptional(cparser, token.TokenTypeElse) {
		t, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypeIf, token.TokenTypeBracketCurlyLeft})
		if err != nil {
			return ifStatement, err
		}
		switch t.Type {
		case token.TokenTypeIf:
			elifStatement, err := ConstructStatementIfOuter(cparser, details)
			if err != nil {
				return ifStatement, err
			}
			statement.SetStatementIfElse(&ifStatement, append(statement.GetStatementIfElse(&ifStatement), elifStatement))
		case token.TokenTypeBracketCurlyLeft:
			elseStatement, err := ConstructStatementGroup(cparser, details)
			if err != nil {
				return ifStatement, err
			}
			statement.SetStatementIfElse(&ifStatement, append(statement.GetStatementIfElse(&ifStatement), elseStatement))
		}
	}

	return ifStatement, nil
}
