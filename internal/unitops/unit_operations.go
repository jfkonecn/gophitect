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

	if output.GetTypeDefinition() == nil {
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
		_, err := fmt.Fprintf(&sb, "- Add these comments\n%s\n", m.CodeComments)
		if err != nil {
			return "", err
		}
	}

	if len(m.FunctionCalls) > 0 {
		_, err := fmt.Fprintln(&sb, "Make sure you use these functions in the map logic")
		if err != nil {
			return "", err
		}
	}

	for _, functionCall := range m.FunctionCalls {
		if functionCall == nil {
			return "", errors.New("function call in map is nil")
		}

		_, err := fmt.Fprintf(&sb, "- %s\n", (*functionCall).GetFunctionName())
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
	Input         *CollectionVariableDefinition
	Output        *CollectionVariableDefinition
	FilterLogic   string
	FunctionCalls []*FunctionDefinition
	CodeComments  string
}

func (f Filter) GetInputTypes() []*VariableDefinition {
	if f.Input == nil {
		return []*VariableDefinition{nil}
	}

	var input VariableDefinition = f.Input
	return []*VariableDefinition{&input}
}

func (f Filter) GetOutputTypes() []*VariableDefinition {
	if f.Output == nil {
		return []*VariableDefinition{nil}
	}

	var output VariableDefinition = f.Output
	return []*VariableDefinition{&output}
}

