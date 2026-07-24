package unitops

// An Authenticate unit operation determines the identity initiating a flow.
type Authenticate struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (a Authenticate) GetUnitOperationTypeDefinition() string {
	return "An Authenticate unit operation determines the identity initiating a flow."
}

func (a Authenticate) GetInputTypes() []TypeDefinition {
	return a.InputTypes
}

func (a Authenticate) GetOutputTypes() []TypeDefinition {
	return a.OutputTypes
}

func (a Authenticate) GetPrompt() string {
	return a.Prompt
}

func (a Authenticate) GetCodeComments() string {
	return a.CodeComments
}

func (a Authenticate) GetDocumentation() string {
	return a.Documentation
}

func (a Authenticate) GetFunctionCalls() []FunctionDefinition {
	return a.FunctionCalls
}

func (a Authenticate) GetShouldStub() bool {
	return a.ShouldStub
}
