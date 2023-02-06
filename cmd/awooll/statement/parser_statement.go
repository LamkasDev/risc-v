package statement

const ParserStatementTypeDefinitionVariable = 0x000
const ParserStatementTypeAssignment = 0x001
const ParserStatementTypeDefinitionType = 0x002
const ParserStatementTypeIf = 0x003
const ParserStatementTypeGroup = 0x004

type AwooParserStatement struct {
	Type uint16
	Data interface{}
}
