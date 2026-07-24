package unitops

// An Authorize unit operation determines the whether an authenticated identity may perform an action.
type Authorize struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (a Authorize) GetUnitOperationTypeDefinition() string {
	return "An Authorize unit operation determines the whether an authenticated identity may perform an action."
}

func (a Authorize) GetInputTypes() []TypeDefinition {
	return a.InputTypes
}

func (a Authorize) GetOutputTypes() []TypeDefinition {
	return a.OutputTypes
}

func (a Authorize) GetPrompt() string {
	return a.Prompt
}

func (a Authorize) GetCodeComments() string {
	return a.CodeComments
}

func (a Authorize) GetDocumentation() string {
	return a.Documentation
}

func (a Authorize) GetFunctionCalls() []FunctionDefinition {
	return a.FunctionCalls
}

func (a Authorize) GetShouldStub() bool {
	return a.ShouldStub
}
