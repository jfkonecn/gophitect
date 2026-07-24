// Package designdefinition
package designdefinition

import (
	"github.com/jfkonecn/gophitect/internal/unitops"
)

type DesignDefinitionSpec struct {
	TestFileDefinitions       []*unitops.TestFileDefinition
	ProductionFileDefinitions []*unitops.ProductionFileDefinition
}

func New() DesignDefinitionSpec {
	return DesignDefinitionSpec{}
}

func (spec *DesignDefinitionSpec) AddProductionFile(f *unitops.ProductionFileDefinition) {
	if f == nil {
		panic("AddProductionFile: ProductionFileDefinition cannot be nil")
	}
	spec.ProductionFileDefinitions = append(spec.ProductionFileDefinitions, f)
}

func (spec *DesignDefinitionSpec) AddTestFile(f *unitops.TestFileDefinition) {
	if f == nil {
		panic("AddTestFile: TestFileDefinition cannot be nil")
	}
	spec.TestFileDefinitions = append(spec.TestFileDefinitions, f)
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

}
