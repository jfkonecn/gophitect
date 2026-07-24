package unitops

// A Map unit operation converts one value or type into another.
type Map struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (m Map) GetUnitOperationTypeDefinition() string {
	return "A Map unit operation converts one value or type into another."
}

func (m Map) GetInputTypes() []*TypeDefinition {
	return m.InputTypes
}

func (m Map) GetOutputTypes() []*TypeDefinition {
	return m.OutputTypes
}

func (m Map) GetPrompt() string {
	return m.Prompt
}

func (m Map) GetCodeComments() string {
	return m.CodeComments
}

func (m Map) GetDocumentation() string {
	return m.Documentation
}

func (m Map) GetFunctionCalls() []*FunctionDefinition {
	return m.FunctionCalls
}

func (m Map) GetShouldStub() bool {
	return m.ShouldStub
}

// A Filter unit operation removes, rejects, or reroutes data.
type Filter struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (f Filter) GetUnitOperationTypeDefinition() string {
	return "A Filter unit operation removes, rejects, or reroutes data."
}

func (f Filter) GetInputTypes() []*TypeDefinition {
	return f.InputTypes
}

func (f Filter) GetOutputTypes() []*TypeDefinition {
	return f.OutputTypes
}

func (f Filter) GetPrompt() string {
	return f.Prompt
}

func (f Filter) GetCodeComments() string {
	return f.CodeComments
}

func (f Filter) GetDocumentation() string {
	return f.Documentation
}

func (f Filter) GetFunctionCalls() []*FunctionDefinition {
	return f.FunctionCalls
}

func (f Filter) GetShouldStub() bool {
	return f.ShouldStub
}

// A Sort unit operation orders a collection according to defined rules.
type Sort struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (s Sort) GetUnitOperationTypeDefinition() string {
	return "A Sort unit operation orders a collection according to defined rules."
}

func (s Sort) GetInputTypes() []*TypeDefinition {
	return s.InputTypes
}

func (s Sort) GetOutputTypes() []*TypeDefinition {
	return s.OutputTypes
}

func (s Sort) GetPrompt() string {
	return s.Prompt
}

func (s Sort) GetCodeComments() string {
	return s.CodeComments
}

func (s Sort) GetDocumentation() string {
	return s.Documentation
}

func (s Sort) GetFunctionCalls() []*FunctionDefinition {
	return s.FunctionCalls
}

func (s Sort) GetShouldStub() bool {
	return s.ShouldStub
}

// A Distribution unit operation chooses where data goes next, such as branching.
type Distribution struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (d Distribution) GetUnitOperationTypeDefinition() string {
	return "A Distribution unit operation chooses where data goes next, such as branching."
}

func (d Distribution) GetInputTypes() []*TypeDefinition {
	return d.InputTypes
}

func (d Distribution) GetOutputTypes() []*TypeDefinition {
	return d.OutputTypes
}

func (d Distribution) GetPrompt() string {
	return d.Prompt
}

func (d Distribution) GetCodeComments() string {
	return d.CodeComments
}

func (d Distribution) GetDocumentation() string {
	return d.Documentation
}

func (d Distribution) GetFunctionCalls() []*FunctionDefinition {
	return d.FunctionCalls
}

func (d Distribution) GetShouldStub() bool {
	return d.ShouldStub
}

// A Validate unit operation checks data integrity.
type Validate struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (v Validate) GetUnitOperationTypeDefinition() string {
	return "A Validate unit operation checks data integrity."
}

func (v Validate) GetInputTypes() []*TypeDefinition {
	return v.InputTypes
}

func (v Validate) GetOutputTypes() []*TypeDefinition {
	return v.OutputTypes
}

func (v Validate) GetPrompt() string {
	return v.Prompt
}

func (v Validate) GetCodeComments() string {
	return v.CodeComments
}

func (v Validate) GetDocumentation() string {
	return v.Documentation
}

func (v Validate) GetFunctionCalls() []*FunctionDefinition {
	return v.FunctionCalls
}

func (v Validate) GetShouldStub() bool {
	return v.ShouldStub
}

// An Authenticate unit operation determines the identity initiating a flow.
type Authenticate struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (a Authenticate) GetUnitOperationTypeDefinition() string {
	return "An Authenticate unit operation determines the identity initiating a flow."
}

func (a Authenticate) GetInputTypes() []*TypeDefinition {
	return a.InputTypes
}

func (a Authenticate) GetOutputTypes() []*TypeDefinition {
	return a.OutputTypes
}

func (a Authenticate) GetPrompt() string {
	return a.Prompt
}

func (a Authenticate) GetCodeComments() string {
	return a.CodeComments
}

func (a Authenticate) GetDocumentation() string {
	return a.Documentation
}

func (a Authenticate) GetFunctionCalls() []*FunctionDefinition {
	return a.FunctionCalls
}

