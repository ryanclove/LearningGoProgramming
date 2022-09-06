package main

import (
	"fmt"
	"sync"
)

/*
Fix the race condition you created in the previous exercise by using a mutex
	● it makes sense to remove runtime.Gosched()

	go run -race main.go
*/

// Todd's solution here:

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 10
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}

// ====================================================================================================

// This is the code I had in Hands-On 3. I tried making it mutex but could not figure out why
// it was outputting a race condition still.

/*
var counter = 0

var gs = 5

var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines start:", runtime.NumGoroutine())

	wg.Add(gs)
	incrementer1()
	wg.Wait()

	fmt.Println("Goroutines before system exit:", runtime.NumGoroutine())
	fmt.Println("End counter value:", counter)
}

func incrementer1() {
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Printf("Counter value: %d\t 1st incrementer Goroutines : %d\n", counter, runtime.NumGoroutine())
	}
	wg.Add(gs)
	incrementer2()
}

func incrementer2() {
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Printf("Counter value: %d\t 2nd incrementer Goroutines : %d\n", counter, runtime.NumGoroutine())
	}
}
*/
