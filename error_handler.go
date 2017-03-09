package croncache

import (
	"fmt"
)

var (
	HandleFatalError func(v ...interface{}) = DefaultFatalErrorHandler
)

func DefaultFatalErrorHandler(v ...interface{}) {
	template := ""
	for i := 0; i < len(v); i++ {
		template += "%+v "
	}
	panic(fmt.Sprintf(template, v...))
}
