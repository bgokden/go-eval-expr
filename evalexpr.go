package evalexpr

import (
	"fmt"

	"github.com/antonmedv/expr"
)

type Source interface {
	GetValue() (interface{}, error)
}

func EvalWithSources(expression string, sources map[string]Source) (interface{}, map[string]interface{}, error) {
	// Parse
	program, err := expr.Parse(expression)
	if err != nil {
		return nil, nil, err
	}
	// Get Resources
	env := make(map[string]interface{}, len(program.Constants))
	for _, constant := range program.Constants {
		// fmt.Printf("%v : %v\n", v, reflect.TypeOf(v))
		if constantName, ok := constant.(string); ok {
			if source, exists := sources[constantName]; exists {
				env[constantName], _ = source.GetValue()
			} else {
				return nil, nil, fmt.Errorf("Undefined Source: %v", constantName)
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
