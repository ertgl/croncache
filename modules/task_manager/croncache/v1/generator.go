package v1

import (
	"github.com/ertgl/croncache/lib"
)

func Generator() lib.TaskManager {
	tm := NewTaskManager()
	return tm
}
