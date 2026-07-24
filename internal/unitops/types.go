// Package unitops
package unitops

type TypeDefinition interface {
	GetPrompt() string
	GetDependentTypes() []*TypeDefinition
}

type UnitOperation interface {
	GetInputTypes() []*TypeDefinition
	GetOutputTypes() []*TypeDefinition
	GetPrompt() string
	GetFunctionCalls() []*FunctionDefinition
}

type FunctionDefinition interface {
	GetFunctionName() string
	GetPrompt() string
	GetUnitOperations() []*UnitOperation
}

type TestCase interface {
	GetName() string
	GetPrompt() string
}

type TestSuite interface {
	GetPrompt() string
	GetFunctionDefinition() *FunctionDefinition
	GetTestCases() []*TestCase
}

type TestFileDefinition interface {
	GetPrompt() string
	GetFilePath() string
	GetTestSuites() []*TestSuite
}

type ProductionFileDefinition interface {
	GetPrompt() string
	GetFilePath() string
	GetTypeDefinitions() []*TypeDefinition
	GetFunctionDefinitions() []*FunctionDefinition
}
