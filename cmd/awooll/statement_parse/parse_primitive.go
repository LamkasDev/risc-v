package statement_parse

import (
	"fmt"
	"math"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
	"golang.org/x/exp/constraints"
)

// TODO: make this generic (to be used in primitive nodes also?).
func GetPrimitiveValue[K constraints.Integer](context lexer_context.AwooLexerContext, t lexer_token.AwooLexerToken) K {
	primType := context.Types.All[lexer_token.GetTokenPrimitiveType(&t)]
	switch primType.Size {
	case 1:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(int8))
		} else {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(uint8))
		}
	case 2:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(int16))
		} else {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(uint16))
		}
	case 4:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(int32))
		} else {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(uint32))
		}
	case 8:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(int64))
		} else {
			return K(lexer_token.GetTokenPrimitiveValue(&t).(uint64))
		}
	}

	return K(0)
}

func GetPrimitiveLimits(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (int64, int64) {
	primType := cparser.Context.Lexer.Types.All[lexer_token.GetTokenPrimitiveType(&t)]
	primBytes := float64(8 * primType.Size)
	if primType.Flags&types.AwooTypeFlagsSign == 1 {
		return int64(math.Pow(2, primBytes-1)) - 1, -int64(math.Pow(2, primBytes-1))
	} else {
		return int64(math.Pow(2, primBytes)) - 1, 0
	}
}

func CreateNodePrimitiveSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	// TODO: handle unsigned.
	primValue := GetPrimitiveValue[int64](cparser.Context.Lexer, t)
	primUp, primDown := GetPrimitiveLimits(cparser, t)
	if primValue > primUp {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s > %s", awerrors.ErrorPrimitiveOverflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(primUp)))
	}
	if primValue < primDown {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s < %s", awerrors.ErrorPrimitiveUnderflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(primDown)))
	}

	return node.CreateNodePrimitive(t), nil
}
