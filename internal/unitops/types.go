// Package unitops
package unitops

type TypeDefinition interface {
	GetTypeName() string
	GetCodeComments() string
	GetDependentTypes() []TypeDefinition
	GetFilePath() string
}

type UnitOperation interface {
	GetUnitOperationTypeDefinition() string
	GetInputTypes() []TypeDefinition
	GetOutputTypes() []TypeDefinition
	GetPrompt() string
	GetCodeComments() string
	GetDocumentation() string
	GetFunctionCalls() []FunctionDefinition
}

type FunctionDefinition interface {
	GetFunctionName()
	GetCodeComments() string
	GetUnitOperations() []UnitOperation
}

type FileDefinition interface {
	FileDefinitionPrompt() string
	GetFilePath() string
	GetCodeComments() string
	GetTypeDefinitions() []TypeDefinition
	GetFunctionDefinitions() []FunctionDefinition
}

type Map struct {
}

type Filter struct {
}

type Sort struct {
}

type Distribution struct {
}

type Validate struct {
}

type Authenticate struct {
}

type Authorize struct {
}

type GlobalStateRead struct {
}

type GlobalStateWrite struct {
}

type InputOutput struct {
}

type Panic struct {
}
