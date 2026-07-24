package unitops

type BasicTestCase struct {
	Name             string
	DefinitionPrompt string
	CodeComments     string
	ShouldStub       bool
}

func (t BasicTestCase) GetName() string {
	return t.Name
}

func (t BasicTestCase) GetDefinitionPrompt() string {
	return t.DefinitionPrompt
}

func (t BasicTestCase) GetCodeComments() string {
	return t.CodeComments
}

func (t BasicTestCase) GetShouldStub() bool {
	return t.ShouldStub
}
