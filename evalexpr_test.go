package evalexpr_test

import (
	"testing"
	"time"

	evalexpr "github.com/bgokden/go-eval-expr"
	sources "github.com/bgokden/go-eval-expr/sources"
	"github.com/stretchr/testify/assert"
)

func TestEvalWithSources(t *testing.T) {
	expression := "a > 0"

	memorySources := sources.NewMemorySources()
	memorySources.SetValue("a", 3)

	output, env, err := evalexpr.EvalWithSources(expression, memorySources)
	assert.Nil(t, err)
	assert.Equal(t, true, output)
	assert.Equal(t, 3, env["a"])
}

func TestEvaluatorWithPrometheus(t *testing.T) {
	cachedSources := sources.NewCachedSources(-1, 10*time.Minute)

	limitValue := float64(1500)

	evaluator := &evalexpr.Evaluator{
		Expression: "value_from_prometheus > comparison_value",
		Sources:    cachedSources,
	}

	cachedSources.SetSource("value_from_prometheus", sources.NewPrometheusSource(
		"http://demo.robustperception.io:9090",
		"rate(prometheus_tsdb_head_samples_appended_total[5m])",
	), 0)

	cachedSources.SetSource("comparison_value", sources.NewMemorySource(
		limitValue,
	), 0)

	output, env, err := evaluator.Eval()
	assert.Nil(t, err)
	assert.Equal(t, true, output)
	assert.NotNil(t, env["value_from_prometheus"])
	value_from_prometheus, ok := env["value_from_prometheus"].(float64)
	assert.True(t, ok)
	assert.True(t, value_from_prometheus > limitValue)
	assert.NotNil(t, env["comparison_value"])
	comparison_value, ok := env["comparison_value"].(float64)
	assert.True(t, ok)
	assert.Equal(t, limitValue, comparison_value)
}
