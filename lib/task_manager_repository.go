package lib

type TaskManagerRepository interface {
	Register(moduleName string, generator TaskManagerGenerator) error
	Resolve(moduleName string) (TaskManagerGenerator, error)
	Map() map[string]TaskManagerGenerator
}
