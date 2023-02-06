package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementDefinition(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	tNode := statement.GetStatementDefinitionVariableType(&s)
	// TODO: make this shit handler poitners.
	t := node.GetNodeTypeType(&tNode)
	nameNode := statement.GetStatementDefinitionVariableIdentifier(&s)
	name := node.GetNodeIdentifierValue(&nameNode)
	valueNode := statement.GetStatementDefinitionVariableValue(&s)
	dest, err := compiler_context.PushCompilerScopeCurrentMemory(context, name, t)
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToCompileStatement, err)
	}
	d, err = CompileNodeValueFast(context, valueNode, d)
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToCompileStatement, err)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSW,
		SourceTwo:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(dest),
	}, d)
}
