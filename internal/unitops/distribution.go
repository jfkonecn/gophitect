package unitops

// A Distribution unit operation chooses where data goes next, such as branching.
type Distribution struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (d Distribution) GetUnitOperationTypeDefinition() string {
	return "A Distribution unit operation chooses where data goes next, such as branching."
}

func (d Distribution) GetInputTypes() []TypeDefinition {
	return d.InputTypes
}

func (d Distribution) GetOutputTypes() []TypeDefinition {
	return d.OutputTypes
}

func (d Distribution) GetPrompt() string {
	return d.Prompt
}

func (d Distribution) GetCodeComments() string {
	return d.CodeComments
}

func (d Distribution) GetDocumentation() string {
	return d.Documentation
}

func (d Distribution) GetFunctionCalls() []FunctionDefinition {
	return d.FunctionCalls
}

func (d Distribution) GetShouldStub() bool {
	return d.ShouldStub
}
