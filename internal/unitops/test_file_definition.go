package unitops

type BasicTestFileDefinition struct {
	DefinitionPrompt string
	FilePath         string
	CodeComments     string
	TestSuites       []TestSuite
}

func (t BasicTestFileDefinition) FileDefinitionPrompt() string {
	return t.DefinitionPrompt
}

func (t BasicTestFileDefinition) GetFilePath() string {
	return t.FilePath
}

func (t BasicTestFileDefinition) GetCodeComments() string {
	return t.CodeComments
}

func (t BasicTestFileDefinition) GetTestSuites() []TestSuite {
	return t.TestSuites
}
