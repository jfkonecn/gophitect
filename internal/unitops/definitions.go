package unitops

import (
	"fmt"
	"strings"
)

type BasicTypeDefinition struct {
	TypeName            string
	CodeComments        string
	DependentTypes      []*TypeDefinition
	ShouldStub          bool
	ProgrammingLanguage string
}

func (t BasicTypeDefinition) GetPrompt() (string, error) {
	var sb strings.Builder

	_, err := fmt.Fprintln(&sb, "# Type Definition")
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "%s\n", t.TypeName)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

func (t BasicTypeDefinition) GetDependentTypes() []*TypeDefinition {
	return t.DependentTypes
}

type BasicFunctionDefinition struct {
	FunctionName   string
	Prompt         string
	UnitOperations []*UnitOperation
}

func (f BasicFunctionDefinition) GetFunctionName() string {
	return f.FunctionName
}

func (f BasicFunctionDefinition) GetPrompt() string {
	return f.Prompt
}

func (f BasicFunctionDefinition) GetUnitOperations() []*UnitOperation {
	return f.UnitOperations
}

type BasicTestCase struct {
	Name   string
	Prompt string
}

func (t BasicTestCase) GetName() string {
	return t.Name
}

func (t BasicTestCase) GetPrompt() string {
	return t.Prompt
}

type BasicTestSuite struct {
	Prompt             string
	FunctionDefinition *FunctionDefinition
	TestCases          []*TestCase
}

func (t BasicTestSuite) GetPrompt() string {
	return t.Prompt
}

func (t BasicTestSuite) GetFunctionDefinition() *FunctionDefinition {
	return t.FunctionDefinition
}

func (t BasicTestSuite) GetTestCases() []*TestCase {
	return t.TestCases
}

type BasicTestFileDefinition struct {
	Prompt     string
	FilePath   string
	TestSuites []*TestSuite
}

func (t BasicTestFileDefinition) GetPrompt() string {
	return t.Prompt
}

func (t BasicTestFileDefinition) GetFilePath() string {
	return t.FilePath
}

func (t BasicTestFileDefinition) GetTestSuites() []*TestSuite {
	return t.TestSuites
}

type BasicProductionFileDefinition struct {
	Prompt              string
	FilePath            string
	TypeDefinitions     []*TypeDefinition
	FunctionDefinitions []*FunctionDefinition
}

func (p BasicProductionFileDefinition) GetPrompt() string {
	return p.Prompt
}

func (p BasicProductionFileDefinition) GetFilePath() string {
	return p.FilePath
}

func (p BasicProductionFileDefinition) GetTypeDefinitions() []*TypeDefinition {
	return p.TypeDefinitions
}

func (p BasicProductionFileDefinition) GetFunctionDefinitions() []*FunctionDefinition {
	return p.FunctionDefinitions
}
