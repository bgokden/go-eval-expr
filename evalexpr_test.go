package evalexpr_test

import (
	"fmt"
	"testing"

	evalexpr "github.com/bgokden/go-eval-expr"
	"github.com/stretchr/testify/assert"
)

type MemorySources struct {
	SourceMap map[string]evalexpr.Source
}

func (ms *MemorySources) GetSource(reference string) (evalexpr.Source, error) {
	if ms.SourceMap != nil {
		if source, ok := ms.SourceMap[reference]; ok {
			return source, nil
		}
	}
	return nil, fmt.Errorf("Source %v does not exist", reference)
}

type MemorySource struct {
	Value interface{}
}

func (ms *MemorySource) GetValue() (interface{}, error) {
	return ms.Value, nil
}

func TestEvalWithSources(t *testing.T) {
	expression0 := "a > 0"

	sources := &MemorySources{
		SourceMap: map[string]evalexpr.Source{
			"a": &MemorySource{
				Value: 3,
			},
		},
	}

	output, env, err := evalexpr.EvalWithSources(expression0, sources)
	assert.Nil(t, err)
	assert.Equal(t, true, output)
	assert.Equal(t, 3, env["a"])
}
