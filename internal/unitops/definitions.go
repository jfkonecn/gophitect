package unitops

type BasicTypeDefinition struct {
	TypeName       string
	CodeComments   string
	DependentTypes []*TypeDefinition
	FilePath       string
	ShouldStub     bool
}

func (t BasicTypeDefinition) GetTypeName() string {
	return t.TypeName
}

func (t BasicTypeDefinition) GetCodeComments() string {
	return t.CodeComments
}

func (t BasicTypeDefinition) GetDependentTypes() []*TypeDefinition {
	return t.DependentTypes
}

func (t BasicTypeDefinition) GetFilePath() string {
	return t.FilePath
}

func (t BasicTypeDefinition) GetShouldStub() bool {
	return t.ShouldStub
}

type BasicFunctionDefinition struct {
	FunctionName   string
	CodeComments   string
	UnitOperations []*UnitOperation
	ShouldStub     bool
}

func (f BasicFunctionDefinition) GetFunctionName() string {
	return f.FunctionName
}

func (f BasicFunctionDefinition) GetCodeComments() string {
	return f.CodeComments
}

func (f BasicFunctionDefinition) GetUnitOperations() []*UnitOperation {
	return f.UnitOperations
}

func (f BasicFunctionDefinition) GetShouldStub() bool {
	return f.ShouldStub
}

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

type BasicTestSuite struct {
	CodeComments       string
	FunctionDefinition *FunctionDefinition
	TestCases          []*TestCase
}

func (t BasicTestSuite) GetCodeComments() string {
	return t.CodeComments
}

func (t BasicTestSuite) GetFunctionDefinition() *FunctionDefinition {
	return t.FunctionDefinition
}

func (t BasicTestSuite) GetTestCases() []*TestCase {
	return t.TestCases
}

type BasicTestFileDefinition struct {
	DefinitionPrompt string
	FilePath         string
	CodeComments     string
	TestSuites       []*TestSuite
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

func (t BasicTestFileDefinition) GetTestSuites() []*TestSuite {
	return t.TestSuites
}

type BasicProductionFileDefinition struct {
	DefinitionPrompt    string
	FilePath            string
	CodeComments        string
	TypeDefinitions     []*TypeDefinition
	FunctionDefinitions []*FunctionDefinition
}

func (p BasicProductionFileDefinition) FileDefinitionPrompt() string {
	return p.DefinitionPrompt
}

func (p BasicProductionFileDefinition) GetFilePath() string {
	return p.FilePath
}

func (p BasicProductionFileDefinition) GetCodeComments() string {
	return p.CodeComments
}

func (p BasicProductionFileDefinition) GetTypeDefinitions() []*TypeDefinition {
	return p.TypeDefinitions
}

func (p BasicProductionFileDefinition) GetFunctionDefinitions() []*FunctionDefinition {
	return p.FunctionDefinitions
}
