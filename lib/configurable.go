package lib

type Configurable interface {
	ImportConfig([]byte) error
	ExportConfig() ([]byte, error)
}
