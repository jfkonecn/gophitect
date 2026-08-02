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
	assertContains(t, prompt, "## Programming Language\nImplement the function in Go.")
	assertContains(t, prompt, "- Add these comments\nReturn a display-safe name.")
	assertContains(t, prompt, "Use these unit operations in order")
	assertContains(t, prompt, "1. Map rawName of type string to trimmedName of type string")
	assertContains(t, prompt, "- Add these comments\nTrim whitespace.")
	assertContains(t, prompt, "2. Map trimmedName of type string to normalizedName of type string")
}

func TestBasicFunctionDefinitionGetPromptUsesProgrammingLanguageInformation(t *testing.T) {
	function := BasicFunctionDefinition{
		FunctionName:                   "NormalizeName",
		ProgrammingLanguageInformation: "Implement the function in TypeScript.",
		PackageNamespace:               "Put this function in the user namespace.",
	}

	prompt, err := function.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "## Programming Language\nImplement the function in TypeScript.")
	assertContains(t, prompt, "## Package/Namespace\nPut this function in the user namespace.")
}

func TestStructTypeDefinitionGetPromptUsesProgrammingLanguageInformation(t *testing.T) {
	typeDefinition := StructTypeDefinition{
		TypeName:                       "User",
		ProgrammingLanguageInformation: "Define the type in Go.",
		PackageNamespace:               "Put this type in package user.",
	}

	prompt, err := typeDefinition.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "## Programming Language\nDefine the type in Go.")
	assertContains(t, prompt, "## Package/Namespace\nPut this type in package user.")
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

func TestDefinitionsGetPackageNamespace(t *testing.T) {
	structType := StructTypeDefinition{PackageNamespace: "Put this type in package user."}
	if structType.GetPackageNamespace() != "Put this type in package user." {
		t.Fatalf("StructTypeDefinition.GetPackageNamespace() = %q, want %q", structType.GetPackageNamespace(), "Put this type in package user.")
	}

	function := BasicFunctionDefinition{PackageNamespace: "Put this function in package user."}
	if function.GetPackageNamespace() != "Put this function in package user." {
		t.Fatalf("BasicFunctionDefinition.GetPackageNamespace() = %q, want %q", function.GetPackageNamespace(), "Put this function in package user.")
	}

	testSuite := BasicTestSuite{PackageNamespace: "Put these tests in package user_test."}
	if testSuite.GetPackageNamespace() != "Put these tests in package user_test." {
		t.Fatalf("BasicTestSuite.GetPackageNamespace() = %q, want %q", testSuite.GetPackageNamespace(), "Put these tests in package user_test.")
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

func TestBasicTestCaseGetPrompt(t *testing.T) {
	testCase := BasicTestCase{
		Name:        "valid user",
		Description: "Returns a normalized user when input is valid.",
	}

	prompt, err := testCase.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "test case valid user")
	assertContains(t, prompt, "Use this test case description\nReturns a normalized user when input is valid.")
}

func TestBasicTestCaseGetPromptErrors(t *testing.T) {
	_, err := (BasicTestCase{}).GetPrompt()
	if err == nil {
		t.Fatal("GetPrompt() returned nil error")
	}

	if err.Error() != "test case's name is empty" {
		t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), "test case's name is empty")
	}
}

func TestBasicTestSuiteGetPrompt(t *testing.T) {
	function := FunctionDefinition(BasicFunctionDefinition{
		FunctionName: "NormalizeUser",
		FilePath:     "internal/user.go",
	})
	validUserCase := TestCase(BasicTestCase{
		Name:        "valid user",
		Description: "Returns a normalized user when input is valid.",
	})
	invalidUserCase := TestCase(BasicTestCase{
		Name:        "invalid user",
		Description: "Returns a validation error when input is invalid.",
	})
	testSuite := BasicTestSuite{
		Name:               "NormalizeUser",
		FilePath:           "internal/user_test.go",
		FunctionDefinition: &function,
		TestCases:          []*TestCase{&validUserCase, &invalidUserCase},
	}

	prompt, err := testSuite.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "# Test Suite For NormalizeUser")
	assertContains(t, prompt, "Create the test suite in internal/user_test.go")
	assertContains(t, prompt, "## Programming Language\nImplement the function in Go.")
	assertContains(t, prompt, "Test function NormalizeUser from internal/user.go")
	assertContains(t, prompt, "Use these test cases")
	assertContains(t, prompt, "1. test case valid user")
	assertContains(t, prompt, "Returns a normalized user when input is valid.")
	assertContains(t, prompt, "2. test case invalid user")
	assertContains(t, prompt, "Returns a validation error when input is invalid.")
}

func TestBasicTestSuiteGetPromptUsesProgrammingLanguageInformation(t *testing.T) {
	function := FunctionDefinition(BasicFunctionDefinition{FunctionName: "NormalizeUser"})
	testSuite := BasicTestSuite{
		Name:                           "NormalizeUser",
		ProgrammingLanguageInformation: "Write the tests in Go using the testing package.",
		PackageNamespace:               "Put these tests in package user_test.",
		FunctionDefinition:             &function,
	}

	prompt, err := testSuite.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "## Programming Language\nWrite the tests in Go using the testing package.")
	assertContains(t, prompt, "## Package/Namespace\nPut these tests in package user_test.")
}

func TestBasicTestSuiteGetPromptErrors(t *testing.T) {
	function := FunctionDefinition(BasicFunctionDefinition{FunctionName: "NormalizeUser"})
	testCase := TestCase(BasicTestCase{Name: "valid user"})
	unnamedTestCase := TestCase(BasicTestCase{})

	tests := []struct {
		name      string
		testSuite BasicTestSuite
		want      string
	}{
		{
			name:      "empty name",
			testSuite: BasicTestSuite{},
			want:      "test suite's name is empty",
		},
		{
			name: "nil function definition",
			testSuite: BasicTestSuite{
				Name: "NormalizeUser",
			},
			want: "test suite's function definition is nil",
		},
		{
			name: "nil test case",
			testSuite: BasicTestSuite{
				Name:               "NormalizeUser",
				FunctionDefinition: &function,
				TestCases:          []*TestCase{nil},
			},
			want: "test case in test suite is nil",
		},
		{
			name: "test case prompt error",
			testSuite: BasicTestSuite{
				Name:               "NormalizeUser",
				FunctionDefinition: &function,
				TestCases:          []*TestCase{&testCase, &unnamedTestCase},
			},
			want: "test case's name is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.testSuite.GetPrompt()
			if err == nil {
				t.Fatal("GetPrompt() returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), tt.want)
			}
		})
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
