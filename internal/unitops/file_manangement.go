package unitops

import (
	"fmt"
	"sort"
	"strings"
)

type FileDefinition struct {
	FilePath            string
	TestSuites          []*TestSuite
	TypeDefinitions     []*TypeDefinition
	FunctionDefinitions []*FunctionDefinition
}

func BuildFileDefinitions(
	testSuites []*TestSuite,
	typeDefinitions []*TypeDefinition,
	functionDefinitions []*FunctionDefinition,
) ([]*FileDefinition, error) {
	filesByPath := map[string]*FileDefinition{}

	fileForPath := func(filePath string) *FileDefinition {
		fileDefinition, ok := filesByPath[filePath]
		if !ok {
			fileDefinition = &FileDefinition{FilePath: filePath}
			filesByPath[filePath] = fileDefinition
		}

		return fileDefinition
	}

	for _, testSuite := range testSuites {
		if testSuite == nil {
			return nil, fmt.Errorf("test suite is nil")
		}

		fileDefinition := fileForPath((*testSuite).GetFilePath())
		fileDefinition.TestSuites = append(fileDefinition.TestSuites, testSuite)
	}

	for _, typeDefinition := range typeDefinitions {
		if typeDefinition == nil {
			return nil, fmt.Errorf("type definition is nil")
		}

		fileDefinition := fileForPath((*typeDefinition).GetFilePath())
		fileDefinition.TypeDefinitions = append(fileDefinition.TypeDefinitions, typeDefinition)
	}

	for _, functionDefinition := range functionDefinitions {
		if functionDefinition == nil {
			return nil, fmt.Errorf("function definition is nil")
		}

		fileDefinition := fileForPath((*functionDefinition).GetFilePath())
		fileDefinition.FunctionDefinitions = append(fileDefinition.FunctionDefinitions, functionDefinition)
	}

	filePaths := make([]string, 0, len(filesByPath))
	for filePath := range filesByPath {
		filePaths = append(filePaths, filePath)
	}
	sort.Strings(filePaths)

	fileDefinitions := make([]*FileDefinition, 0, len(filePaths))
	for _, filePath := range filePaths {
		fileDefinitions = append(fileDefinitions, filesByPath[filePath])
	}

	return fileDefinitions, nil
}

func (f FileDefinition) GetPrompt() (string, error) {
	var sb strings.Builder

	_, err := fmt.Fprintf(&sb, "# File Definition For %s\n", f.FilePath)
	if err != nil {
		return "", err
	}

	if len(f.TypeDefinitions) > 0 {
		_, err = fmt.Fprintln(&sb, "## Types")
		if err != nil {
			return "", err
		}
	}

	for _, typeDefinition := range f.TypeDefinitions {
		if typeDefinition == nil {
			return "", fmt.Errorf("type definition in file is nil")
		}

		_, err = fmt.Fprintf(&sb, "- %s\n", (*typeDefinition).GetTypeName())
		if err != nil {
			return "", err
		}
	}

	if len(f.FunctionDefinitions) > 0 {
		_, err = fmt.Fprintln(&sb, "## Functions")
		if err != nil {
			return "", err
		}
	}

	for _, functionDefinition := range f.FunctionDefinitions {
		if functionDefinition == nil {
			return "", fmt.Errorf("function definition in file is nil")
		}

		_, err = fmt.Fprintf(&sb, "- %s\n", (*functionDefinition).GetFunctionName())
		if err != nil {
			return "", err
		}
	}

	if len(f.TestSuites) > 0 {
		_, err = fmt.Fprintln(&sb, "## Test Suites")
		if err != nil {
			return "", err
		}
	}

	for _, testSuite := range f.TestSuites {
		if testSuite == nil {
			return "", fmt.Errorf("test suite in file is nil")
		}

		_, err = fmt.Fprintf(&sb, "- %s\n", (*testSuite).GetID())
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}
