package main

import (
	"context"
	"fmt"
)

/*
In Go servers, each incoming request is handled in its own goroutine.
Request handlers often start additional goroutines to access backends such as databases and RPC services.
The set of goroutines working on a request typically needs access to request-specific values such as
the identity of the end user, authorization tokens, and the request’s deadline. When a request is
canceled or times out, all the goroutines working on that request should exit quickly so the system
can reclaim any resources they are using.

At Google, we developed a context package that makes it easy to pass request-scoped values,
cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.
The package is publicly available as context

Context allows running and managing multiple goroutines within a specific process.
It provides execution, cancellation, shutdown and other operations for a set of goroutines in order to
prevent goroutine leaks.

Context is similar to Executor framework in java to manage tasks and threads.
*/
func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("---------")

	ctx, cancelFunc := context.WithCancel(ctx)
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel function:\t\t", cancelFunc)
	fmt.Printf("cancel function type:\t%T\n", cancelFunc)
	fmt.Println("---------")

	fmt.Println("Calling cancel func")
	cancelFunc()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel function:\t\t", cancelFunc)
	fmt.Printf("cancel function type:\t%T\n", cancelFunc)
	fmt.Println("---------")
}
