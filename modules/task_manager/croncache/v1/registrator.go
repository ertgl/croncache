package v1

import (
	"github.com/ertgl/croncache"
)

func init() {
	err := croncache.TaskManagerRepository().Register(MODULE_NAME, Generator)
	if err != nil {
		croncache.HandleFatalError(err)
	}
}
