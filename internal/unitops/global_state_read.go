package unitops

// A GlobalStateRead unit operation reads values whose lifetime extends beyond the current call stack
type GlobalStateRead struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (g GlobalStateRead) GetUnitOperationTypeDefinition() string {
	return "A GlobalStateRead unit operation reads values whose lifetime extends beyond the current call stack"
}

func (g GlobalStateRead) GetInputTypes() []TypeDefinition {
	return g.InputTypes
}

func (g GlobalStateRead) GetOutputTypes() []TypeDefinition {
	return g.OutputTypes
}

func (g GlobalStateRead) GetPrompt() string {
	return g.Prompt
}

func (g GlobalStateRead) GetCodeComments() string {
	return g.CodeComments
}

func (g GlobalStateRead) GetDocumentation() string {
	return g.Documentation
}

func (g GlobalStateRead) GetFunctionCalls() []FunctionDefinition {
	return g.FunctionCalls
}

func (g GlobalStateRead) GetShouldStub() bool {
	return g.ShouldStub
}
