package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
WaitGroup

A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of goroutines to
wait for. Then each of the goroutines runs and calls Done when
finished. At the same time, Wait can be used to block until all
goroutines have finished.

Add adds delta, which may be negative, to the WaitGroup counter.
If the counter becomes zero, all goroutines blocked on Wait are
released. If the counter goes negative, Add panics.

Done decrements the WaitGroup counter by one.

Wait blocks until the WaitGroup counter is zero.
*/

// has package scope
var wg sync.WaitGroup

func init() {
	// runs before main at the beginning of the program
}

// func main is the main Goroutine, so that's 1
func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	// func (wg *WaitGroup) Add(delta int)   - delta can be negative
	wg.Add(1) // add one thing to wait for, removes it when it's done
	// another Goroutine
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	// func (wg *WaitGroup) Done()
	// when foo is done, we declare it
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
