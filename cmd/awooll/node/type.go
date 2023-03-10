package node

import "github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"

type AwooParserNodeDataType struct {
	Value uint16
}

func GetNodeTypeType(n *AwooParserNode) uint16 {
	return n.Data.(AwooParserNodeDataType).Value
}

func SetNodeTypeType(n *AwooParserNode, value uint16) {
	d := n.Data.(AwooParserNodeDataType)
	d.Value = value
	n.Data = d
}

func CreateNodeType(t lexer_token.AwooLexerToken) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeType,
			Token: t,
			Data: AwooParserNodeDataType{
				Value: lexer_token.GetTokenTypeId(&t),
			},
		},
	}
}
