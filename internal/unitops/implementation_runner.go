package unitops

import "fmt"

type PromptLLMFunc func(prompt string) (string, error)
type ReadFileFunc func(filePath string) (string, error)
type WriteFileFunc func(filePath string, content string) error

type ImplementationRunner struct {
	PromptLLM PromptLLMFunc
	ReadFile  ReadFileFunc
	WriteFile WriteFileFunc
}

func (r ImplementationRunner) Implement(
	functionDefinitions []*FunctionDefinition,
	testSuites []*TestSuite,
	completedTypeIDs []string,
	completedFunctionIDs []string,
	completedTestIDs []string,
) error {
	if r.PromptLLM == nil {
		return fmt.Errorf("prompt LLM function is nil")
	}
	if r.ReadFile == nil {
		return fmt.Errorf("read file function is nil")
	}
	if r.WriteFile == nil {
		return fmt.Errorf("write file function is nil")
	}

	typeDefinitions, err := collectTypeDefinitions(functionDefinitions)
	if err != nil {
		return err
	}

	completedTypes := idSet(completedTypeIDs)
	if err := r.implementTypes(typeDefinitions, completedTypes); err != nil {
		return err
	}

	functionDefinitions, err = collectFunctionDefinitions(functionDefinitions)
	if err != nil {
		return err
	}

	completedFunctions := idSet(completedFunctionIDs)
	if err := r.implementFunctions(functionDefinitions, completedFunctions); err != nil {
		return err
	}

	completedTests := idSet(completedTestIDs)
	return r.implementTestSuites(testSuites, completedTests)
}

func (r ImplementationRunner) implementTypes(typeDefinitions []*TypeDefinition, completed map[string]bool) error {
	remaining := map[string]*TypeDefinition{}
	for _, typeDefinition := range typeDefinitions {
		if typeDefinition == nil {
			return fmt.Errorf("type definition is nil")
		}

		typeValue := *typeDefinition
		if typeValue.IsBuiltin() || completed[typeValue.GetID()] {
			continue
		}

		remaining[typeValue.GetID()] = typeDefinition
	}

	for len(remaining) > 0 {
		implementedAny := false
		for id, typeDefinition := range remaining {
			if !typeDependenciesComplete(*typeDefinition, completed) {
				continue
			}

			if err := r.implementDefinition((*typeDefinition).GetFilePath(), *typeDefinition); err != nil {
				return err
			}

			completed[id] = true
			delete(remaining, id)
			implementedAny = true
		}

		if !implementedAny {
			return fmt.Errorf("type definitions have unresolved dependencies")
		}
	}

	return nil
}

func (r ImplementationRunner) implementFunctions(functionDefinitions []*FunctionDefinition, completed map[string]bool) error {
	remaining := map[string]*FunctionDefinition{}
	for _, functionDefinition := range functionDefinitions {
		if functionDefinition == nil {
			return fmt.Errorf("function definition is nil")
		}

		functionValue := *functionDefinition
		if completed[functionValue.GetID()] {
			continue
		}

		remaining[functionValue.GetID()] = functionDefinition
	}

	for len(remaining) > 0 {
		implementedAny := false
		for id, functionDefinition := range remaining {
			if !functionDependenciesComplete(*functionDefinition, completed) {
				continue
			}

			if err := r.implementDefinition((*functionDefinition).GetFilePath(), *functionDefinition); err != nil {
				return err
			}

			completed[id] = true
			delete(remaining, id)
			implementedAny = true
		}

		if !implementedAny {
			return fmt.Errorf("function definitions have unresolved dependencies")
		}
	}

	return nil
}

func (r ImplementationRunner) implementTestSuites(testSuites []*TestSuite, completed map[string]bool) error {
	for _, testSuite := range testSuites {
		if testSuite == nil {
			return fmt.Errorf("test suite is nil")
		}

		testSuiteValue := *testSuite
		if completed[testSuiteValue.GetID()] {
			continue
		}

		if err := r.implementDefinition(testSuiteValue.GetFilePath(), testSuiteValue); err != nil {
			return err
		}

		completed[testSuiteValue.GetID()] = true
	}

	return nil
}

type promptDefinition interface {
	GetPrompt() (string, error)
}

