package evalexpr

import (
	"fmt"

	"github.com/antonmedv/expr"
)

type Sources interface {
	GetSource(string) (Source, error)
}

type Source interface {
	GetValue() (interface{}, error)
}

type Evaluator struct {
	Expression string
	Sources    Sources
}

func (e *Evaluator) Eval() (interface{}, map[string]interface{}, error) {
	if e.Expression != "" && e.Sources != nil {
		return EvalWithSources(e.Expression, e.Sources)
	}
	return nil, nil, fmt.Errorf("Evaluator not initialized yet")
}

func EvalWithSources(expression string, sources Sources) (interface{}, map[string]interface{}, error) {
	// Parse
	program, err := expr.Parse(expression)
	if err != nil {
		return nil, nil, err
	}
	// Get Resources
	env := make(map[string]interface{}, len(program.Constants))
	for _, constant := range program.Constants {
		if constantReference, ok := constant.(string); ok {
			if source, sourceErr := sources.GetSource(constantReference); sourceErr == nil {
				tempValue, valueErr := source.GetValue()
				if valueErr != nil {
					return nil, nil, valueErr
				}
				env[constantReference] = tempValue
			} else {
				return nil, nil, sourceErr
			}
		}
	}

	// Evaluate resources
	output, err := expr.Run(program, env)
	if err != nil {
		return nil, nil, err
	}

	return output, env, nil
}
