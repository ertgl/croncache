package lib

import (
	"github.com/ertgl/croncache/models"
)

type TaskManager interface {
	Dependency
	Configurable
	Process(*models.Task)
	Run()
}
