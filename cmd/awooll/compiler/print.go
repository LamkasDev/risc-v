package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintNewCompile(ccompiler *AwooCompiler, s *statement.AwooParserStatement, data []byte) {
	text := ""
	for i := 0; i < len(data); i += 4 {
		text += fmt.Sprintf("[%#x %#x %#x %#x] ", data[i], data[i+1], data[i+2], data[i+3])
	}

	logger.Log("%s %s  %s\n",
		parser.PrintStatement(&ccompiler.Settings.Parser, &ccompiler.Context.Parser, s),
		gchalk.Gray("➔"),
		gchalk.Cyan(text),
	)
}
