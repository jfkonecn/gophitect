package unitops

import (
	"errors"
	"testing"
)

type failingUnitOperation struct{}

func (f failingUnitOperation) GetInputTypes() []*VariableDefinition {
	return nil
}

func (f failingUnitOperation) GetOutputTypes() []*VariableDefinition {
	return nil
}

func (f failingUnitOperation) GetPrompt() (string, error) {
	return "", errors.New("unit operation failed")
}

func (f failingUnitOperation) GetFunctionCalls() []*FunctionDefinition {
	return nil
}

func TestBasicFunctionDefinitionGetPrompt(t *testing.T) {
	firstOperation := UnitOperation(Map{
		Input:        primitiveVariable("rawName", "string"),
		Output:       primitiveVariable("trimmedName", "string"),
		CodeComments: "Trim whitespace.",
	})
	secondOperation := UnitOperation(Map{
		Input:  primitiveVariable("trimmedName", "string"),
		Output: primitiveVariable("normalizedName", "string"),
	})
	function := BasicFunctionDefinition{
		FunctionName:   "NormalizeName",
		CodeComments:   "Return a display-safe name.",
		UnitOperations: []*UnitOperation{&firstOperation, &secondOperation},
	}

	prompt, err := function.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "implement function NormalizeName")
	assertContains(t, prompt, "- Add these comments\nReturn a display-safe name.")
	assertContains(t, prompt, "Use these unit operations in order")
	assertContains(t, prompt, "1. Map rawName of type string to trimmedName of type string")
	assertContains(t, prompt, "- Add these comments\nTrim whitespace.")
	assertContains(t, prompt, "2. Map trimmedName of type string to normalizedName of type string")
}

func TestBasicFunctionDefinitionGetPromptErrors(t *testing.T) {
	failingOperation := UnitOperation(failingUnitOperation{})

	tests := []struct {
		name     string
		function BasicFunctionDefinition
		want     string
	}{
		{
			name:     "empty name",
			function: BasicFunctionDefinition{},
			want:     "function's name is empty",
		},
		{
			name: "nil unit operation",
			function: BasicFunctionDefinition{
				FunctionName:   "DoWork",
				UnitOperations: []*UnitOperation{nil},
			},
			want: "unit operation in function is nil",
		},
		{
			name: "unit operation prompt error",
			function: BasicFunctionDefinition{
				FunctionName:   "DoWork",
				UnitOperations: []*UnitOperation{&failingOperation},
			},
			want: "unit operation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.function.GetPrompt()
			if err == nil {
				t.Fatal("GetPrompt() returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), tt.want)
			}
		})
	}
}

func TestBasicFunctionDefinitionGetInputs(t *testing.T) {
	firstOperation := UnitOperation(Map{
		Input:  primitiveVariable("rawName", "string"),
		Output: primitiveVariable("trimmedName", "string"),
	})
	secondOperation := UnitOperation(Map{
		Input:  primitiveVariable("trimmedName", "string"),
		Output: primitiveVariable("normalizedName", "string"),
	})
	thirdOperation := UnitOperation(Map{
		Input:  primitiveVariable("locale", "string"),
		Output: primitiveVariable("localizedName", "string"),
	})
	function := BasicFunctionDefinition{
		UnitOperations: []*UnitOperation{&firstOperation, &secondOperation, &thirdOperation},
	}

	inputs := function.GetInputs()
	if len(inputs) != 2 {
		t.Fatalf("GetInputs() returned %d inputs, want 2", len(inputs))
	}

	assertVariable(t, inputs[0], "rawName", "string")
	assertVariable(t, inputs[1], "locale", "string")
}

func TestBasicFunctionDefinitionGetOutput(t *testing.T) {
	firstOperation := UnitOperation(Map{
		Input:  primitiveVariable("rawName", "string"),
		Output: primitiveVariable("trimmedName", "string"),
	})
	secondOperation := UnitOperation(Map{
		Input:  primitiveVariable("trimmedName", "string"),
		Output: primitiveVariable("normalizedName", "string"),
	})
	function := BasicFunctionDefinition{
		UnitOperations: []*UnitOperation{&firstOperation, &secondOperation},
	}

	output := function.GetOutput()
	if output == nil {
		t.Fatal("GetOutput() returned nil")
	}

	if (*output).GetTypeName() != "string" {
		t.Fatalf("GetOutput() type = %q, want %q", (*output).GetTypeName(), "string")
	}
}

