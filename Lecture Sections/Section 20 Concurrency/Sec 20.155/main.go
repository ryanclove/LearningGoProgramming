package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
Atomic

Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

We can use package atomic to also prevent a race condition in our incrementer code.
atomic package - https://pkg.go.dev/sync/atomic

*/

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64 // anytime use int64, think package atomic

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // adds to counter address by 1
			runtime.Gosched()
			// this is safe, no race conditions with atomic load, read address from counter
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
