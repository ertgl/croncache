package lib

type Dependency interface {
	ModuleName() string
	Initialize() error
	IoC() IoC
	SetIoC(IoC)
}