func (f Filter) GetPrompt() (string, error) {
	var sb strings.Builder

	if f.Input == nil {
		return "", errors.New("filter's input is nil")
	}
	input := *f.Input

	if input.GetTypeDefinition() == nil {
		return "", errors.New("filter's input type definition is nil")
	}
	inputTypeDefinition := *input.GetTypeDefinition()

	if f.Output == nil {
		return "", errors.New("filter's output is nil")
	}
	output := *f.Output

	if output.GetTypeDefinition() == nil {
		return "", errors.New("filter's output type definition is nil")
	}
	outputTypeDefinition := *output.GetTypeDefinition()

	_, err := fmt.Fprintf(&sb, "filter %s of type %s to %s of type %s\n",
		input.GetVariableName(),
		inputTypeDefinition.GetTypeName(),
		output.GetVariableName(),
		outputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "Do this logic to filter\n%s\n", f.FilterLogic)
	if err != nil {
		return "", err
	}

	if f.CodeComments != "" {
		_, err := fmt.Fprintf(&sb, "- Add these comments\n%s\n", f.CodeComments)
		if err != nil {
			return "", err
		}
	}

	if len(f.FunctionCalls) > 0 {
		_, err := fmt.Fprintln(&sb, "Make sure you use these functions in the filter logic")
		if err != nil {
			return "", err
		}
	}

	for _, functionCall := range f.FunctionCalls {
		if functionCall == nil {
			return "", errors.New("function call in filter is nil")
		}

		_, err := fmt.Fprintf(&sb, "- %s\n", (*functionCall).GetFunctionName())
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (f Filter) GetFunctionCalls() []*FunctionDefinition {
	return f.FunctionCalls
}

// A Sort unit operation orders a collection according to defined rules.
type Sort struct {
	Input         *VariableDefinition
	FunctionCalls []*FunctionDefinition
	SortLogic     string
	CodeComments  string
}

func (s Sort) GetInputTypes() []*VariableDefinition {
	return []*VariableDefinition{s.Input}
}

func (s Sort) GetOutputTypes() []*VariableDefinition {
	return nil
}

func (s Sort) GetPrompt() (string, error) {
	var sb strings.Builder

	if s.Input == nil {
		return "", errors.New("sort's input is nil")
	}
	input := *s.Input

	if input.GetTypeDefinition() == nil {
		return "", errors.New("sort's input type definition is nil")
	}
	inputTypeDefinition := *input.GetTypeDefinition()

	_, err := fmt.Fprintf(&sb, "sort %s of type %s\n",
		input.GetVariableName(),
		inputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "Do this logic to sort\n%s\n", s.SortLogic)
	if err != nil {
		return "", err
	}

	if s.CodeComments != "" {
		_, err := fmt.Fprintf(&sb, "- Add these comments\n%s\n", s.CodeComments)
		if err != nil {
			return "", err
		}
	}

	if len(s.FunctionCalls) > 0 {
		_, err := fmt.Fprintln(&sb, "Make sure you use these functions in the sort logic")
		if err != nil {
			return "", err
		}
	}

	for _, functionCall := range s.FunctionCalls {
		if functionCall == nil {
			return "", errors.New("function call in sort is nil")
		}

		_, err := fmt.Fprintf(&sb, "- %s\n", (*functionCall).GetFunctionName())
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (s Sort) GetFunctionCalls() []*FunctionDefinition {
	return s.FunctionCalls
}

type DistributionCondition struct {
	Condition     string
	Output        *VariableDefinition
	FunctionCalls []*FunctionDefinition
}

// A Distribution unit operation chooses where data goes next, such as branching.
type Distribution struct {
	Input      *VariableDefinition
	Conditions []DistributionCondition
}

func (d Distribution) GetInputTypes() []*VariableDefinition {
	return []*VariableDefinition{d.Input}
}

func (d Distribution) GetOutputTypes() []*VariableDefinition {
	var outputs []*VariableDefinition
	for _, c := range d.Conditions {
		outputs = append(outputs, c.Output)
	}
	return outputs
}

func (d Distribution) GetPrompt() (string, error) {
	var sb strings.Builder

	if d.Input == nil {
		return "", errors.New("distribution's input is nil")
	}
	input := *d.Input

	if input.GetTypeDefinition() == nil {
		return "", errors.New("distribution's input type definition is nil")
	}
	inputTypeDefinition := *input.GetTypeDefinition()

	if len(d.Conditions) == 0 {
		return "", errors.New("distribution's conditions are empty")
	}

	_, err := fmt.Fprintf(&sb, "distribute %s of type %s based on these conditions\n",
		input.GetVariableName(),
		inputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintln(&sb, "Evaluate each condition and route the input to the matching output.")
	if err != nil {
		return "", err
	}

	for i, condition := range d.Conditions {
		if condition.Condition == "" {
			return "", errors.New("distribution condition is empty")
		}

		if condition.Output == nil {
			return "", errors.New("distribution condition's output is nil")
		}
		output := *condition.Output

		if output.GetTypeDefinition() == nil {
			return "", errors.New("distribution condition's output type definition is nil")
		}
		outputTypeDefinition := *output.GetTypeDefinition()

		_, err := fmt.Fprintf(&sb, "%d. When %s\n", i+1, condition.Condition)
		if err != nil {
			return "", err
		}

		_, err = fmt.Fprintf(&sb, "- Route to %s of type %s\n",
			output.GetVariableName(),
			outputTypeDefinition.GetTypeName())
		if err != nil {
			return "", err
		}

		if len(condition.FunctionCalls) > 0 {
			_, err := fmt.Fprintln(&sb, "- Make sure you use these functions for this condition")
			if err != nil {
				return "", err
			}
		}

		for _, functionCall := range condition.FunctionCalls {
			if functionCall == nil {
				return "", errors.New("function call in distribution condition is nil")
			}

			_, err := fmt.Fprintf(&sb, "  - %s\n", (*functionCall).GetFunctionName())
			if err != nil {
				return "", err
			}
		}
	}

	return sb.String(), nil
}

func (d Distribution) GetFunctionCalls() []*FunctionDefinition {
	var functionCalls []*FunctionDefinition
	for _, c := range d.Conditions {
		functionCalls = append(functionCalls, c.FunctionCalls...)
	}
	return functionCalls
}

// A Validate unit operation checks data integrity.
type Validate struct {
	Input         *VariableDefinition
	SuccessOutput *VariableDefinition
	FailureOutput *VariableDefinition
	FunctionCalls []*FunctionDefinition
}

func (v Validate) GetInputTypes() []*VariableDefinition {
	return []*VariableDefinition{v.Input}
}

func (v Validate) GetOutputTypes() []*VariableDefinition {
	return []*VariableDefinition{v.SuccessOutput, v.FailureOutput}
}

func (v Validate) GetPrompt() (string, error) {
	var sb strings.Builder

	if v.Input == nil {
		return "", errors.New("validate's input is nil")
	}
	input := *v.Input

	if input.GetTypeDefinition() == nil {
		return "", errors.New("validate's input type definition is nil")
	}
	inputTypeDefinition := *input.GetTypeDefinition()

	if v.SuccessOutput == nil {
		return "", errors.New("validate's success output is nil")
	}
	successOutput := *v.SuccessOutput

	if successOutput.GetTypeDefinition() == nil {
		return "", errors.New("validate's success output type definition is nil")
	}
	successOutputTypeDefinition := *successOutput.GetTypeDefinition()

	if v.FailureOutput == nil {
		return "", errors.New("validate's failure output is nil")
	}
	failureOutput := *v.FailureOutput

	if failureOutput.GetTypeDefinition() == nil {
		return "", errors.New("validate's failure output type definition is nil")
	}
	failureOutputTypeDefinition := *failureOutput.GetTypeDefinition()

	_, err := fmt.Fprintf(&sb, "validate %s of type %s\n",
		input.GetVariableName(),
		inputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintln(&sb, "Route the input based on whether validation succeeds or fails.")
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "- On success, route to %s of type %s\n",
		successOutput.GetVariableName(),
		successOutputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "- On failure, route to %s of type %s\n",
		failureOutput.GetVariableName(),
		failureOutputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	if len(v.FunctionCalls) > 0 {
		_, err := fmt.Fprintln(&sb, "Make sure you use these functions in the validate logic")
		if err != nil {
			return "", err
		}
	}

	for _, functionCall := range v.FunctionCalls {
		if functionCall == nil {
			return "", errors.New("function call in validate is nil")
		}

		_, err := fmt.Fprintf(&sb, "- %s\n", (*functionCall).GetFunctionName())
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (v Validate) GetFunctionCalls() []*FunctionDefinition {
	return v.FunctionCalls
}

// An Authenticate unit operation determines the identity initiating a flow.
type Authenticate struct {
	Input         *VariableDefinition
	SuccessOutput *VariableDefinition
	FailureOutput *VariableDefinition
	FunctionCalls []*FunctionDefinition
}

func (a Authenticate) GetInputTypes() []*VariableDefinition {
	return []*VariableDefinition{a.Input}
}

func (a Authenticate) GetOutputTypes() []*VariableDefinition {
	return []*VariableDefinition{a.SuccessOutput, a.FailureOutput}
}

func (a Authenticate) GetPrompt() (string, error) {
	var sb strings.Builder

	if a.Input == nil {
		return "", errors.New("authenticate's input is nil")
	}
	input := *a.Input

	if input.GetTypeDefinition() == nil {
		return "", errors.New("authenticate's input type definition is nil")
	}
	inputTypeDefinition := *input.GetTypeDefinition()

	if a.SuccessOutput == nil {
		return "", errors.New("authenticate's success output is nil")
	}
	successOutput := *a.SuccessOutput

	if successOutput.GetTypeDefinition() == nil {
		return "", errors.New("authenticate's success output type definition is nil")
	}
	successOutputTypeDefinition := *successOutput.GetTypeDefinition()

	if a.FailureOutput == nil {
		return "", errors.New("authenticate's failure output is nil")
	}
	failureOutput := *a.FailureOutput

	if failureOutput.GetTypeDefinition() == nil {
		return "", errors.New("authenticate's failure output type definition is nil")
	}
	failureOutputTypeDefinition := *failureOutput.GetTypeDefinition()

	_, err := fmt.Fprintf(&sb, "authenticate %s of type %s\n",
		input.GetVariableName(),
		inputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintln(&sb, "Route the input based on whether authentication succeeds or fails.")
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "- On success, route to %s of type %s\n",
		successOutput.GetVariableName(),
		successOutputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(&sb, "- On failure, route to %s of type %s\n",
		failureOutput.GetVariableName(),
		failureOutputTypeDefinition.GetTypeName())
	if err != nil {
		return "", err
	}

	if len(a.FunctionCalls) > 0 {
		_, err := fmt.Fprintln(&sb, "Make sure you use these functions in the authenticate logic")
		if err != nil {
			return "", err
		}
	}

	for _, functionCall := range a.FunctionCalls {
		if functionCall == nil {
			return "", errors.New("function call in authenticate is nil")
		}

		_, err := fmt.Fprintf(&sb, "- %s\n", (*functionCall).GetFunctionName())
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
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
