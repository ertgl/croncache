package lib

type ExecuterRepository interface {
	Register(moduleName string, generator ExecuterGenerator) error
	Resolve(moduleName string) (ExecuterGenerator, error)
	Map() map[string]ExecuterGenerator
}
