package unitops

// A Sort unit operation orders a collection according to defined rules.
type Sort struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (s Sort) GetUnitOperationTypeDefinition() string {
	return "A Sort unit operation orders a collection according to defined rules."
}

func (s Sort) GetInputTypes() []TypeDefinition {
	return s.InputTypes
}

func (s Sort) GetOutputTypes() []TypeDefinition {
	return s.OutputTypes
}

func (s Sort) GetPrompt() string {
	return s.Prompt
}

func (s Sort) GetCodeComments() string {
	return s.CodeComments
}

func (s Sort) GetDocumentation() string {
	return s.Documentation
}

func (s Sort) GetFunctionCalls() []FunctionDefinition {
	return s.FunctionCalls
}

func (s Sort) GetShouldStub() bool {
	return s.ShouldStub
}
