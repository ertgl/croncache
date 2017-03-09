package v1

import (
	"github.com/ertgl/croncache/lib"
)

type IoC struct {
	lib.IoC
	moduleName  string
	node        lib.Node
	executer    lib.Executer
	taskManager lib.TaskManager
}

func NewIoC() lib.IoC {
	ioc := &IoC{
		moduleName: MODULE_NAME,
	}
	return ioc
}

func (ioc *IoC) ModuleName() string {
	return ioc.moduleName
}

func (ioc *IoC) Initialize() error {
	var err error = nil
	return err
}

func (ioc *IoC) Node() lib.Node {
	return ioc.node
}

func (ioc *IoC) SetNode(node lib.Node) {
	ioc.node = node
	if ioc.Node().IoC() != ioc {
		ioc.Node().SetIoC(ioc)
	}
}

func (ioc *IoC) Executer() lib.Executer {
	return ioc.executer
}

func (ioc *IoC) SetExecuter(executer lib.Executer) {
	ioc.executer = executer
	if ioc.Executer().IoC() != ioc {
		ioc.Executer().SetIoC(ioc)
	}
}

func (ioc *IoC) TaskManager() lib.TaskManager {
	return ioc.taskManager
}

func (ioc *IoC) SetTaskManager(taskManager lib.TaskManager) {
	ioc.taskManager = taskManager
	if ioc.TaskManager().IoC() != ioc {
		ioc.TaskManager().SetIoC(ioc)
	}
}
