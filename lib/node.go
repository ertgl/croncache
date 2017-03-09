package lib

import (
	"log"
	"sync"
)

type Node interface {
	Dependency
	Configurable
	IsHalted() bool
	SetIsHalted(bool)
	HaltSignal() chan bool
	WaitGroup() *sync.WaitGroup
	Start() error
	Stop() error
	Log() *log.Logger
}
