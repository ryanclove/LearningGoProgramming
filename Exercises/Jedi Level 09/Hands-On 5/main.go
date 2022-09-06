package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Fix the race condition you created in exercise #4
by using package atomic
*/

func main() {
	var wg sync.WaitGroup

	var incrementer int64
	gs := 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// atomic is nice, just call address and delta int
			atomic.AddInt64(&incrementer, 1)
			fmt.Println("Incrementer:\t", atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:\t", incrementer)
}
