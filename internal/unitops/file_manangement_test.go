package unitops

import "testing"

func TestBuildFileDefinitions(t *testing.T) {
	userType := TypeDefinition(StructTypeDefinition{
		TypeName: "User",
		FilePath: "internal/user.go",
	})
	accountType := TypeDefinition(StructTypeDefinition{
		TypeName: "Account",
		FilePath: "internal/account.go",
	})
	normalizeUserFunction := FunctionDefinition(BasicFunctionDefinition{
		FunctionName: "NormalizeUser",
		FilePath:     "internal/user.go",
	})
	loadUserFunction := FunctionDefinition(BasicFunctionDefinition{
		FunctionName: "LoadUser",
		FilePath:     "internal/user.go",
	})
	normalizeUserSuite := TestSuite(BasicTestSuite{
		Name:     "NormalizeUser",
		FilePath: "internal/user_test.go",
	})

	fileDefinitions, err := BuildFileDefinitions(
		[]*TestSuite{&normalizeUserSuite},
		[]*TypeDefinition{&userType, &accountType},
		[]*FunctionDefinition{&normalizeUserFunction, &loadUserFunction},
	)
	if err != nil {
		t.Fatalf("BuildFileDefinitions() returned error: %v", err)
	}

	if len(fileDefinitions) != 3 {
		t.Fatalf("BuildFileDefinitions() returned %d files, want 3", len(fileDefinitions))
	}

	assertFileDefinition(t, fileDefinitions[0], "internal/account.go", 0, 1, 0)
	assertFileDefinition(t, fileDefinitions[1], "internal/user.go", 0, 1, 2)
	assertFileDefinition(t, fileDefinitions[2], "internal/user_test.go", 1, 0, 0)

	if (*fileDefinitions[1].TypeDefinitions[0]).GetTypeName() != "User" {
		t.Fatalf("user file type = %q, want %q", (*fileDefinitions[1].TypeDefinitions[0]).GetTypeName(), "User")
	}

	if (*fileDefinitions[1].FunctionDefinitions[0]).GetFunctionName() != "NormalizeUser" {
		t.Fatalf("first user file function = %q, want %q", (*fileDefinitions[1].FunctionDefinitions[0]).GetFunctionName(), "NormalizeUser")
	}

	if (*fileDefinitions[1].FunctionDefinitions[1]).GetFunctionName() != "LoadUser" {
		t.Fatalf("second user file function = %q, want %q", (*fileDefinitions[1].FunctionDefinitions[1]).GetFunctionName(), "LoadUser")
	}
}

func TestBuildFileDefinitionsErrors(t *testing.T) {
	tests := []struct {
		name                string
		testSuites          []*TestSuite
		typeDefinitions     []*TypeDefinition
		functionDefinitions []*FunctionDefinition
		want                string
	}{
		{
			name:       "nil test suite",
			testSuites: []*TestSuite{nil},
			want:       "test suite is nil",
		},
		{
			name:            "nil type definition",
			typeDefinitions: []*TypeDefinition{nil},
			want:            "type definition is nil",
		},
		{
			name:                "nil function definition",
			functionDefinitions: []*FunctionDefinition{nil},
			want:                "function definition is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := BuildFileDefinitions(tt.testSuites, tt.typeDefinitions, tt.functionDefinitions)
			if err == nil {
				t.Fatal("BuildFileDefinitions() returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("BuildFileDefinitions() error = %q, want %q", err.Error(), tt.want)
			}
		})
	}
}

func assertFileDefinition(
	t *testing.T,
	fileDefinition *FileDefinition,
	filePath string,
	testSuiteCount int,
	typeDefinitionCount int,
	functionDefinitionCount int,
) {
	t.Helper()

	if fileDefinition == nil {
		t.Fatal("file definition is nil")
	}

	if fileDefinition.FilePath != filePath {
		t.Fatalf("FilePath = %q, want %q", fileDefinition.FilePath, filePath)
	}

	if len(fileDefinition.TestSuites) != testSuiteCount {
		t.Fatalf("len(TestSuites) = %d, want %d", len(fileDefinition.TestSuites), testSuiteCount)
	}

	if len(fileDefinition.TypeDefinitions) != typeDefinitionCount {
		t.Fatalf("len(TypeDefinitions) = %d, want %d", len(fileDefinition.TypeDefinitions), typeDefinitionCount)
	}

	if len(fileDefinition.FunctionDefinitions) != functionDefinitionCount {
		t.Fatalf("len(FunctionDefinitions) = %d, want %d", len(fileDefinition.FunctionDefinitions), functionDefinitionCount)
	}
}
