package repositories

import (
	"errors"
)

import (
	"github.com/ertgl/croncache/lib"
)

var (
	NodeTypeAlreadyExist error = errors.New("Node type already exists")
	NodeTypeDoesNotExist error = errors.New("Node type does not exist")
)

type NodeRepository struct {
	lib.NodeRepository
	storage map[string]lib.NodeGenerator
}

func NewNodeRepository() lib.NodeRepository {
	r := &NodeRepository{
		storage: make(map[string]lib.NodeGenerator, 0),
	}
	return r
}

func (r *NodeRepository) get(moduleName string) (lib.NodeGenerator, bool) {
	if g, ok := r.storage[moduleName]; ok {
		return g, true
	}
	return nil, false
}

func (r *NodeRepository) Register(moduleName string, g lib.NodeGenerator) error {
	if _, ok := r.get(moduleName); ok {
		return NodeTypeAlreadyExist
	}
	r.storage[moduleName] = g
	return nil
}

func (r *NodeRepository) Resolve(moduleName string) (lib.NodeGenerator, error) {
	if g, ok := r.get(moduleName); ok {
		return g, nil
	}
	return nil, NodeTypeDoesNotExist
}

func (r *NodeRepository) Map() map[string]lib.NodeGenerator {
	return r.storage
}
