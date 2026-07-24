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

// A Filter unit operation removes, rejects, or reroutes data.
type Filter struct {
}

// A Sort unit operation orders a collection according to defined rules.
type Sort struct {
}

// A Distribution unit operation chooses where data goes next, such as branching.
type Distribution struct {
}

// A Validate unit operation checks data integrity.
type Validate struct {
}

// An Authenticate unit operation determines the identity initiating a flow.
type Authenticate struct {
}

// An Authorize unit operation determines the whether an authenticated identity may perform an action.
type Authorize struct {
}

// A GlobalStateRead unit operation reads values whose lifetime extends beyond the current call stack
type GlobalStateRead struct {
}

// A GlobalStateWrite unit operation writes values whose lifetime extends beyond the current call stack
type GlobalStateWrite struct {
}

// An InputOutput unit operation communicates outside the program, including HTTP, databases, etc...
type InputOutput struct {
}

// A Panic unit operation terminates execution of the program.
type Panic struct {
}
