package gosafe

import (
	"fmt"
	"runtime"
)

func safe(f func() error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 1024)
				n := runtime.Stack(buf, false) // Capture the stack trace
				err = fmt.Errorf("panic: occurred in goroutine: %+v, Stack trace: %s", r, string(buf[:n]))
			}
		}()
		return f()
	}
}
