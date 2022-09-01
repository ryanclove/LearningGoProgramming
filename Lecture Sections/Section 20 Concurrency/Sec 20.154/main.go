package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Mutex

A “mutex” is a mutual exclusion lock. Mutexes allow us to lock our code
so that only one goroutine can access that locked chunk of code at a time.

Mutex in sync package - Effective Go - https://pkg.go.dev/sync#Mutex

RWMutex - Effective Go - https://pkg.go.dev/sync#RWMutex

RW means Reading/Writing, so here you can lock things for reading but when someone's going
to write to it, lock it for writing. Allows unlimited reads at the same time,
since nothing is changing
*/

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// mutual exclusion lock
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// code gets locked down, no one else can access the counter variable
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			// code unlocks so counter variable is available elsewhere
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter) // prints 100
}
