package v1

import (
	"github.com/ertgl/croncache"
)

func init() {
	err := croncache.NodeRepository().Register(MODULE_NAME, Generator)
	if err != nil {
		croncache.HandleFatalError(err)
	}
}
