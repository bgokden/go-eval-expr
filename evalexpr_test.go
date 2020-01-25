package evalexpr_test

import (
	"testing"

	evalexpr "github.com/bgokden/go-eval-expr"
	"github.com/stretchr/testify/assert"
)

type MemorySource struct {
	Value interface{}
}

func (ms *MemorySource) GetValue() (interface{}, error) {
	return ms.Value, nil
}

func TestEvalWithSources(t *testing.T) {
	expression0 := "a > 0"

	sources := map[string]evalexpr.Source{
		"a": &MemorySource{
			Value: 3,
		},
	}

	output, env, err := evalexpr.EvalWithSources(expression0, sources)
	assert.Nil(t, err)
	assert.Equal(t, true, output)
	assert.Equal(t, 3, env["a"])
}
