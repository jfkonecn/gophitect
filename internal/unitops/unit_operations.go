package unitops

import (
	"errors"
	"fmt"
	"strings"
)

// A Map unit operation converts one value or type into another.
type Map struct {
	Input         *VariableDefinition
	Output        *VariableDefinition
	FunctionCalls []*FunctionDefinition
	CodeComments  string
}

func (m Map) GetInputTypes() []*VariableDefinition {
	return []*VariableDefinition{m.Input}
}

func (m Map) GetOutputTypes() []*VariableDefinition {
	return []*VariableDefinition{m.Output}
}

func (m Map) GetPrompt() (string, error) {
	var sb strings.Builder

	if m.Input == nil {
		return "", errors.New("map's input is nil")
	}
	input := *m.Input

	if input.GetTypeDefinition() == nil {
		return "", errors.New("map's input type definition is nil")
	}
	inputTypeDefinition := *input.GetTypeDefinition()

	if m.Output == nil {
		return "", errors.New("map's output is nil")
	}
	output := *m.Output

	if input.GetTypeDefinition() == nil {
		return "", errors.New("map's output type definition is nil")
	}
	outputTypeDefinition := *output.GetTypeDefinition()

	_, err := fmt.Fprintf(&sb, "Map %s of type %s to %s of type %s\n",
		input.GetVariableName(),
		inputTypeDefinition.GetTypeName(),
		output.GetVariableName(),
		outputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	if m.CodeComments != "" {
		_, err := fmt.Fprintf(&sb, "Add these comments %s\n", m.CodeComments)
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (m Map) GetFunctionCalls() []*FunctionDefinition {
	return m.FunctionCalls
}

// A Filter unit operation removes, rejects, or reroutes data.
type Filter struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (f Filter) GetFunctionCalls() []*FunctionDefinition {
	return f.FunctionCalls
}

// A Sort unit operation orders a collection according to defined rules.
type Sort struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (s Sort) GetFunctionCalls() []*FunctionDefinition {
	return s.FunctionCalls
}

// A Distribution unit operation chooses where data goes next, such as branching.
type Distribution struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (d Distribution) GetFunctionCalls() []*FunctionDefinition {
	return d.FunctionCalls
}

// A Validate unit operation checks data integrity.
type Validate struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (v Validate) GetFunctionCalls() []*FunctionDefinition {
	return v.FunctionCalls
}

// An Authenticate unit operation determines the identity initiating a flow.
type Authenticate struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (a Authenticate) GetFunctionCalls() []*FunctionDefinition {
	return a.FunctionCalls
}

// An Authorize unit operation determines whether an authenticated identity may perform an action.
type Authorize struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (a Authorize) GetFunctionCalls() []*FunctionDefinition {
	return a.FunctionCalls
}

// A GlobalStateRead unit operation reads values whose lifetime extends beyond the current call stack.
type GlobalStateRead struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (g GlobalStateRead) GetFunctionCalls() []*FunctionDefinition {
	return g.FunctionCalls
}

// A GlobalStateWrite unit operation writes values whose lifetime extends beyond the current call stack.
type GlobalStateWrite struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (g GlobalStateWrite) GetFunctionCalls() []*FunctionDefinition {
	return g.FunctionCalls
}

// An InputOutput unit operation communicates outside the program, including HTTP, databases, etc.
type InputOutput struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (i InputOutput) GetFunctionCalls() []*FunctionDefinition {
	return i.FunctionCalls
}

// A Panic unit operation terminates execution of the program.
type Panic struct {
	InputTypes    []*TypeDefinition
	OutputTypes   []*TypeDefinition
	Prompt        string
	FunctionCalls []*FunctionDefinition
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

func (p Panic) GetFunctionCalls() []*FunctionDefinition {
	return p.FunctionCalls
}
