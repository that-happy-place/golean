# gosafe

A safer way to trigger goroutines.

## How panic works in go?

Panics in go allows the application to fail fast quickly. In an idea world, we use `panic()` when we are initializing our application. This allows us to fail fast in case of any misconfigurations.

Panics can also be thrown during application runtime when there is a nil pointer or an index out of bound case.

To tackle such panics gracefully, go provides `recover()` function.

Traditionally, we will write a wrapper or an interceptor to our APIs to gracefully handle and recover from such panics.