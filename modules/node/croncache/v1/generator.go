package v1

import (
	"github.com/ertgl/croncache/lib"
)

func Generator() lib.Node {
	n := NewNode()
	return n
}
