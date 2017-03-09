package repositories

import (
	"errors"
)

import (
	"github.com/ertgl/croncache/lib"
)

var (
	IoCTypeAlreadyExist error = errors.New("IoC type already exists")
	IoCTypeDoesNotExist error = errors.New("IoC type does not exist")
)

type IoCRepository struct {
	lib.IoCRepository
	storage map[string]lib.IoCGenerator
}

func NewIoCRepository() lib.IoCRepository {
	r := &IoCRepository{
		storage: make(map[string]lib.IoCGenerator, 0),
	}
	return r
}

func (r *IoCRepository) get(moduleName string) (lib.IoCGenerator, bool) {
	if g, ok := r.storage[moduleName]; ok {
		return g, true
	}
	return nil, false
}

func (r *IoCRepository) Register(moduleName string, g lib.IoCGenerator) error {
	if _, ok := r.get(moduleName); ok {
		return IoCTypeAlreadyExist
	}
	r.storage[moduleName] = g
	return nil
}

func (r *IoCRepository) Resolve(moduleName string) (lib.IoCGenerator, error) {
	if g, ok := r.get(moduleName); ok {
		return g, nil
	}
	return nil, IoCTypeDoesNotExist
}

func (r *IoCRepository) Map() map[string]lib.IoCGenerator {
	return r.storage
}
