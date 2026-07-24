// Package unitops
package unitops

type TypeDefinition interface {
	GetPrompt() (string, error)
	GetDependentTypes() []*TypeDefinition
}

type UnitOperation interface {
	GetInputTypes() []*TypeDefinition
	GetOutputTypes() []*TypeDefinition
	GetPrompt() (string, error)
	GetFunctionCalls() []*FunctionDefinition
}

type FunctionDefinition interface {
	GetFunctionName() string
	GetPrompt() (string, error)
	GetUnitOperations() []*UnitOperation
}

type TestCase interface {
	GetName() string
	GetPrompt() (string, error)
}

type TestSuite interface {
	GetPrompt() (string, error)
	GetFunctionDefinition() *FunctionDefinition
	GetTestCases() []*TestCase
}

type TestFileDefinition interface {
	GetPrompt() (string, error)
	GetFilePath() string
	GetTestSuites() []*TestSuite
}

type ProductionFileDefinition interface {
	GetPrompt() (string, error)
	GetFilePath() string
	GetTypeDefinitions() []*TypeDefinition
	GetFunctionDefinitions() []*FunctionDefinition
}