func (r ImplementationRunner) implementDefinition(filePath string, definition promptDefinition) error {
	existingContent, err := r.ReadFile(filePath)
	if err != nil {
		return err
	}

	prompt, err := definition.GetPrompt()
	if err != nil {
		return err
	}

	response, err := r.PromptLLM(fmt.Sprintf(
		"Existing file content:\n%s\n\n%s\n\n%s",
		existingContent,
		prompt,
		completeFileContentsPrompt,
	))
	if err != nil {
		return err
	}

	return r.WriteFile(filePath, response)
}

func collectTypeDefinitions(functionDefinitions []*FunctionDefinition) ([]*TypeDefinition, error) {
	typeDefinitionsByID := map[string]*TypeDefinition{}

	var addTypeDefinition func(*TypeDefinition) error
	addTypeDefinition = func(typeDefinition *TypeDefinition) error {
		if typeDefinition == nil {
			return nil
		}

		typeValue := *typeDefinition
		if typeValue.IsBuiltin() {
			return nil
		}
		if _, ok := typeDefinitionsByID[typeValue.GetID()]; ok {
			return nil
		}

		typeDefinitionsByID[typeValue.GetID()] = typeDefinition
		for _, dependentType := range typeValue.GetDependentTypes() {
			if err := addTypeDefinition(dependentType); err != nil {
				return err
			}
		}

		return nil
	}

	for _, functionDefinition := range functionDefinitions {
		if functionDefinition == nil {
			return nil, fmt.Errorf("function definition is nil")
		}

		functionValue := *functionDefinition
		for _, input := range functionValue.GetInputs() {
			if input == nil {
				continue
			}

			if err := addTypeDefinition((*input).GetTypeDefinition()); err != nil {
				return nil, err
			}
		}

		if err := addTypeDefinition(functionValue.GetOutput()); err != nil {
			return nil, err
		}
	}

	typeDefinitions := make([]*TypeDefinition, 0, len(typeDefinitionsByID))
	for _, typeDefinition := range typeDefinitionsByID {
		typeDefinitions = append(typeDefinitions, typeDefinition)
	}

	return typeDefinitions, nil
}

func collectFunctionDefinitions(functionDefinitions []*FunctionDefinition) ([]*FunctionDefinition, error) {
	functionDefinitionsByID := map[string]*FunctionDefinition{}

	var addFunctionDefinition func(*FunctionDefinition) error
	addFunctionDefinition = func(functionDefinition *FunctionDefinition) error {
		if functionDefinition == nil {
			return nil
		}

		functionValue := *functionDefinition
		if _, ok := functionDefinitionsByID[functionValue.GetID()]; ok {
			return nil
		}

		functionDefinitionsByID[functionValue.GetID()] = functionDefinition

		for _, unitOperation := range functionValue.GetUnitOperations() {
			if unitOperation == nil {
				continue
			}

			for _, functionCall := range (*unitOperation).GetFunctionCalls() {
				if err := addFunctionDefinition(functionCall); err != nil {
					return err
				}
			}
		}

		return nil
	}

	for _, functionDefinition := range functionDefinitions {
		if functionDefinition == nil {
			return nil, fmt.Errorf("function definition is nil")
		}

		if err := addFunctionDefinition(functionDefinition); err != nil {
			return nil, err
		}
	}

	collectedFunctionDefinitions := make([]*FunctionDefinition, 0, len(functionDefinitionsByID))
	for _, functionDefinition := range functionDefinitionsByID {
		collectedFunctionDefinitions = append(collectedFunctionDefinitions, functionDefinition)
	}

	return collectedFunctionDefinitions, nil
}

func typeDependenciesComplete(typeDefinition TypeDefinition, completed map[string]bool) bool {
	for _, dependentType := range typeDefinition.GetDependentTypes() {
		if dependentType == nil {
			continue
		}

		dependentTypeValue := *dependentType
		if dependentTypeValue.IsBuiltin() {
			continue
		}

		if !completed[dependentTypeValue.GetID()] {
			return false
		}
	}

	return true
}

func functionDependenciesComplete(functionDefinition FunctionDefinition, completed map[string]bool) bool {
	for _, unitOperation := range functionDefinition.GetUnitOperations() {
		if unitOperation == nil {
			continue
		}

		for _, functionCall := range (*unitOperation).GetFunctionCalls() {
			if functionCall == nil {
				continue
			}

			if !completed[(*functionCall).GetID()] {
				return false
			}
		}
	}

	return true
}

func idSet(ids []string) map[string]bool {
	set := map[string]bool{}
	for _, id := range ids {
		set[id] = true
	}

	return set
}
