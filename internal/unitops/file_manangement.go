package unitops

import (
	"fmt"
	"sort"
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
