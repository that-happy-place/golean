package gosafe

import (
	"fmt"
)

func safe(f func() error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: occurred in goroutine: %+v", r)
			}
		}()
		return f()
	}
}
