package v1

import (
	"github.com/ertgl/croncache/lib"
)

func Generator() lib.IoC {
	ioc := NewIoC()
	return ioc
}
