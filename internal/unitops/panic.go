package unitops

// A Panic unit operation terminates execution of the program.
type Panic struct {
	InputTypes    []TypeDefinition
	OutputTypes   []TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []FunctionDefinition
	ShouldStub    bool
}

func (p Panic) GetUnitOperationTypeDefinition() string {
	return "A Panic unit operation terminates execution of the program."
}

func (p Panic) GetInputTypes() []TypeDefinition {
	return p.InputTypes
}

func (p Panic) GetOutputTypes() []TypeDefinition {
	return p.OutputTypes
}

func (p Panic) GetPrompt() string {
	return p.Prompt
}

func (p Panic) GetCodeComments() string {
	return p.CodeComments
}

func (p Panic) GetDocumentation() string {
	return p.Documentation
}

func (p Panic) GetFunctionCalls() []FunctionDefinition {
	return p.FunctionCalls
}

func (p Panic) GetShouldStub() bool {
	return p.ShouldStub
}
