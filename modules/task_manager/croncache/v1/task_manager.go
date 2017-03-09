package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

import (
	"github.com/ertgl/croncache"
	"github.com/ertgl/croncache/lib"
	"github.com/ertgl/croncache/models"
)

var (
	TaskTimeoutError error = errors.New("Task is timed out")
)

type TaskManager struct {
	lib.TaskManager
	moduleName string
	ioc        lib.IoC
	config     *Config
}

func NewTaskManager() lib.TaskManager {
	tm := &TaskManager{
		moduleName: MODULE_NAME,
		config:     NewConfig(),
	}
	return tm
}

func (tm *TaskManager) ModuleName() string {
	return tm.moduleName
}

func (tm *TaskManager) Initialize() error {
	var err error = nil
	return err
}

func (tm *TaskManager) IoC() lib.IoC {
	return tm.ioc
}

func (tm *TaskManager) SetIoC(ioc lib.IoC) {
	tm.ioc = ioc
	if tm.IoC().TaskManager() != tm {
		tm.IoC().SetTaskManager(tm)
	}
}

func (tm *TaskManager) ImportConfig(raw []byte) error {
	err := json.Unmarshal(raw, &tm.config)
	return err
}

func (tm *TaskManager) ExportConfig() ([]byte, error) {
	raw, err := json.MarshalIndent(&tm.config, "", "\t")
	return raw, err
}

func (tm *TaskManager) runTask(task *models.Task) error {
	ch := make(chan models.Cache, 1)
	errChan := make(chan error, 1)
	go func(ch chan models.Cache, errChan chan error) {
		cache, err := tm.IoC().Executer().Execute(
			task.Command,
			task.Args...,
		)
		if err != nil {
			errChan <- err
			return
		}
		ch <- cache
	}(ch, errChan)
	select {
	case cache := <-ch:
		cacheEngineGenerator, err := croncache.CacheEngineRepository().Resolve(
			task.CacheEngineModuleName,
		)
		if err != nil {
			return err
		}
		cacheEngine := cacheEngineGenerator()
		err = cacheEngine.ImportConfig(*task.CacheEngineCredentials)
		if err != nil {
			return err
		}
		err = cacheEngine.Initialize()
		if err != nil {
			return err
		}
		err = cacheEngine.Upsert(cache)
		if err != nil {
			return err
		}
	case <-time.After(task.Timeout.Duration):
		return TaskTimeoutError
	case <-tm.IoC().Node().HaltSignal():
		return nil
	}
	return nil
}

func (tm *TaskManager) Process(task *models.Task) {
	for !tm.IoC().Node().IsHalted() {
		tm.IoC().Node().WaitGroup().Add(1)
		go func(task *models.Task) {
			defer tm.IoC().Node().WaitGroup().Done()
			iter := 0
			for !tm.IoC().Node().IsHalted() {
				err := tm.runTask(task)
				if err != nil {
					task.Log().Fatalln(err)
					iter++
					if iter <= task.IterationOnFail {
						task.Log().Println(fmt.Sprintf(
							"Failover is enabled for %d times",
							iter,
						))
						continue
					}
				}
				break
			}
		}(task)
		time.Sleep(task.Interval.Duration)
	}
}

func (tm *TaskManager) Run() {
	for _, taskFilePath := range tm.config.Tasks {
		select {
		case <-tm.IoC().Node().HaltSignal():
			return
		default:
			var task *models.Task
			raw, err := ioutil.ReadFile(taskFilePath)
			if err != nil {
				tm.IoC().Node().Log().Fatalln(taskFilePath, err)
				continue
			}
			err = json.Unmarshal(raw, &task)
			if err != nil {
				tm.IoC().Node().Log().Fatalln(taskFilePath, err)
				continue
			}
			err = task.Initialize()
			if err != nil {
				tm.IoC().Node().Log().Fatalln(taskFilePath, err)
				continue
			}
			tm.IoC().Node().WaitGroup().Add(1)
			go func(task *models.Task) {
				defer tm.IoC().Node().WaitGroup().Done()
				tm.Process(task)
			}(task)
		}
	}
}
