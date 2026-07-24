package unitops

type BasicTestSuite struct {
	CodeComments       string
	FunctionDefinition FunctionDefinition
	TestCases          []TestCase
}

func (t BasicTestSuite) GetCodeComments() string {
	return t.CodeComments
}

func (t BasicTestSuite) GetFunctionDefinition() FunctionDefinition {
	return t.FunctionDefinition
}

func (t BasicTestSuite) GetTestCases() []TestCase {
	return t.TestCases
}
