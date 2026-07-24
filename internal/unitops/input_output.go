package unitops

// An InputOutput unit operation communicates outside the program, including HTTP, databases, etc...
type InputOutput struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (i InputOutput) GetUnitOperationTypeDefinition() string {
	return "An InputOutput unit operation communicates outside the program, including HTTP, databases, etc..."
}

func (i InputOutput) GetInputTypes() []TypeDefinition {
	return i.InputTypes
}

func (i InputOutput) GetOutputTypes() []TypeDefinition {
	return i.OutputTypes
}

func (i InputOutput) GetPrompt() string {
	return i.Prompt
}

func (i InputOutput) GetCodeComments() string {
	return i.CodeComments
}

func (i InputOutput) GetDocumentation() string {
	return i.Documentation
}

func (i InputOutput) GetFunctionCalls() []FunctionDefinition {
	return i.FunctionCalls
}

func (i InputOutput) GetShouldStub() bool {
	return i.ShouldStub
}
