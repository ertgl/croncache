package repositories

import (
	"errors"
)

import (
	"github.com/ertgl/croncache/lib"
)

var (
	CacheEngineTypeAlreadyExist error = errors.New("Cache engine type already exists")
	CacheEngineTypeDoesNotExist error = errors.New("Cache engine type does not exist")
)

type CacheEngineRepository struct {
	lib.CacheEngineRepository
	storage map[string]lib.CacheEngineGenerator
}

func NewCacheEngineRepository() lib.CacheEngineRepository {
	r := &CacheEngineRepository{
		storage: make(map[string]lib.CacheEngineGenerator, 0),
	}
	return r
}

func (r *CacheEngineRepository) get(moduleName string) (lib.CacheEngineGenerator, bool) {
	if g, ok := r.storage[moduleName]; ok {
		return g, true
	}
	return nil, false
}

func (r *CacheEngineRepository) Register(moduleName string, g lib.CacheEngineGenerator) error {
	if _, ok := r.get(moduleName); ok {
		return CacheEngineTypeAlreadyExist
	}
	r.storage[moduleName] = g
	return nil
}

func (r *CacheEngineRepository) Resolve(moduleName string) (lib.CacheEngineGenerator, error) {
	if g, ok := r.get(moduleName); ok {
		return g, nil
	}
	return nil, CacheEngineTypeDoesNotExist
}

func (r *CacheEngineRepository) Map() map[string]lib.CacheEngineGenerator {
	return r.storage
}
