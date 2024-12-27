package gosafe

import (
	"sync"
)

func Go(f func() error) <-chan error {
	var wg sync.WaitGroup
	wg.Add(1)

	errChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		errChan <- safe(f)()
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	return errChan
}
