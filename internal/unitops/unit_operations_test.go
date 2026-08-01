package unitops

import (
	"strings"
	"testing"
)

type testFunctionDefinition struct {
	name string
}

func (f testFunctionDefinition) GetID() string {
	return f.name
}

func (f testFunctionDefinition) GetFunctionName() string {
	return f.name
}

func (f testFunctionDefinition) GetFilePath() string {
	return ""
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
		Input:        primitiveVariable("payment", "Payment"),
		CodeComments: "Return after the first matching condition.",
		Conditions: []DistributionCondition{
			{
				Condition:     "payment.Amount > 1000",
				Output:        primitiveVariable("manualReviewQueue", "ReviewQueueItem"),
				CodeComments:  "Keep the original payment ID.",
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
	assertContains(t, prompt, "- Add these comments\nReturn after the first matching condition.")
	assertContains(t, prompt, "1. When payment.Amount > 1000")
	assertContains(t, prompt, "- Route to manualReviewQueue of type ReviewQueueItem")
	assertContains(t, prompt, "- Add these comments for this condition\nKeep the original payment ID.")
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

func TestValidateGetPrompt(t *testing.T) {
	operation := Validate{
		Input:         primitiveVariable("user", "User"),
		SuccessOutput: primitiveVariable("validUser", "User"),
		FailureOutput: primitiveVariable("validationError", "ValidationError"),
		CodeComments:  "Include every failed field in the failure output.",
		FunctionCalls: []*FunctionDefinition{functionCall("validateUser")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "validate user of type User")
	assertContains(t, prompt, "Route the input based on whether validation succeeds or fails.")
	assertContains(t, prompt, "- On success, route to validUser of type User")
	assertContains(t, prompt, "- On failure, route to validationError of type ValidationError")
	assertContains(t, prompt, "- Add these comments\nInclude every failed field in the failure output.")
	assertContains(t, prompt, "Make sure you use these functions in the validate logic")
	assertContains(t, prompt, "- validateUser")
}

func TestValidateGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation Validate
		want      string
	}{
		{
			name: "nil input",
			operation: Validate{
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "validate's input is nil",
		},
		{
			name: "nil input type definition",
			operation: Validate{
				Input:         variableWithoutType("input"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "validate's input type definition is nil",
		},
		{
			name: "nil success output",
			operation: Validate{
				Input:         primitiveVariable("input", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "validate's success output is nil",
		},
		{
			name: "nil success output type definition",
			operation: Validate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: variableWithoutType("success"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "validate's success output type definition is nil",
		},
		{
			name: "nil failure output",
			operation: Validate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
			},
			want: "validate's failure output is nil",
		},
		{
			name: "nil failure output type definition",
			operation: Validate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: variableWithoutType("failure"),
			},
			want: "validate's failure output type definition is nil",
		},
		{
			name: "nil function call",
			operation: Validate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in validate is nil",
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

func TestAuthenticateGetPrompt(t *testing.T) {
	operation := Authenticate{
		Input:         primitiveVariable("credentials", "Credentials"),
		SuccessOutput: primitiveVariable("authenticatedUser", "AuthenticatedUser"),
		FailureOutput: primitiveVariable("authFailure", "AuthenticationError"),
		CodeComments:  "Do not log raw credentials.",
		FunctionCalls: []*FunctionDefinition{functionCall("authenticateCredentials")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "authenticate credentials of type Credentials")
	assertContains(t, prompt, "Route the input based on whether authentication succeeds or fails.")
	assertContains(t, prompt, "- On success, route to authenticatedUser of type AuthenticatedUser")
	assertContains(t, prompt, "- On failure, route to authFailure of type AuthenticationError")
	assertContains(t, prompt, "- Add these comments\nDo not log raw credentials.")
	assertContains(t, prompt, "Make sure you use these functions in the authenticate logic")
	assertContains(t, prompt, "- authenticateCredentials")
}

func TestAuthenticateGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation Authenticate
		want      string
	}{
		{
			name: "nil input",
			operation: Authenticate{
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "authenticate's input is nil",
		},
		{
			name: "nil input type definition",
			operation: Authenticate{
				Input:         variableWithoutType("input"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "authenticate's input type definition is nil",
		},
		{
			name: "nil success output",
			operation: Authenticate{
				Input:         primitiveVariable("input", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "authenticate's success output is nil",
		},
		{
			name: "nil success output type definition",
			operation: Authenticate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: variableWithoutType("success"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "authenticate's success output type definition is nil",
		},
		{
			name: "nil failure output",
			operation: Authenticate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
			},
			want: "authenticate's failure output is nil",
		},
		{
			name: "nil failure output type definition",
			operation: Authenticate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: variableWithoutType("failure"),
			},
			want: "authenticate's failure output type definition is nil",
		},
		{
			name: "nil function call",
			operation: Authenticate{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in authenticate is nil",
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

func TestAuthorizeGetPrompt(t *testing.T) {
	operation := Authorize{
		Input:               primitiveVariable("request", "AuthorizedRequest"),
		AuthorizeConditions: "Allow only when request.UserID matches resource owner or user has admin role.",
		SuccessOutput:       primitiveVariable("authorizedRequest", "AuthorizedRequest"),
		FailureOutput:       primitiveVariable("authorizationFailure", "AuthorizationError"),
		CodeComments:        "Audit denied authorization attempts.",
		FunctionCalls:       []*FunctionDefinition{functionCall("authorizeRequest")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "authorize request of type AuthorizedRequest")
	assertContains(t, prompt, "Use these conditions to authorize\nAllow only when request.UserID matches resource owner or user has admin role.")
	assertContains(t, prompt, "Route the input based on whether authorization succeeds or fails.")
	assertContains(t, prompt, "- On success, route to authorizedRequest of type AuthorizedRequest")
	assertContains(t, prompt, "- On failure, route to authorizationFailure of type AuthorizationError")
	assertContains(t, prompt, "- Add these comments\nAudit denied authorization attempts.")
	assertContains(t, prompt, "Make sure you use these functions in the authorize logic")
	assertContains(t, prompt, "- authorizeRequest")
}

func TestAuthorizeGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation Authorize
		want      string
	}{
		{
			name: "nil input",
			operation: Authorize{
				AuthorizeConditions: "input.UserID != \"\"",
				SuccessOutput:       primitiveVariable("success", "string"),
				FailureOutput:       primitiveVariable("failure", "error"),
			},
			want: "authorize's input is nil",
		},
		{
			name: "nil input type definition",
			operation: Authorize{
				Input:               variableWithoutType("input"),
				AuthorizeConditions: "input.UserID != \"\"",
				SuccessOutput:       primitiveVariable("success", "string"),
				FailureOutput:       primitiveVariable("failure", "error"),
			},
			want: "authorize's input type definition is nil",
		},
		{
			name: "empty conditions",
			operation: Authorize{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "authorize's conditions are empty",
		},
		{
			name: "nil success output",
			operation: Authorize{
				Input:               primitiveVariable("input", "string"),
				AuthorizeConditions: "input.UserID != \"\"",
				FailureOutput:       primitiveVariable("failure", "error"),
			},
			want: "authorize's success output is nil",
		},
		{
			name: "nil success output type definition",
			operation: Authorize{
				Input:               primitiveVariable("input", "string"),
				AuthorizeConditions: "input.UserID != \"\"",
				SuccessOutput:       variableWithoutType("success"),
				FailureOutput:       primitiveVariable("failure", "error"),
			},
			want: "authorize's success output type definition is nil",
		},
		{
			name: "nil failure output",
			operation: Authorize{
				Input:               primitiveVariable("input", "string"),
				AuthorizeConditions: "input.UserID != \"\"",
				SuccessOutput:       primitiveVariable("success", "string"),
			},
			want: "authorize's failure output is nil",
		},
		{
			name: "nil failure output type definition",
			operation: Authorize{
				Input:               primitiveVariable("input", "string"),
				AuthorizeConditions: "input.UserID != \"\"",
				SuccessOutput:       primitiveVariable("success", "string"),
				FailureOutput:       variableWithoutType("failure"),
			},
			want: "authorize's failure output type definition is nil",
		},
		{
			name: "nil function call",
			operation: Authorize{
				Input:               primitiveVariable("input", "string"),
				AuthorizeConditions: "input.UserID != \"\"",
				SuccessOutput:       primitiveVariable("success", "string"),
				FailureOutput:       primitiveVariable("failure", "error"),
				FunctionCalls:       []*FunctionDefinition{nil},
			},
			want: "function call in authorize is nil",
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

func TestGlobalStateReadGetPrompt(t *testing.T) {
	operation := GlobalStateRead{
		Output:        primitiveVariable("session", "Session"),
		CodeComments:  "Read from the request-scoped session store.",
		FunctionCalls: []*FunctionDefinition{functionCall("readSession")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "read global state into session of type Session")
	assertContains(t, prompt, "- Add these comments\nRead from the request-scoped session store.")
	assertContains(t, prompt, "Make sure you use these functions in the global state read logic")
	assertContains(t, prompt, "- readSession")
}

func TestGlobalStateReadGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation GlobalStateRead
		want      string
	}{
		{
			name:      "nil output",
			operation: GlobalStateRead{},
			want:      "global state read's output is nil",
		},
		{
			name: "nil output type definition",
			operation: GlobalStateRead{
				Output: variableWithoutType("output"),
			},
			want: "global state read's output type definition is nil",
		},
		{
			name: "nil function call",
			operation: GlobalStateRead{
				Output:        primitiveVariable("output", "string"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in global state read is nil",
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

func TestGlobalStateWriteGetPrompt(t *testing.T) {
	operation := GlobalStateWrite{
		Input:         primitiveVariable("session", "Session"),
		CodeComments:  "Overwrite the current request session.",
		FunctionCalls: []*FunctionDefinition{functionCall("writeSession")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "write session of type Session to global state")
	assertContains(t, prompt, "- Add these comments\nOverwrite the current request session.")
	assertContains(t, prompt, "Make sure you use these functions in the global state write logic")
	assertContains(t, prompt, "- writeSession")
}

func TestGlobalStateWriteGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation GlobalStateWrite
		want      string
	}{
		{
			name:      "nil input",
			operation: GlobalStateWrite{},
			want:      "global state write's input is nil",
		},
		{
			name: "nil input type definition",
			operation: GlobalStateWrite{
				Input: variableWithoutType("input"),
			},
			want: "global state write's input type definition is nil",
		},
		{
			name: "nil function call",
			operation: GlobalStateWrite{
				Input:         primitiveVariable("input", "string"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in global state write is nil",
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

func TestInputOutputGetPrompt(t *testing.T) {
	operation := InputOutput{
		Input:         primitiveVariable("request", "HTTPRequest"),
		SuccessOutput: primitiveVariable("response", "HTTPResponse"),
		FailureOutput: primitiveVariable("requestError", "RequestError"),
		CodeComments:  "Set a timeout before making the request.",
		FunctionCalls: []*FunctionDefinition{functionCall("sendRequest")},
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "perform input/output with request of type HTTPRequest")
	assertContains(t, prompt, "Route the input based on whether the input/output operation succeeds or fails.")
	assertContains(t, prompt, "- On success, route to response of type HTTPResponse")
	assertContains(t, prompt, "- On failure, route to requestError of type RequestError")
	assertContains(t, prompt, "- Add these comments\nSet a timeout before making the request.")
	assertContains(t, prompt, "Make sure you use these functions in the input/output logic")
	assertContains(t, prompt, "- sendRequest")
}

func TestInputOutputGetPromptErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation InputOutput
		want      string
	}{
		{
			name: "nil input",
			operation: InputOutput{
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "input/output's input is nil",
		},
		{
			name: "nil input type definition",
			operation: InputOutput{
				Input:         variableWithoutType("input"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "input/output's input type definition is nil",
		},
		{
			name: "nil success output",
			operation: InputOutput{
				Input:         primitiveVariable("input", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "input/output's success output is nil",
		},
		{
			name: "nil success output type definition",
			operation: InputOutput{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: variableWithoutType("success"),
				FailureOutput: primitiveVariable("failure", "error"),
			},
			want: "input/output's success output type definition is nil",
		},
		{
			name: "nil failure output",
			operation: InputOutput{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
			},
			want: "input/output's failure output is nil",
		},
		{
			name: "nil failure output type definition",
			operation: InputOutput{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: variableWithoutType("failure"),
			},
			want: "input/output's failure output type definition is nil",
		},
		{
			name: "nil function call",
			operation: InputOutput{
				Input:         primitiveVariable("input", "string"),
				SuccessOutput: primitiveVariable("success", "string"),
				FailureOutput: primitiveVariable("failure", "error"),
				FunctionCalls: []*FunctionDefinition{nil},
			},
			want: "function call in input/output is nil",
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

func TestPanicGetPrompt(t *testing.T) {
	operation := Panic{
		Description:  "Stop processing because the invariant cannot be recovered.",
		CodeComments: "Include the invariant name in the panic message.",
	}

	prompt, err := operation.GetPrompt()
	if err != nil {
		t.Fatalf("GetPrompt() returned error: %v", err)
	}

	assertContains(t, prompt, "panic with this description\nStop processing because the invariant cannot be recovered.")
	assertContains(t, prompt, "- Add these comments\nInclude the invariant name in the panic message.")
}

func TestPanicGetPromptErrors(t *testing.T) {
	_, err := (Panic{}).GetPrompt()
	if err == nil {
		t.Fatal("GetPrompt() returned nil error")
	}

	if err.Error() != "panic's description is empty" {
		t.Fatalf("GetPrompt() error = %q, want %q", err.Error(), "panic's description is empty")
	}
}

func TestFilterImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Filter{}
}

func TestDistributionImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Distribution{}
}

func TestValidateImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Validate{}
}

func TestAuthenticateImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Authenticate{}
}

func TestAuthorizeImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Authorize{}
}

func TestGlobalStateReadImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = GlobalStateRead{}
}

func TestGlobalStateWriteImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = GlobalStateWrite{}
}

func TestInputOutputImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = InputOutput{}
}

func TestPanicImplementsUnitOperation(t *testing.T) {
	var _ UnitOperation = Panic{}
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
