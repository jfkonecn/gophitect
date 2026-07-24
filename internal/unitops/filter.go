package unitops

// A Filter unit operation removes, rejects, or reroutes data.
type Filter struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (f Filter) GetUnitOperationTypeDefinition() string {
	return "A Filter unit operation removes, rejects, or reroutes data."
}

func (f Filter) GetInputTypes() []TypeDefinition {
	return f.InputTypes
}

func (f Filter) GetOutputTypes() []TypeDefinition {
	return f.OutputTypes
}

func (f Filter) GetPrompt() string {
	return f.Prompt
}

func (f Filter) GetCodeComments() string {
	return f.CodeComments
}

func (f Filter) GetDocumentation() string {
	return f.Documentation
}

func (f Filter) GetFunctionCalls() []FunctionDefinition {
	return f.FunctionCalls
}

func (f Filter) GetShouldStub() bool {
	return f.ShouldStub
}
