package unitops

// A Validate unit operation checks data integrity.
type Validate struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (v Validate) GetUnitOperationTypeDefinition() string {
	return "A Validate unit operation checks data integrity."
}

func (v Validate) GetInputTypes() []TypeDefinition {
	return v.InputTypes
}

func (v Validate) GetOutputTypes() []TypeDefinition {
	return v.OutputTypes
}

func (v Validate) GetPrompt() string {
	return v.Prompt
}

func (v Validate) GetCodeComments() string {
	return v.CodeComments
}

func (v Validate) GetDocumentation() string {
	return v.Documentation
}

func (v Validate) GetFunctionCalls() []FunctionDefinition {
	return v.FunctionCalls
}

func (v Validate) GetShouldStub() bool {
	return v.ShouldStub
}
