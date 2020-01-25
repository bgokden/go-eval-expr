package sources

import (
	"fmt"

	evalexpr "github.com/bgokden/go-eval-expr"
)

type MemorySources struct {
	SourceMap map[string]evalexpr.Source
}

func NewMemorySources() *MemorySources {
	return &MemorySources{
		SourceMap: make(map[string]evalexpr.Source),
	}
}

func (ms *MemorySources) GetSource(reference string) (evalexpr.Source, error) {
	if ms.SourceMap != nil {
		if source, ok := ms.SourceMap[reference]; ok {
			return source, nil
		}
	}
	return nil, fmt.Errorf("Source %v does not exist", reference)
}

func (ms *MemorySources) SetSource(reference string, source evalexpr.Source) error {
	if ms.SourceMap != nil {
		ms.SourceMap[reference] = source
		return nil
	}
	return fmt.Errorf("SourceMap is not initialized")
}

func (ms *MemorySources) SetValue(reference string, value interface{}) error {
	source := &MemorySource{
		Value: value,
	}
	return ms.SetSource(reference, source)
}

func NewMemorySource(value interface{}) *MemorySource {
	return &MemorySource{
		Value: value,
	}
}

type MemorySource struct {
	Value interface{}
}

func (ms *MemorySource) GetValue() (interface{}, error) {
	return ms.Value, nil
}
