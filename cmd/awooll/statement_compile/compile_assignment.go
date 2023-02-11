package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementAssignment(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	variableNameNode := statement.GetStatementAssignmentIdentifier(&s)
	variableName := node.GetNodeIdentifierValue(&variableNameNode)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, variableName)
	if err != nil {
		return d, err
	}
	assignmentValueNode := statement.GetStatementAssignmentValue(&s)
	d, err = CompileNodeValue(ccompiler, assignmentValueNode, d, &details)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsSave[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type].Size],
		SourceOne:   cpu.AwooRegisterSavedZero,
		SourceTwo:   details.Register,
		Immediate:   uint32(variableMemory.Start),
	}, d)
}
