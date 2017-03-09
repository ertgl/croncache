package lib

type CacheEngineRepository interface {
	Register(moduleName string, generator CacheEngineGenerator) error
	Resolve(moduleName string) (CacheEngineGenerator, error)
	Map() map[string]CacheEngineGenerator
}
