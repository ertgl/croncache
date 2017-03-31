package v1

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"time"
)

import (
	"github.com/ertgl/croncache/lib"
	"github.com/ertgl/croncache/models"
	"github.com/ertgl/croncache/utils"
)

type Executer struct {
	lib.Executer
	moduleName string
	ioc        lib.IoC
}

func NewExecuter() lib.Executer {
	e := &Executer{
		moduleName: MODULE_NAME,
	}
	return e
}

func (e *Executer) ModuleName() string {
	return e.moduleName
}

func (e *Executer) Initialize() error {
	var err error = nil
	return err
}

func (e *Executer) IoC() lib.IoC {
	return e.ioc
}

func (e *Executer) SetIoC(ioc lib.IoC) {
	e.ioc = ioc
	if e.IoC().Executer() != e {
		e.IoC().SetExecuter(e)
	}
}

func (e *Executer) Execute(command string, timeout models.Duration, args ...string) (models.Cache, error) {
	cache := models.Cache{}
	command = utils.ReplaceOSVariables(command)
	for _, arg := range args {
		arg = utils.ReplaceOSVariables(arg)
	}
	cmd := exec.Command(command, args...)
	done := make(chan int, 1)
	e.IoC().Node().WaitGroup().Add(1)
	go func(cmd *exec.Cmd, timeout models.Duration, done chan int) {
		defer func() {
			e.IoC().Node().WaitGroup().Done()
		}()
	Waiting:
		for {
			select {
			case <-done:
				break Waiting
			case <-time.After(timeout.Duration):
				cmd.Process.Kill()
				break Waiting
			}
		}
	}(cmd, timeout, done)
	out, err := cmd.Output()
	select {
	case done <- 1:
	}
	if err != nil {
		return cache, err
	}
	buffer := bytes.NewBuffer(out)
	err = json.NewDecoder(buffer).Decode(&cache)
	if err != nil {
		return cache, err
	}
	return cache, nil
}
