package unitops

type BasicProductionFileDefinition struct {
	DefinitionPrompt    string
	FilePath            string
	CodeComments        string
	TypeDefinitions     []TypeDefinition
	FunctionDefinitions []FunctionDefinition
}

func (p BasicProductionFileDefinition) FileDefinitionPrompt() string {
	return p.DefinitionPrompt
}

func (p BasicProductionFileDefinition) GetFilePath() string {
	return p.FilePath
}

func (p BasicProductionFileDefinition) GetCodeComments() string {
	return p.CodeComments
}

func (p BasicProductionFileDefinition) GetTypeDefinitions() []TypeDefinition {
	return p.TypeDefinitions
}

func (p BasicProductionFileDefinition) GetFunctionDefinitions() []FunctionDefinition {
	return p.FunctionDefinitions
}
