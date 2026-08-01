// Package unitops
package unitops

type TypeDefinition interface {
	GetID() string
	GetPrompt() (string, error)
	GetDependentTypes() []*TypeDefinition
	IsBuiltin() bool
	GetTypeName() string
	GetFilePath() string
}
type VariableDefinition interface {
	GetVariableName() string
	GetTypeDefinition() *TypeDefinition
}

type UnitOperation interface {
	GetInputTypes() []*VariableDefinition
	GetOutputTypes() []*VariableDefinition
	GetPrompt() (string, error)
	GetFunctionCalls() []*FunctionDefinition
}

type FunctionDefinition interface {
	GetID() string
	GetFunctionName() string
	GetFilePath() string
	GetPrompt() (string, error)
	GetInputs() []*VariableDefinition
	GetOutput() *TypeDefinition
	GetUnitOperations() []*UnitOperation
}

type TestCase interface {
	GetName() string
	GetPrompt() (string, error)
}

type TestSuite interface {
	GetID() string
	GetPrompt() (string, error)
	GetFilePath() string
	GetFunctionDefinition() *FunctionDefinition
	GetTestCases() []*TestCase
}
