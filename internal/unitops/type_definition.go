package unitops

type BasicTypeDefinition struct {
	TypeName       string
	CodeComments   string
	DependentTypes []TypeDefinition
	FilePath       string
	ShouldStub     bool
}

func (t BasicTypeDefinition) GetTypeName() string {
	return t.TypeName
}

func (t BasicTypeDefinition) GetCodeComments() string {
	return t.CodeComments
}

func (t BasicTypeDefinition) GetDependentTypes() []TypeDefinition {
	return t.DependentTypes
}

func (t BasicTypeDefinition) GetFilePath() string {
	return t.FilePath
}

func (t BasicTypeDefinition) GetShouldStub() bool {
	return t.ShouldStub
}
