package lib

import (
	"github.com/ertgl/croncache/models"
)

type Executer interface {
	Dependency
	Execute(command string, timeout models.Duration, args ...string) (models.Cache, error)
}
