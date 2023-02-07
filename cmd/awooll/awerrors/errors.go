package awerrors

import "errors"

var ErrorFailedToSelectProgram = errors.New("failed to select program")
var ErrorPrimitiveOverflow = errors.New("primitive overflow")
var ErrorPrimitiveUnderflow = errors.New("primitive underflow")
var ErrorUnknownVariable = errors.New("unknown variable")
var ErrorAlreadyDefinedVariable = errors.New("already defined variable")
var ErrorUnknownFunction = errors.New("unknown function")
var ErrorIllegalCharacter = errors.New("illegal characters")
var ErrorExpectedToken = errors.New("expected one of")
var ErrorUnexpectedToken = errors.New("unexpected token")
var ErrorExpectedStatement = errors.New("expected statement")
var ErrorUnexpectedStatement = errors.New("unexpected statement")
var ErrorNoMoreTokens = errors.New("no more tokens")
var ErrorFailedToParse = errors.New("failed to parse")
var ErrorFailedToGetVariableFromScope = errors.New("failed to get variable from scope")
var ErrorFailedToPushVariableIntoScope = errors.New("failed to push variable into scope")
var ErrorFailedToGetFunctionFromScope = errors.New("failed to get function from scope")
var ErrorFailedToPushFunctionIntoScope = errors.New("failed to push function into scope")
var ErrorFailedToEncodeInstruction = errors.New("failed to encode instruction")
var ErrorCantCompileOperator = errors.New("no idea how to compile operator")
var ErrorCantCompileNode = errors.New("no idea how to compile node")
var ErrorCantCompileStatement = errors.New("no idea how to compile statement")
var ErrorFailedToCompileOperator = errors.New("failed to compile operator")
var ErrorFailedToCompileNode = errors.New("failed to compile node")
var ErrorFailedToCompileStatement = errors.New("failed to compile statement")
var ErrorFailedToConstructNode = errors.New("failed to construct node")
var ErrorFailedToConstructExpression = errors.New("failed to construct expression")
var ErrorFailedToConstructStatement = errors.New("failed to construct statement")
var ErrorFailedToCompileProgramHeader = errors.New("failed to compile program header")
