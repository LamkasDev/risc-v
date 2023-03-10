package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

func CompileStatementGroup(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	var err error
	for _, n := range statement.GetStatementGroupBody(&s) {
		d, err = CompileStatement(ccompiler, n, d)
		if err != nil {
			return d, err
		}
	}
	return d, nil
}
