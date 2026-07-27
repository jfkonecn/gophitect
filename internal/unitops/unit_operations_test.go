package unitops

import (
	"strings"
	"testing"
)

type testFunctionDefinition struct {
	name string
}

func (f testFunctionDefinition) GetFunctionName() string {
	return f.name
}

func (f testFunctionDefinition) GetPrompt() (string, error) {
	return "", nil
}

func (f testFunctionDefinition) GetInputs() []*VariableDefinition {
	return nil
}

func (f testFunctionDefinition) GetOutput() *TypeDefinition {
	return nil
}

func (f testFunctionDefinition) GetUnitOperations() []*UnitOperation {
	return nil
}

func primitiveVariable(name, primitiveType string) *VariableDefinition {
	typeDefinition := TypeDefinition(&PrimitiveTypeDefinition{primitiveType: primitiveType})
	variable := VariableDefinition(&BasicVariableDefinition{
		Name:           name,
		TypeDefinition: &typeDefinition,
	})

	return &variable
}

func collectionVariable(name, collectionType, primitiveType string) *CollectionVariableDefinition {
	typeDefinition := TypeDefinition(&PrimitiveTypeDefinition{primitiveType: primitiveType})
	collectionDefinition := TypeDefinition(&CollectionTypeDefinition{
		CollectionType: collectionType,
		TypeDefinition: typeDefinition,
	})

	return &CollectionVariableDefinition{
		Name:           name,
		TypeDefinition: &collectionDefinition,
	}
}

func functionCall(name string) *FunctionDefinition {
	function := FunctionDefinition(testFunctionDefinition{name: name})
	return &function
}

