package repositories

import (
	"errors"
)

import (
	"github.com/ertgl/croncache/lib"
)

var (
	TaskManagerTypeAlreadyExist error = errors.New("Task manager type already exists")
	TaskManagerTypeDoesNotExist error = errors.New("Task manager type does not exist")
)

type TaskManagerRepository struct {
	lib.TaskManagerRepository
	storage map[string]lib.TaskManagerGenerator
}

func NewTaskManagerRepository() lib.TaskManagerRepository {
	r := &TaskManagerRepository{
		storage: make(map[string]lib.TaskManagerGenerator, 0),
	}
	return r
}

func (r *TaskManagerRepository) get(moduleName string) (lib.TaskManagerGenerator, bool) {
	if g, ok := r.storage[moduleName]; ok {
		return g, true
	}
	return nil, false
}

func (r *TaskManagerRepository) Register(moduleName string, g lib.TaskManagerGenerator) error {
	if _, ok := r.get(moduleName); ok {
		return TaskManagerTypeAlreadyExist
	}
	r.storage[moduleName] = g
	return nil
}

func (r *TaskManagerRepository) Resolve(moduleName string) (lib.TaskManagerGenerator, error) {
	if g, ok := r.get(moduleName); ok {
		return g, nil
	}
	return nil, TaskManagerTypeDoesNotExist
}

func (r *TaskManagerRepository) Map() map[string]lib.TaskManagerGenerator {
	return r.storage
}
