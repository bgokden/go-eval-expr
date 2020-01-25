package sources

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type PrometheusSource struct {
	Address string
	Query   string
}

func NewPrometheusSource(address, query string) *PrometheusSource {
	return &PrometheusSource{
		Address: address,
		Query:   query,
	}
}

func (ps *PrometheusSource) GetValue() (interface{}, error) {
	return query(ps.Address, ps.Query)
}

func query(address, query string) (float64, error) {
	client, err := api.NewClient(api.Config{
		Address: address,
	})
	if err != nil {
		return 0, err
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, query, time.Now())
	if err != nil {
		return 0, err
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}

	switch result.Type() {
	case model.ValVector:
		if vector, ok := result.(model.Vector); ok {
			return float64(vector[0].Value), nil
		}
	case model.ValScalar:
		if scalar, ok := result.(*model.Scalar); ok {
			return float64(scalar.Value), nil
		}
	}

	if scalar, ok := result.(*model.Scalar); ok {
		return float64(scalar.Value), nil
	}

	return 0, fmt.Errorf("Error getting result")
}
