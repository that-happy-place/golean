package gosafe

import (
	"fmt"
	"sync"
)

func Go(f func() error) <-chan error {
	var wg sync.WaitGroup
	wg.Add(1)

	errChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: occurred in goroutine: %+v", r)
			}
		}()
		errChan <- f()
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	return errChan
}
