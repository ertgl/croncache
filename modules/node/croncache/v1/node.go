package v1

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

import (
	"github.com/ertgl/croncache"
	"github.com/ertgl/croncache/lib"
	"github.com/ertgl/croncache/utils"
)

type Node struct {
	lib.Node
	moduleName string
	ioc        lib.IoC
	config     *Config
	waitGroup  *sync.WaitGroup
	isHalted   bool
	haltSignal chan bool
	logger     *log.Logger
}

func NewNode() lib.Node {
	n := &Node{
		moduleName: MODULE_NAME,
		config:     NewConfig(),
		waitGroup:  new(sync.WaitGroup),
		haltSignal: make(chan bool, 1),
	}
	return n
}

func (n *Node) ModuleName() string {
	return n.moduleName
}

func (n *Node) Initialize() error {
	n.config.LogFilePath = utils.ReplaceOSVariables(n.config.LogFilePath)
	f, err := os.OpenFile(n.config.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 755)
	if err != nil {
		return err
	}
	n.logger = log.New(f, "[Node] ", log.LstdFlags|log.Llongfile)
	n.logger.SetOutput(f)
	croncache.HandleFatalError = n.Log().Println
	err = n.IoC().TaskManager().ImportConfig(*n.config.TaskManager)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node) IoC() lib.IoC {
	return n.ioc
}

func (n *Node) SetIoC(ioc lib.IoC) {
	n.ioc = ioc
	if n.IoC().Node() != n {
		n.IoC().SetNode(n)
	}
}

func (n *Node) ImportConfig(raw []byte) error {
	err := json.Unmarshal(raw, &n.config)
	return err
}

func (n *Node) ExportConfig() ([]byte, error) {
	raw, err := json.MarshalIndent(&n.config, "", "\t")
	return raw, err
}

func (n *Node) IsHalted() bool {
	return n.isHalted
}

func (n *Node) SetIsHalted(isHalted bool) {
	n.isHalted = isHalted
}

func (n *Node) HaltSignal() chan bool {
	return n.haltSignal
}

func (n *Node) WaitGroup() *sync.WaitGroup {
	return n.waitGroup
}

func (n *Node) Start() error {
	var err error = nil
	n.WaitGroup().Add(1)
	go func() {
		defer n.WaitGroup().Done()
		n.IoC().TaskManager().Run()
	}()
	n.WaitGroup().Wait()
	return err
}

func (n *Node) Stop() error {
	var err error = nil
	n.SetIsHalted(true)
	n.HaltSignal() <- true
	return err
}

func (n *Node) Log() *log.Logger {
	return n.logger
}
