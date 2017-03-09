package v1

import (
	"github.com/ertgl/croncache/lib"
)

func Generator() lib.Executer {
	e := NewExecuter()
	return e
}
