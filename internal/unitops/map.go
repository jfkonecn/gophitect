package unitops

// A Map unit operation converts one value or type into another.
type Map struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (m Map) GetUnitOperationTypeDefinition() string {
	return "A Map unit operation converts one value or type into another."
}

func (m Map) GetInputTypes() []TypeDefinition {
	return m.InputTypes
}

func (m Map) GetOutputTypes() []TypeDefinition {
	return m.OutputTypes
}

func (m Map) GetPrompt() string {
	return m.Prompt
}

func (m Map) GetCodeComments() string {
	return m.CodeComments
}

func (m Map) GetDocumentation() string {
	return m.Documentation
}

func (m Map) GetFunctionCalls() []FunctionDefinition {
	return m.FunctionCalls
}

func (m Map) GetShouldStub() bool {
	return m.ShouldStub
}
