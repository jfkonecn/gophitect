package unitops

import (
	"fmt"
	"strings"
)

type BasicVariableDefinition struct {
	Name           string
	TypeDefinition *TypeDefinition
}

func (v *BasicVariableDefinition) GetVariableName() string {
	return v.Name
}

func (v *BasicVariableDefinition) GetTypeDefinition() *TypeDefinition {
	return v.TypeDefinition
}

type CollectionVariableDefinition struct {
	Name           string
	TypeDefinition *CollectionTypeDefinition
}

func (v *CollectionVariableDefinition) GetVariableName() string {
	return v.Name
}

func (v *CollectionVariableDefinition) GetTypeDefinition() *TypeDefinition {
	return v.TypeDefinition
}

type PrimitiveTypeDefinition struct {
	primitiveType string
}

func (t *PrimitiveTypeDefinition) GetPrompt() (string, error) {
	panic("PrimitiveTypeDefinition does not have a prompt")
}

func (t *PrimitiveTypeDefinition) GetTypeName() string {
	return fmt.Sprintf("%s", t.primitiveType)
}

func (t *PrimitiveTypeDefinition) IsBuiltin() bool {
	return true
}

func (t *PrimitiveTypeDefinition) GetDependentTypes() []*TypeDefinition {
	return nil
}

type CollectionTypeDefinition struct {
	CollectionType string
	TypeDefinition TypeDefinition
}

func (t *CollectionTypeDefinition) GetPrompt() (string, error) {
	panic("CollectionTypeDefinition does not have a prompt")
}

func (t *CollectionTypeDefinition) GetTypeName() string {
	return fmt.Sprintf("A %s of %s", t.CollectionType, t.TypeDefinition)
}

func (t *CollectionTypeDefinition) IsBuiltin() bool {
	return true
}

func (t *CollectionTypeDefinition) GetDependentTypes() []*TypeDefinition {
	if t.TypeDefinition.IsBuiltin() {
		return nil
	} else {
		return []*TypeDefinition{&t.TypeDefinition}
	}
}

type StructTypeField struct {
	Name           string
	CodeComment    string
	TypeDefinition *TypeDefinition
}

type StructTypeDefinition struct {
	TypeName     string
	CodeComments string
	Fields       []*StructTypeField
}

func (t StructTypeDefinition) GetPrompt() (string, error) {
	var sb strings.Builder

	_, err := fmt.Fprintf(&sb, "# Type Definition For %s\n", t.TypeName)
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintln(&sb, "You are going make a type definition. Here are the specs:")
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintln(&sb, "## Code Comments")
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "Please Add The Following Code Comments\n%s\n", t.CodeComments)
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintln(&sb, "## Fields")
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintln(&sb, "The format will be #. <field-name>:<field-type> -- <code-comments>")
	if err != nil {
		return "", err
	}

	for i, field := range t.Fields {
		if field.TypeDefinition == nil {
			return "", fmt.Errorf("%s has a nil type definition for the field \"%s\"", t.TypeName, field.Name)
		}
		typeDefinition := *field.TypeDefinition

		_, err = fmt.Fprintf(&sb, "%d. %s:%s", i, field.Name, typeDefinition.GetTypeName())
		if err != nil {
			return "", err
		}

		comments := field.CodeComment
		if comments == "" {
			comments = "No comments"
		}

		_, err = fmt.Fprintf(&sb, " -- %s", comments)
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (t StructTypeDefinition) IsBuiltin() bool {
	return false
}

func (t StructTypeDefinition) GetDependentTypes() []*TypeDefinition {
	var definitions []*TypeDefinition
	for _, field := range t.Fields {
		if (*field.TypeDefinition).IsBuiltin() {
			definitions = append(definitions, field.TypeDefinition)
		}
	}
	return definitions
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
	ProgrammingLanguage string
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
