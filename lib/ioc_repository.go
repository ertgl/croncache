package lib

type IoCRepository interface {
	Register(moduleName string, generator IoCGenerator) error
	Resolve(moduleName string) (IoCGenerator, error)
	Map() map[string]IoCGenerator
}
