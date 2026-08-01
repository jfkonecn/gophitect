package unitops

import (
	"fmt"
	"slices"
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
	TypeDefinition *TypeDefinition
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

func (t *PrimitiveTypeDefinition) GetID() string {
	return defaultDefinitionID(t.GetFilePath(), t.GetTypeName())
}

func (t *PrimitiveTypeDefinition) GetFilePath() string {
	return ""
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
	return fmt.Sprintf("A %s of %s", t.CollectionType, t.TypeDefinition.GetTypeName())
}

func (t *CollectionTypeDefinition) GetID() string {
	return defaultDefinitionID(t.GetFilePath(), t.GetTypeName())
}

func (t *CollectionTypeDefinition) GetFilePath() string {
	return t.TypeDefinition.GetFilePath()
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
	ID           string
	TypeName     string
	FilePath     string
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

func (t StructTypeDefinition) GetTypeName() string {
	return t.TypeName
}

func (t StructTypeDefinition) GetID() string {
	if t.ID != "" {
		return t.ID
	}

	return defaultDefinitionID(t.FilePath, t.TypeName)
}

func (t StructTypeDefinition) GetFilePath() string {
	return t.FilePath
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
	ID             string
	FunctionName   string
	FilePath       string
	CodeComments   string
	UnitOperations []*UnitOperation
}

func (f BasicFunctionDefinition) GetFunctionName() string {
	return f.FunctionName
}

func (f BasicFunctionDefinition) GetID() string {
	if f.ID != "" {
		return f.ID
	}

	return defaultDefinitionID(f.FilePath, f.FunctionName)
}

func (f BasicFunctionDefinition) GetFilePath() string {
	return f.FilePath
}

func (f BasicFunctionDefinition) GetPrompt() (string, error) {
	var sb strings.Builder

	if f.FunctionName == "" {
		return "", fmt.Errorf("function's name is empty")
	}

	_, err := fmt.Fprintf(&sb, "implement function %s\n", f.FunctionName)
	if err != nil {
		return "", err
	}

	if f.CodeComments != "" {
		_, err := fmt.Fprintf(&sb, "- Add these comments\n%s\n", f.CodeComments)
		if err != nil {
			return "", err
		}
	}

	if len(f.UnitOperations) > 0 {
		_, err := fmt.Fprintln(&sb, "Use these unit operations in order")
		if err != nil {
			return "", err
		}
	}

	for i, unitOperation := range f.UnitOperations {
		if unitOperation == nil {
			return "", fmt.Errorf("unit operation in function is nil")
		}

		unitOperationPrompt, err := (*unitOperation).GetPrompt()
		if err != nil {
			return "", err
		}

		_, err = fmt.Fprintf(&sb, "%d. %s", i+1, unitOperationPrompt)
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (f BasicFunctionDefinition) GetInputs() []*VariableDefinition {
	var inputs []*VariableDefinition
	seenInputs := map[string]bool{}
	produced := map[string]bool{}

	for _, unitOperation := range f.UnitOperations {
		if unitOperation == nil {
			continue
		}

		for _, input := range (*unitOperation).GetInputTypes() {
			if input == nil {
				continue
			}

			key := variableDefinitionKey(input)
			if produced[key] || seenInputs[key] {
				continue
			}

			inputs = append(inputs, input)
			seenInputs[key] = true
		}

		for _, output := range (*unitOperation).GetOutputTypes() {
			if output == nil {
				continue
			}

			produced[variableDefinitionKey(output)] = true
		}
	}

	return inputs
}

func (f BasicFunctionDefinition) GetOutput() *TypeDefinition {
	for _, unitOperation := range slices.Backward(f.UnitOperations) {

		if unitOperation == nil {
			continue
		}

		outputs := (*unitOperation).GetOutputTypes()
		for _, output := range slices.Backward(outputs) {
			if output == nil {
				continue
			}

			return (*output).GetTypeDefinition()
		}
	}

	return nil
}

func (f BasicFunctionDefinition) GetUnitOperations() []*UnitOperation {
	return f.UnitOperations
}

func variableDefinitionKey(variableDefinition *VariableDefinition) string {
	if variableDefinition == nil {
		return ""
	}

	variable := *variableDefinition
	if variable.GetTypeDefinition() == nil {
		return variable.GetVariableName()
	}

	typeDefinition := *variable.GetTypeDefinition()
	return fmt.Sprintf("%s:%s", variable.GetVariableName(), typeDefinition.GetTypeName())
}

type BasicTestCase struct {
	Name        string
	Description string
}

func (t BasicTestCase) GetName() string {
	return t.Name
}

func (t BasicTestCase) GetPrompt() (string, error) {
	var sb strings.Builder

	if t.Name == "" {
		return "", fmt.Errorf("test case's name is empty")
	}

	_, err := fmt.Fprintf(&sb, "test case %s\n", t.Name)
	if err != nil {
		return "", err
	}

	if t.Description != "" {
		_, err = fmt.Fprintf(&sb, "Use this test case description\n%s\n", t.Description)
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

type BasicTestSuite struct {
	ID                 string
	Name               string
	FilePath           string
	FunctionDefinition *FunctionDefinition
	TestCases          []*TestCase
}

func (t BasicTestSuite) GetID() string {
	if t.ID != "" {
		return t.ID
	}

	return defaultDefinitionID(t.FilePath, t.Name)
}

func (t BasicTestSuite) GetPrompt() (string, error) {
	var sb strings.Builder

	if t.Name == "" {
		return "", fmt.Errorf("test suite's name is empty")
	}

	if t.FunctionDefinition == nil {
		return "", fmt.Errorf("test suite's function definition is nil")
	}
	functionDefinition := *t.FunctionDefinition

	_, err := fmt.Fprintf(&sb, "# Test Suite For %s\n", t.Name)
	if err != nil {
		return "", err
	}

	if t.FilePath != "" {
		_, err = fmt.Fprintf(&sb, "Create the test suite in %s\n", t.FilePath)
		if err != nil {
			return "", err
		}
	}

	_, err = fmt.Fprintf(&sb, "Test function %s", functionDefinition.GetFunctionName())
	if err != nil {
		return "", err
	}

	if functionDefinition.GetFilePath() != "" {
		_, err = fmt.Fprintf(&sb, " from %s", functionDefinition.GetFilePath())
		if err != nil {
			return "", err
		}
	}

	_, err = fmt.Fprintln(&sb)
	if err != nil {
		return "", err
	}

	if len(t.TestCases) > 0 {
		_, err = fmt.Fprintln(&sb, "Use these test cases")
		if err != nil {
			return "", err
		}
	}

	for i, testCase := range t.TestCases {
		if testCase == nil {
			return "", fmt.Errorf("test case in test suite is nil")
		}

		testCasePrompt, err := (*testCase).GetPrompt()
		if err != nil {
			return "", err
		}

		_, err = fmt.Fprintf(&sb, "%d. %s", i+1, testCasePrompt)
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (t BasicTestSuite) GetFilePath() string {
	return t.FilePath
}

func (t BasicTestSuite) GetFunctionDefinition() *FunctionDefinition {
	return t.FunctionDefinition
}

func (t BasicTestSuite) GetTestCases() []*TestCase {
	return t.TestCases
}

func defaultDefinitionID(filePath, name string) string {
	return filePath + name
}
