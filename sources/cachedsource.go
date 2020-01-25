package sources

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"

	evalexpr "github.com/bgokden/go-eval-expr"
)

type CachedSources struct {
	Cache *cache.Cache
}

func NewCachedSources(defaultExpration, cleanupInterval time.Duration) *CachedSources {
	return &CachedSources{
		Cache: cache.New(defaultExpration, cleanupInterval),
	}
}

func (cs *CachedSources) GetSource(reference string) (evalexpr.Source, error) {
	if cs.Cache != nil {
		if sourceInterface, exists := cs.Cache.Get(reference); exists {
			if source, ok := sourceInterface.(evalexpr.Source); ok {
				return source, nil
			}
		}
	}
	return nil, fmt.Errorf("Source %v does not exist", reference)
}

func (cs *CachedSources) SetSource(reference string, source evalexpr.Source, expirationPeriod time.Duration) error {
	if cs.Cache != nil {
		cs.Cache.Set(reference, source, expirationPeriod)
		return nil
	}
	return fmt.Errorf("Cache is not initialized")
}

func (cs *CachedSources) SetValue(reference string, value interface{}) error {
	source := &MemorySource{
		Value: value,
	}
	return cs.SetSource(reference, source, cache.DefaultExpiration)
}
