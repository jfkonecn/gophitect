package unitops

type BasicFunctionDefinition struct {
	FunctionName   string
	CodeComments   string
	UnitOperations []UnitOperation
	ShouldStub     bool
}

func (f BasicFunctionDefinition) GetFunctionName() string {
	return f.FunctionName
}

func (f BasicFunctionDefinition) GetCodeComments() string {
	return f.CodeComments
}

func (f BasicFunctionDefinition) GetUnitOperations() []UnitOperation {
	return f.UnitOperations
}

func (f BasicFunctionDefinition) GetShouldStub() bool {
	return f.ShouldStub
}
