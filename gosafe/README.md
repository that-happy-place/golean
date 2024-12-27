# gosafe

A safer way to trigger goroutines.

## How panic works in go?

Panics in go allows the application to fail fast quickly. In an idea world, we use `panic()` when we are initializing our application. This allows us to fail fast in case of any misconfigurations.

Panics can also be thrown during application runtime when there is a nil pointer or an index out of bound case.

To tackle such panics gracefully, go provides `recover()` function.

Traditionally, when we write a Go application with APIs/RPC/Consumer logic, we will write a wrapper or an interceptor to our APIs to gracefully handle and recover from such panics.

But, when it comes to async processing with goroutines, if a panic is thrown from within a goroutine, then the application will exit directly. Regardless of interceptors/wrappers that we have written.

## Why is this package needed?

In an ideal world, we should add necessary conditions before accessing an attribute when implementing something. And make sure such runtime panics never occur.

But in a fast developing environment, it is out of our control. We have seen these kinds of cases when incremental releases are done where an originally expected non-nil value is accidentally sent as nil by an upstream service. In these sencarios, the entire service goes down.

The motivation behind this package: While failing fast is good, it shouldn't necessarily lead us to a complete downtime.

## Installation

```shell
go get github.com/that-happy-place/golean/gosafe@latest
```

## Usage

### Triggering a single goroutine

```go
import (
    "github.com/that-happy-place/golean/gosafe"
)

func main() {
    errChan := gosafe.Go(func() error{
        // do stuff
        return nil  // return error if any
    })

    err := <-errChan
    if err != nil {
        // handle your error or panic here
    }
}
```

### Triggering a group of goroutines and waiting for them to end.

```go
import (
    "context"

    "github.com/that-happy-place/golean/gosafe"
)

func main() {
    group, ctx := gosafe.GoGroupWithContext(context.Background()) // use the returned context in your goroutines to propagate cancellations in case of single goroutine failing.

    group.Go(func() error{
        // do something
        return nil // return error if any
    })

    group.Go(func() error{
        // do something
        return nil // return error if any
    })

    // Wait() waits until all goroutines are done, or one of them returned an error. In case of errors, ctx returned will get cancelled.
    err := group.Wait()
    if err != nil {
        // handle your error or panic here
    }
}
```