# go-eval-expr
A go library to evaluate expressions backed by data


Go Eval Expr allows you to define different sources to evaluate expressions:

This example shows where you define one value from prometheus while another from memory:

```go
// -1 means no expiration period
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
// output is true
assert.Nil(t, err)
assert.Equal(t, true, output)
```
