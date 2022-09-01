package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Race Condition

Race conditions are not good code. A race condition will give unpredictable results. We
will see how to fix this race condition in the next video.
*/

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 04
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// ==================================================================
			// can do either time.Sleep or runtime.Gosched:

			// I will choose runtime.GoSched so time.sleep will be commented out

			// make the goroutine sleep for a second
			//time.Sleep(time.Second)

			// Gosched yields the processor, allowing other goroutines to run.
			// It does not suspend the current goroutine, so execution resumes automatically.
			runtime.Gosched() // go ahead and run something else if you want to

			//==================================================================

			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter) // prints 104, should be 100
}
