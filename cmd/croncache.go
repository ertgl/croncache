package main

import (
	"flag"
	"io/ioutil"
)

import (
	"github.com/ertgl/croncache"
)

import (
	_ "github.com/ertgl/croncache/modules/cache_engine/radix/v1"
	_ "github.com/ertgl/croncache/modules/executer/croncache/v1"
	_ "github.com/ertgl/croncache/modules/ioc/croncache/v1"
	_ "github.com/ertgl/croncache/modules/node/croncache/v1"
	_ "github.com/ertgl/croncache/modules/task_manager/croncache/v1"
)

var (
	configFilePath *string = flag.String("config", "node.json", "Config file path.")
)

var (
	iocModuleName         string = "ioc/croncache/v1"
	executerModuleName    string = "executer/croncache/v1"
	taskManagerModuleName string = "task_manager/croncache/v1"
	nodeModuleName        string = "node/croncache/v1"
)

var (
	app *croncache.Application
)

func init() {
	flag.Parse()

	app = croncache.NewApplication()
	err := app.Initialize()
	if err != nil {
		croncache.HandleFatalError(err)
	}

	iocGenerator, err := croncache.IoCRepository().Resolve(iocModuleName)
	if err != nil {
		croncache.HandleFatalError(err)
	}
	app.SetIoC(iocGenerator())

	executerGenerator, err := croncache.ExecuterRepository().Resolve(executerModuleName)
	if err != nil {
		croncache.HandleFatalError(err)
	}
	app.IoC().SetExecuter(executerGenerator())
	err = app.IoC().Executer().Initialize()
	if err != nil {
		croncache.HandleFatalError(err)
	}

	taskManagerGenerator, err := croncache.TaskManagerRepository().Resolve(taskManagerModuleName)
	if err != nil {
		croncache.HandleFatalError(err)
	}
	app.IoC().SetTaskManager(taskManagerGenerator())
	err = app.IoC().TaskManager().Initialize()
	if err != nil {
		croncache.HandleFatalError(err)
	}

	nodeGenerator, err := croncache.NodeRepository().Resolve(nodeModuleName)
	if err != nil {
		croncache.HandleFatalError(err)
	}
	app.IoC().SetNode(nodeGenerator())
	nodeConfigRaw, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		croncache.HandleFatalError(err)
	}
	err = app.IoC().Node().ImportConfig(nodeConfigRaw)
	if err != nil {
		croncache.HandleFatalError(err)
	}
	err = app.IoC().Node().Initialize()
	if err != nil {
		croncache.HandleFatalError(err)
	}
}

func main() {
	err := app.Run()
	if err != nil {
		croncache.HandleFatalError(err)
	}
}
