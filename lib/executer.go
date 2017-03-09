package lib

import (
	"github.com/ertgl/croncache/models"
)

type Executer interface {
	Dependency
	Execute(command string, args ...string) (models.Cache, error)
}
