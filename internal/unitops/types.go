// Package unitops
package unitops

type TypeDefinition interface {
	GetTypeName() string
	GetCodeComments() string
	GetDependentTypes() []TypeDefinition
	GetFilePath() string
	GetShouldStub() bool
}

type UnitOperation interface {
	GetUnitOperationTypeDefinition() string
	GetInputTypes() []TypeDefinition
	GetOutputTypes() []TypeDefinition
	GetPrompt() string
	GetCodeComments() string
	GetDocumentation() string
	GetFunctionCalls() []FunctionDefinition
	GetShouldStub() bool
}

type FunctionDefinition interface {
	GetFunctionName()
	GetCodeComments() string
	GetUnitOperations() []UnitOperation
	GetShouldStub() bool
}

type TestCase interface {
	GetName() string
	GetDefinitionPrompt() string
	GetCodeComments() string
	GetShouldStub() bool
}

type TestSuite interface {
	GetCodeComments() string
	GetFunctionDefinition() FunctionDefinition
	GetTestCases() []TestCase
}

type TestFileDefinition interface {
	FileDefinitionPrompt() string
	GetFilePath() string
	GetCodeComments() string
	GetTestSuites() []TestSuite
}

type ProductionFileDefinition interface {
	FileDefinitionPrompt() string
	GetFilePath() string
	GetCodeComments() string
	GetTypeDefinitions() []TypeDefinition
	GetFunctionDefinitions() []FunctionDefinition
}