func TestMapGetPrompt(t *testing.T) {
	operation := Map{
		Input:         primitiveVariable("rawName", "string"),
		Output:        primitiveVariable("normalizedName", "string"),
		CodeComments:  "Trim whitespace before returning.",
		FunctionCalls: []*FunctionDefinition{functionCall("normalizeString")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "Map rawName of type string to normalizedName of type string")
	assertContains(t, prompt, "- Add these comments\nTrim whitespace before returning.")
	assertContains(t, prompt, "Make sure you use these functions in the map logic")
	assertContains(t, prompt, "- normalizeString")
}

func TestMapGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation Map
		want      string
	}{
		{
			name:      "nil input",
			operation: Map{Output: primitiveVariable("output", "string")},
			want:      "map's input is nil",
		},
		{
			name: "nil input type definition",
			operation: Map{
				Input:  variableWithoutType("input"),
				Output: primitiveVariable("output", "string"),
			},
			want: "map's input type definition is nil",
		},
		{
			name:      "nil output",
			operation: Map{Input: primitiveVariable("input", "string")},
			want:      "map's output is nil",
		},
		{
			name: "nil output type definition",
			operation: Map{
				Input:  primitiveVariable("input", "string"),
				Output: variableWithoutType("output"),
			},
			want: "map's output type definition is nil",
		},
		{
			name: "nil function call",
			operation: Map{
				Input:         primitiveVariable("input", "string"),
				Output:        primitiveVariable("output", "string"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in map is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.operation.GetPrompt()
			if err == nil {
				t.Fatal("GetPrompt() returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), tt.want)
			}
		})
	}
}

func TestFilterGetPrompt(t *testing.T) {
	operation := Filter{
		Input:         collectionVariable("users", "slice", "User"),
		Output:        collectionVariable("activeUsers", "slice", "User"),
		FilterLogic:   "Keep users where Active is true.",
		CodeComments:  "Preserve original order.",
		FunctionCalls: []*FunctionDefinition{functionCall("isActiveUser")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "filter users of type A slice of User to activeUsers of type A slice of User")
	assertContains(t, prompt, "Do this logic to filter\nKeep users where Active is true.")
	assertContains(t, prompt, "- Add these comments\nPreserve original order.")
	assertContains(t, prompt, "Make sure you use these functions in the filter logic")
	assertContains(t, prompt, "- isActiveUser")
}

func TestFilterGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation Filter
		want      string
	}{
		{
			name:      "nil input",
			operation: Filter{Output: collectionVariable("output", "slice", "string")},
			want:      "filter's input is nil",
		},
		{
			name: "nil input type definition",
			operation: Filter{
				Input:  collectionVariableWithoutType("input"),
				Output: collectionVariable("output", "slice", "string"),
			},
			want: "filter's input type definition is nil",
		},
		{
			name:      "nil output",
			operation: Filter{Input: collectionVariable("input", "slice", "string")},
			want:      "filter's output is nil",
		},
		{
			name: "nil output type definition",
			operation: Filter{
				Input:  collectionVariable("input", "slice", "string"),
				Output: collectionVariableWithoutType("output"),
			},
			want: "filter's output type definition is nil",
		},
		{
			name: "nil function call",
			operation: Filter{
				Input:         collectionVariable("input", "slice", "string"),
				Output:        collectionVariable("output", "slice", "string"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in filter is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.operation.GetPrompt()
			if err == nil {
				t.Fatal("GetPrompt() returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), tt.want)
			}
		})
	}
}

func TestSortGetPrompt(t *testing.T) {
	operation := Sort{
		Input:         primitiveVariable("users", "[]User"),
		SortLogic:     "Sort users by LastName, then FirstName.",
		CodeComments:  "Use a stable sort.",
		FunctionCalls: []*FunctionDefinition{functionCall("compareUsersByName")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "sort users of type []User")
	assertContains(t, prompt, "Do this logic to sort\nSort users by LastName, then FirstName.")
	assertContains(t, prompt, "- Add these comments\nUse a stable sort.")
	assertContains(t, prompt, "Make sure you use these functions in the sort logic")
	assertContains(t, prompt, "- compareUsersByName")
}

func TestSortGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation Sort
		want      string
	}{
		{
			name:      "nil input",
			operation: Sort{},
			want:      "sort's input is nil",
		},
		{
			name: "nil input type definition",
			operation: Sort{
				Input: variableWithoutType("input"),
			},
			want: "sort's input type definition is nil",
		},
		{
			name: "nil function call",
			operation: Sort{
				Input:         primitiveVariable("input", "[]string"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in sort is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.operation.GetPrompt()
			if err == nil {
				t.Fatal("GetPrompt() returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), tt.want)
			}
		})
	}
}

func TestDistributionGetPrompt(t *testing.T) {
	operation := Distribution{
		Input: primitiveVariable("payment", "Payment"),
		Conditions: []DistributionCondition{
			{
				Condition:     "payment.Amount > 1000",
				Output:        primitiveVariable("manualReviewQueue", "ReviewQueueItem"),
				FunctionCalls: []*FunctionDefinition{functionCall("requiresManualReview")},
			},
			{
				Condition: "payment.Amount <= 1000",
				Output:    primitiveVariable("autoApprovalQueue", "ApprovalQueueItem"),
			},
		},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "distribute payment of type Payment based on these conditions")
	assertContains(t, prompt, "Evaluate each condition and route the input to the matching output.")
	assertContains(t, prompt, "1. When payment.Amount > 1000")
	assertContains(t, prompt, "- Route to manualReviewQueue of type ReviewQueueItem")
	assertContains(t, prompt, "- Make sure you use these functions for this condition")
	assertContains(t, prompt, "  - requiresManualReview")
	assertContains(t, prompt, "2. When payment.Amount <= 1000")
	assertContains(t, prompt, "- Route to autoApprovalQueue of type ApprovalQueueItem")
}

func TestDistributionGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation Distribution
		want      string
	}{
		{
			name:      "nil input",
			operation: Distribution{},
			want:      "distribution's input is nil",
		},
		{
			name: "nil input type definition",
			operation: Distribution{
				Input: variableWithoutType("input"),
				Conditions: []DistributionCondition{
					{
						Condition: "input.Valid",
						Output:    primitiveVariable("output", "string"),
					},
				},
			},
			want: "distribution's input type definition is nil",
		},
		{
			name: "empty conditions",
			operation: Distribution{
				Input: primitiveVariable("input", "string"),
			},
			want: "distribution's conditions are empty",
		},
		{
			name: "empty condition",
			operation: Distribution{
				Input: primitiveVariable("input", "string"),
				Conditions: []DistributionCondition{
					{
						Output: primitiveVariable("output", "string"),
					},
				},
			},
			want: "distribution condition is empty",
		},
		{
			name: "nil output",
			operation: Distribution{
				Input: primitiveVariable("input", "string"),
				Conditions: []DistributionCondition{
					{
						Condition: "input != \"\"",
					},
				},
			},
			want: "distribution condition's output is nil",
		},
		{
			name: "nil output type definition",
			operation: Distribution{
				Input: primitiveVariable("input", "string"),
				Conditions: []DistributionCondition{
					{
						Condition: "input != \"\"",
						Output:    variableWithoutType("output"),
					},
				},
			},
			want: "distribution condition's output type definition is nil",
		},
		{
			name: "nil function call",
			operation: Distribution{
				Input: primitiveVariable("input", "string"),
				Conditions: []DistributionCondition{
					{
						Condition:     "input != \"\"",
						Output:        primitiveVariable("output", "string"),
						FunctionCalls: []*FunctionDefinition{nil},
					},
				},
			},
			want: "function call in distribution condition is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.operation.GetPrompt()
			if err == nil {
				t.Fatal("GetPrompt() returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), tt.want)
			}
		})
	}
}

func TestFilterImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Filter{}
}

func TestDistributionImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Distribution{}
}

func variableWithoutType(name string) *VariableDefinition {
	variable := VariableDefinition(&BasicVariableDefinition{Name: name})
	return &variable
}

func collectionVariableWithoutType(name string) *CollectionVariableDefinition {
	return &CollectionVariableDefinition{Name: name}
}

func assertContains(t *testing.T, value, substring string) {
	t.Helper()

	if !strings.Contains(value, substring) {
		t.Fatalf("expected %q to contain %q", value, substring)
	}
}