func (a Authenticate) GetShouldStub() bool {
	return a.ShouldStub
}

// An Authorize unit operation determines the whether an authenticated identity may perform an action.
type Authorize struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (a Authorize) GetUnitOperationTypeDefinition() string {
	return "An Authorize unit operation determines the whether an authenticated identity may perform an action."
}

func (a Authorize) GetInputTypes() []*TypeDefinition {
	return a.InputTypes
}

func (a Authorize) GetOutputTypes() []*TypeDefinition {
	return a.OutputTypes
}

func (a Authorize) GetPrompt() string {
	return a.Prompt
}

func (a Authorize) GetCodeComments() string {
	return a.CodeComments
}

func (a Authorize) GetDocumentation() string {
	return a.Documentation
}

func (a Authorize) GetFunctionCalls() []*FunctionDefinition {
	return a.FunctionCalls
}

func (a Authorize) GetShouldStub() bool {
	return a.ShouldStub
}

// A GlobalStateRead unit operation reads values whose lifetime extends beyond the current call stack
type GlobalStateRead struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (g GlobalStateRead) GetUnitOperationTypeDefinition() string {
	return "A GlobalStateRead unit operation reads values whose lifetime extends beyond the current call stack"
}

func (g GlobalStateRead) GetInputTypes() []*TypeDefinition {
	return g.InputTypes
}

func (g GlobalStateRead) GetOutputTypes() []*TypeDefinition {
	return g.OutputTypes
}

func (g GlobalStateRead) GetPrompt() string {
	return g.Prompt
}

func (g GlobalStateRead) GetCodeComments() string {
	return g.CodeComments
}

func (g GlobalStateRead) GetDocumentation() string {
	return g.Documentation
}

func (g GlobalStateRead) GetFunctionCalls() []*FunctionDefinition {
	return g.FunctionCalls
}

func (g GlobalStateRead) GetShouldStub() bool {
	return g.ShouldStub
}

// A GlobalStateWrite unit operation writes values whose lifetime extends beyond the current call stack
type GlobalStateWrite struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (g GlobalStateWrite) GetUnitOperationTypeDefinition() string {
	return "A GlobalStateWrite unit operation writes values whose lifetime extends beyond the current call stack"
}

func (g GlobalStateWrite) GetInputTypes() []*TypeDefinition {
	return g.InputTypes
}

func (g GlobalStateWrite) GetOutputTypes() []*TypeDefinition {
	return g.OutputTypes
}

func (g GlobalStateWrite) GetPrompt() string {
	return g.Prompt
}

func (g GlobalStateWrite) GetCodeComments() string {
	return g.CodeComments
}

func (g GlobalStateWrite) GetDocumentation() string {
	return g.Documentation
}

func (g GlobalStateWrite) GetFunctionCalls() []*FunctionDefinition {
	return g.FunctionCalls
}

func (g GlobalStateWrite) GetShouldStub() bool {
	return g.ShouldStub
}

// An InputOutput unit operation communicates outside the program, including HTTP, databases, etc...
type InputOutput struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (i InputOutput) GetUnitOperationTypeDefinition() string {
	return "An InputOutput unit operation communicates outside the program, including HTTP, databases, etc..."
}

func (i InputOutput) GetInputTypes() []*TypeDefinition {
	return i.InputTypes
}

func (i InputOutput) GetOutputTypes() []*TypeDefinition {
	return i.OutputTypes
}

func (i InputOutput) GetPrompt() string {
	return i.Prompt
}

func (i InputOutput) GetCodeComments() string {
	return i.CodeComments
}

func (i InputOutput) GetDocumentation() string {
	return i.Documentation
}

func (i InputOutput) GetFunctionCalls() []*FunctionDefinition {
	return i.FunctionCalls
}

func (i InputOutput) GetShouldStub() bool {
	return i.ShouldStub
}

// A Panic unit operation terminates execution of the program.
type Panic struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	CodeComments  string
	Documentation string
	FunctionCalls []*FunctionDefinition
	ShouldStub    bool
}

func (p Panic) GetUnitOperationTypeDefinition() string {
	return "A Panic unit operation terminates execution of the program."
}

func (p Panic) GetInputTypes() []*TypeDefinition {
	return p.InputTypes
}

func (p Panic) GetOutputTypes() []*TypeDefinition {
	return p.OutputTypes
}

func (p Panic) GetPrompt() string {
	return p.Prompt
}

func (p Panic) GetCodeComments() string {
	return p.CodeComments
}

func (p Panic) GetDocumentation() string {
	return p.Documentation
}

func (p Panic) GetFunctionCalls() []*FunctionDefinition {
	return p.FunctionCalls
}

func (p Panic) GetShouldStub() bool {
	return p.ShouldStub
}
