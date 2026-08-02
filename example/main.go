package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jfkonecn/gophitect/internal/unitops"
)

const ollamaModel = "gemma4:latest"

func main() {
	userType := unitops.TypeDefinition(unitops.StructTypeDefinition{
		ID:                             "type:user",
		TypeName:                       "User",
		FilePath:                       "example/generated/user.go",
		ProgrammingLanguageInformation: "Define the type in Go.",
		PackageNamespace:               "Put this type in package generated.",
		CodeComments:                   "User represents a person in the system.",
	})
	normalizedUserType := unitops.TypeDefinition(unitops.StructTypeDefinition{
		ID:                             "type:normalized-user",
		TypeName:                       "NormalizedUser",
		FilePath:                       "example/generated/user.go",
		ProgrammingLanguageInformation: "Define the type in Go.",
		PackageNamespace:               "Put this type in package generated.",
		CodeComments:                   "NormalizedUser is safe to display and persist.",
		Fields: []*unitops.StructTypeField{
			{
				Name:           "Original",
				CodeComment:    "The original user value before normalization.",
				TypeDefinition: &userType,
			},
		},
	})

	input := unitops.VariableDefinition(&unitops.BasicVariableDefinition{
		Name:           "user",
		TypeDefinition: &userType,
	})
	output := unitops.VariableDefinition(&unitops.BasicVariableDefinition{
		Name:           "normalizedUser",
		TypeDefinition: &normalizedUserType,
	})
	operation := unitops.UnitOperation(unitops.Map{
		Input:        &input,
		Output:       &output,
		CodeComments: "Trim display fields and preserve the original user data.",
	})
	normalizeUser := unitops.FunctionDefinition(unitops.BasicFunctionDefinition{
		ID:                             "function:normalize-user",
		FunctionName:                   "NormalizeUser",
		FilePath:                       "example/generated/user.go",
		ProgrammingLanguageInformation: "Implement the function in Go.",
		PackageNamespace:               "Put this function in package generated.",
		CodeComments:                   "NormalizeUser returns a display-safe user value.",
		UnitOperations:                 []*unitops.UnitOperation{&operation},
	})
	validUserCase := unitops.TestCase(unitops.BasicTestCase{
		Name:        "valid user",
		Description: "Returns a normalized user for populated input.",
	})
	testSuite := unitops.TestSuite(unitops.BasicTestSuite{
		ID:                             "test:normalize-user",
		Name:                           "NormalizeUser",
		FilePath:                       "example/generated/user_test.go",
		ProgrammingLanguageInformation: "Write the tests in Go using the testing package.",
		PackageNamespace:               "Put these tests in package generated.",
		FunctionDefinition:             &normalizeUser,
		TestCases:                      []*unitops.TestCase{&validUserCase},
	})

	runner := unitops.ImplementationRunner{
		PromptLLM: promptOllama,
		ReadFile:  readFile,
		WriteFile: writeFile,
	}

	err := runner.Implement(
		[]*unitops.FunctionDefinition{&normalizeUser},
		[]*unitops.TestSuite{&testSuite},
		nil,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("generated example/generated/user.go and example/generated/user_test.go")
}

func promptOllama(prompt string) (string, error) {
	requestBody, err := json.Marshal(map[string]any{
		"model":  ollamaModel,
		"prompt": prompt,
		"stream": false,
	})
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, ollamaURL()+"/api/generate", bytes.NewReader(requestBody))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Minute}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("ollama returned %s: %s", response.Status, string(body))
	}

	var generateResponse struct {
		Response string `json:"response"`
	}
	if err := json.Unmarshal(body, &generateResponse); err != nil {
		return "", err
	}

	return generateResponse.Response, nil
}

func ollamaURL() string {
	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		return "http://localhost:11434"
	}

	return strings.TrimRight(host, "/")
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func writeFile(filePath string, content string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return err
	}

	return os.WriteFile(filePath, []byte(content), 0o644)
}
