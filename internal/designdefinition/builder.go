// Package designdefinition
package designdefinition

import (
	"github.com/jfkonecn/gophitect/internal/unitops"
)

type DesignDefinitionSpec struct {
	FileDefinitions []*unitops.FileDefinition
}

func New() DesignDefinitionSpec {
	return DesignDefinitionSpec{}
}

func (spec *DesignDefinitionSpec) AddProductionFile(f *unitops.FileDefinition) {
	if f == nil {
		panic("AddProductionFile: FileDefinition cannot be nil")
	}
	spec.FileDefinitions = append(spec.FileDefinitions, f)
}

func (spec *DesignDefinitionSpec) AddTestFile(f *unitops.FileDefinition) {
	if f == nil {
		panic("AddTestFile: FileDefinition cannot be nil")
	}
	spec.FileDefinitions = append(spec.FileDefinitions, f)
}

type TypeDefinition struct {
	name           string
	filePath       string
	prompt         string
	dependentTypes []*TypeDefinition
}

type FunctionDefinition struct {
	name        string
	filePath    string
	prompt      string
	inputTypes  []*TypeDefinition
	outputTypes []*TypeDefinition
}

type DesignDefinition struct {
	Types []*TypeDefinition
}

func (spec *DesignDefinitionSpec) Build() DesignDefinition {
	// dfs for type definitions
	d := DesignDefinition{}

	return d
}
