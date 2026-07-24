package unitops

// A GlobalStateWrite unit operation writes values whose lifetime extends beyond the current call stack
type GlobalStateWrite struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (g GlobalStateWrite) GetUnitOperationTypeDefinition() string {
	return "A GlobalStateWrite unit operation writes values whose lifetime extends beyond the current call stack"
}

func (g GlobalStateWrite) GetInputTypes() []TypeDefinition {
	return g.InputTypes
}

func (g GlobalStateWrite) GetOutputTypes() []TypeDefinition {
	return g.OutputTypes
}

func (g GlobalStateWrite) GetPrompt() string {
	return g.Prompt
}

func (g GlobalStateWrite) GetCodeComments() string {
	return g.CodeComments
}

func (g GlobalStateWrite) GetDocumentation() string {
	return g.Documentation
}

func (g GlobalStateWrite) GetFunctionCalls() []FunctionDefinition {
	return g.FunctionCalls
}

func (g GlobalStateWrite) GetShouldStub() bool {
	return g.ShouldStub
}
