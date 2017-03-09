package croncache

import (
	"github.com/ertgl/croncache/lib"
)

type Application struct {
	ioc lib.IoC
}

func NewApplication() *Application {
	app := &Application{}
	return app
}

func (app *Application) Initialize() error {
	var err error = nil
	return err
}

func (app *Application) IoC() lib.IoC {
	return app.ioc
}

func (app *Application) SetIoC(ioc lib.IoC) {
	app.ioc = ioc
}

func (app *Application) Run() error {
	err := app.IoC().Node().Start()
	return err
}
