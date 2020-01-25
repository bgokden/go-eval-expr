package sources_test

import (
	"fmt"
	"testing"

	"github.com/bgokden/go-eval-expr/sources"
	"github.com/stretchr/testify/assert"
)

func TestPrometheusSource(t *testing.T) {
	prometheusSource := sources.NewPrometheusSource(
		"http://demo.robustperception.io:9090",
		"rate(prometheus_tsdb_head_samples_appended_total[5m])",
	)

	value, err := prometheusSource.GetValue()
	assert.Nil(t, err)
	fmt.Printf("result: %v\n", value)
}
