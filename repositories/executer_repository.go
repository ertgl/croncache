package repositories

import (
	"errors"
)

import (
	"github.com/ertgl/croncache/lib"
)

var (
	ExecuterTypeAlreadyExist error = errors.New("Executer type already exists")
	ExecuterTypeDoesNotExist error = errors.New("Executer type does not exist")
)

type ExecuterRepository struct {
	lib.ExecuterRepository
	storage map[string]lib.ExecuterGenerator
}

func NewExecuterRepository() lib.ExecuterRepository {
	r := &ExecuterRepository{
		storage: make(map[string]lib.ExecuterGenerator, 0),
	}
	return r
}

func (r *ExecuterRepository) get(moduleName string) (lib.ExecuterGenerator, bool) {
	if g, ok := r.storage[moduleName]; ok {
		return g, true
	}
	return nil, false
}

func (r *ExecuterRepository) Register(moduleName string, g lib.ExecuterGenerator) error {
	if _, ok := r.get(moduleName); ok {
		return ExecuterTypeAlreadyExist
	}
	r.storage[moduleName] = g
	return nil
}

func (r *ExecuterRepository) Resolve(moduleName string) (lib.ExecuterGenerator, error) {
	if g, ok := r.get(moduleName); ok {
		return g, nil
	}
	return nil, ExecuterTypeDoesNotExist
}

func (r *ExecuterRepository) Map() map[string]lib.ExecuterGenerator {
	return r.storage
}
