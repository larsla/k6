package js

import (
	"time"

	"github.com/dop251/goja"
)

func registerGlobals(rt *goja.Runtime) {
	rt.Set("setTimeout", setTimeout)
}

func setTimeout(callback goja.Value, ms goja.Value) error {
	time.Sleep(time.Duration(ms.ToFloat() * float64(time.Millisecond)))
	fn, ok := goja.AssertFunction(callback)
	if ok {
		fn(goja.Undefined())
	}

	return nil
}
