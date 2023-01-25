package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeIdentifier(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	id := node.GetNodeIdentifierValue(&n)
	src, _ := compiler_context.GetCompilerScopeMemory(context, id)
	if details.Expression {
		return encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionLW,
			Destination: cpu.AwooRegisterTemporaryOne,
			Immediate:   uint32(src),
		}, d)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionLW,
		Destination: cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(src),
	}, d)
}
