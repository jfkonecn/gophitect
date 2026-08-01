package unitops

import (
	"strings"
	"testing"
)

func TestImplementationRunnerImplementsInDependencyOrder(t *testing.T) {
	baseType := TypeDefinition(StructTypeDefinition{
		ID:       "type:base",
		TypeName: "Base",
		FilePath: "internal/types.go",
	})
	dependentType := TypeDefinition(StructTypeDefinition{
		ID:       "type:dependent",
		TypeName: "Dependent",
		FilePath: "internal/types.go",
		Fields: []*StructTypeField{
			{
				Name:           "Base",
				TypeDefinition: &baseType,
			},
		},
	})
	input := VariableDefinition(&BasicVariableDefinition{
		Name:           "input",
		TypeDefinition: &dependentType,
	})
	output := VariableDefinition(&BasicVariableDefinition{
		Name:           "output",
		TypeDefinition: &baseType,
	})

	helperFunction := FunctionDefinition(BasicFunctionDefinition{
		ID:           "function:helper",
		FunctionName: "Helper",
		FilePath:     "internal/helper.go",
	})
	operation := UnitOperation(Map{
		Input:         &input,
		Output:        &output,
		FunctionCalls: []*FunctionDefinition{&helperFunction},
	})
	mainFunction := FunctionDefinition(BasicFunctionDefinition{
		ID:             "function:main",
		FunctionName:   "Main",
		FilePath:       "internal/main.go",
		UnitOperations: []*UnitOperation{&operation},
	})
	testSuite := TestSuite(BasicTestSuite{
		ID:                 "test:main",
		Name:               "Main",
		FilePath:           "internal/main_test.go",
		FunctionDefinition: &mainFunction,
	})

	var prompts []string
	var writes []string
	runner := ImplementationRunner{
		ReadFile: func(filePath string) (string, error) {
			return "existing content for " + filePath, nil
		},
		PromptLLM: func(prompt string) (string, error) {
			prompts = append(prompts, prompt)
			return prompt, nil
		},
		WriteFile: func(filePath string, content string) error {
			writes = append(writes, filePath+":"+promptKind(content))
			return nil
		},
	}

	err := runner.Implement(
		[]*FunctionDefinition{&mainFunction},
		[]*TestSuite{&testSuite},
		nil,
		nil,
		nil,
	)
	if err != nil {
		t.Fatalf("Implement() returned error: %v", err)
	}

	wantWrites := []string{
		"internal/types.go:type:Base",
		"internal/types.go:type:Dependent",
		"internal/helper.go:function:Helper",
		"internal/main.go:function:Main",
		"internal/main_test.go:test:Main",
	}
	assertStringSlice(t, writes, wantWrites)

	for _, prompt := range prompts {
		assertContains(t, prompt, "Existing file content:")
		assertContains(t, prompt, completeFileContentsPrompt)
	}
}

func TestImplementationRunnerSkipsCompletedIDs(t *testing.T) {
	userType := TypeDefinition(StructTypeDefinition{
		ID:       "type:user",
		TypeName: "User",
		FilePath: "internal/user.go",
	})
	input := VariableDefinition(&BasicVariableDefinition{
		Name:           "user",
		TypeDefinition: &userType,
	})
	operation := UnitOperation(Map{
		Input:  &input,
		Output: &input,
	})
	function := FunctionDefinition(BasicFunctionDefinition{
		ID:             "function:normalize-user",
		FunctionName:   "NormalizeUser",
		FilePath:       "internal/user.go",
		UnitOperations: []*UnitOperation{&operation},
	})
	testSuite := TestSuite(BasicTestSuite{
		ID:                 "test:normalize-user",
		Name:               "NormalizeUser",
		FilePath:           "internal/user_test.go",
		FunctionDefinition: &function,
	})

	var writes []string
	runner := ImplementationRunner{
		ReadFile: func(filePath string) (string, error) {
			return "", nil
		},
		PromptLLM: func(prompt string) (string, error) {
			return prompt, nil
		},
		WriteFile: func(filePath string, content string) error {
			writes = append(writes, filePath+":"+promptKind(content))
			return nil
		},
	}

	err := runner.Implement(
		[]*FunctionDefinition{&function},
		[]*TestSuite{&testSuite},
		[]string{"type:user"},
		[]string{"function:normalize-user"},
		[]string{"test:normalize-user"},
	)
	if err != nil {
		t.Fatalf("Implement() returned error: %v", err)
	}

	if len(writes) != 0 {
		t.Fatalf("writes = %v, want none", writes)
	}
}

func TestImplementationRunnerRequiresCallbacks(t *testing.T) {
	err := (ImplementationRunner{}).Implement(nil, nil, nil, nil, nil)
	if err == nil {
		t.Fatal("Implement() returned nil error")
	}

	if err.Error() != "prompt LLM function is nil" {
		t.Fatalf("Implement() error = %q, want %q", err.Error(), "prompt LLM function is nil")
	}
}

func promptKind(prompt string) string {
	switch {
	case strings.Contains(prompt, "# Type Definition For Base"):
		return "type:Base"
	case strings.Contains(prompt, "# Type Definition For Dependent"):
		return "type:Dependent"
	case strings.Contains(prompt, "implement function Helper"):
		return "function:Helper"
	case strings.Contains(prompt, "implement function Main"):
		return "function:Main"
	case strings.Contains(prompt, "# Test Suite For Main"):
		return "test:Main"
	case strings.Contains(prompt, "# Type Definition For User"):
		return "type:User"
	case strings.Contains(prompt, "implement function NormalizeUser"):
		return "function:NormalizeUser"
	case strings.Contains(prompt, "# Test Suite For NormalizeUser"):
		return "test:NormalizeUser"
	default:
		return "unknown"
	}
}

func assertStringSlice(t *testing.T, got, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("got %d values %v, want %d values %v", len(got), got, len(want), want)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %q, want %q; all got %v", i, got[i], want[i], got)
		}
	}
}
