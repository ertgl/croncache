package lib

type NodeRepository interface {
	Register(moduleName string, generator NodeGenerator) error
	Resolve(moduleName string) (NodeGenerator, error)
	Map() map[string]NodeGenerator
}
