package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
Context
- https://go.dev/blog/context
- https://pkg.go.dev/context

So you don't leak Goroutines

Is a more advanced topic

In Go servers, each incoming request is handled in its own goroutine. Request handlers often
start additional goroutines to access backends such as databases and RPC services. The set of
goroutines working on a request typically needs access to request-specific values such as the
identity of the end user, authorization tokens, and the request's deadline. When a request is
canceled or times out, all the goroutines working on that request should exit quickly so the
system can reclaim any resources they are using. At Google, we developed a context package
that makes it easy to pass request-scoped values, cancellation signals, and deadlines across
API boundaries to all the goroutines involved in handling a request. The package is publicly
available as context. This article describes how to use the package and provides a complete
working example.
*/

func main() {
	// func Background() Context
	// Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests,
	// and as the top-level Context for incoming requests.
	ctx := context.Background()

	// =====================================================================================
	//                         Context Background Info and Cancels
	// =====================================================================================

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err()) // nil
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")

	// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	ctx, cancel := context.WithCancel(ctx)

	// before cancel
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err()) // nil
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")

	// after cancel
	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err()) // changes to canceled
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("=========================================================")

	// =====================================================================================
	//                           Example of Using Context
	// =====================================================================================
	fmt.Println("Example:")
	contextExample()

	fmt.Println("=========================================================")

	// =====================================================================================
	//                          Go package example of Using Context
	// =====================================================================================
	fmt.Println("Go's Select package Example:")
	packgeGoContextExample()

	fmt.Println("About to exit")
}

func contextExample() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done(): // if something is returned, returns a closed a channel
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num goroutines 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutines 3:", runtime.NumGoroutine())
}

func packgeGoContextExample() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}