func TestBasicFunctionDefinitionImplementsFunctionDefinition(t *testing.T) {
	var _ FunctionDefinition = BasicFunctionDefinition{}
}

func TestBasicDefinitionsImplementInterfaces(t *testing.T) {
	var _ TypeDefinition = StructTypeDefinition{}
	var _ TestCase = BasicTestCase{}
	var _ TestSuite = BasicTestSuite{}
	var _ TestFileDefinition = BasicTestFileDefinition{}
	var _ ProductionFileDefinition = BasicProductionFileDefinition{}
}

func TestDefinitionsGetFilePath(t *testing.T) {
	structType := StructTypeDefinition{
		TypeName: "User",
		FilePath: "internal/user.go",
	}
	if structType.GetFilePath() != "internal/user.go" {
		t.Fatalf("StructTypeDefinition.GetFilePath() = %q, want %q", structType.GetFilePath(), "internal/user.go")
	}

	function := BasicFunctionDefinition{
		FunctionName: "NormalizeUser",
		FilePath:     "internal/user.go",
	}
	if function.GetFilePath() != "internal/user.go" {
		t.Fatalf("BasicFunctionDefinition.GetFilePath() = %q, want %q", function.GetFilePath(), "internal/user.go")
	}

	testSuite := BasicTestSuite{
		FilePath: "internal/user_test.go",
	}
	if testSuite.GetFilePath() != "internal/user_test.go" {
		t.Fatalf("BasicTestSuite.GetFilePath() = %q, want %q", testSuite.GetFilePath(), "internal/user_test.go")
	}
}

func TestDefinitionsGetIDDefaultsToFilePathAndName(t *testing.T) {
	structType := StructTypeDefinition{
		TypeName: "User",
		FilePath: "internal/user.go",
	}
	if structType.GetID() != "internal/user.goUser" {
		t.Fatalf("StructTypeDefinition.GetID() = %q, want %q", structType.GetID(), "internal/user.goUser")
	}

	function := BasicFunctionDefinition{
		FunctionName: "NormalizeUser",
		FilePath:     "internal/user.go",
	}
	if function.GetID() != "internal/user.goNormalizeUser" {
		t.Fatalf("BasicFunctionDefinition.GetID() = %q, want %q", function.GetID(), "internal/user.goNormalizeUser")
	}

	testSuite := BasicTestSuite{
		Name:     "NormalizeUser",
		FilePath: "internal/user_test.go",
	}
	if testSuite.GetID() != "internal/user_test.goNormalizeUser" {
		t.Fatalf("BasicTestSuite.GetID() = %q, want %q", testSuite.GetID(), "internal/user_test.goNormalizeUser")
	}
}

func TestDefinitionsGetIDUsesExplicitID(t *testing.T) {
	structType := StructTypeDefinition{
		ID:       "type:user",
		TypeName: "User",
		FilePath: "internal/user.go",
	}
	if structType.GetID() != "type:user" {
		t.Fatalf("StructTypeDefinition.GetID() = %q, want %q", structType.GetID(), "type:user")
	}

	function := BasicFunctionDefinition{
		ID:           "function:normalize-user",
		FunctionName: "NormalizeUser",
		FilePath:     "internal/user.go",
	}
	if function.GetID() != "function:normalize-user" {
		t.Fatalf("BasicFunctionDefinition.GetID() = %q, want %q", function.GetID(), "function:normalize-user")
	}

	testSuite := BasicTestSuite{
		ID:       "suite:normalize-user",
		Name:     "NormalizeUser",
		FilePath: "internal/user_test.go",
	}
	if testSuite.GetID() != "suite:normalize-user" {
		t.Fatalf("BasicTestSuite.GetID() = %q, want %q", testSuite.GetID(), "suite:normalize-user")
	}
}

func assertVariable(t *testing.T, variableDefinition *VariableDefinition, name, typeName string) {
	t.Helper()

	if variableDefinition == nil {
		t.Fatal("variable definition is nil")
	}

	variable := *variableDefinition
	if variable.GetVariableName() != name {
		t.Fatalf("variable name = %q, want %q", variable.GetVariableName(), name)
	}

	if variable.GetTypeDefinition() == nil {
		t.Fatal("variable type definition is nil")
	}

	typeDefinition := *variable.GetTypeDefinition()
	if typeDefinition.GetTypeName() != typeName {
		t.Fatalf("variable type = %q, want %q", typeDefinition.GetTypeName(), typeName)
	}
}
