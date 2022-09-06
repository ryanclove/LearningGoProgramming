package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
● Using goroutines, create an incrementer program
	○ have a variable to hold the incrementer value
	○ launch a bunch of goroutines
		■ each goroutine should
			● read the incrementer value
				○ store it in a new variable
			● yield the processor with runtime.Gosched()
			● increment the new variable
			● write the value in the new variable back to the incrementer
			variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag
● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej

in the terminal, run:
go run -race main.go

got this output in terminal after the print statements:
Found 1 data race(s)
exit status 66
*/

var counter = 0

var gs = 5

var wg sync.WaitGroup

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
			v := counter      // the race condition
			runtime.Gosched() // the race condition
			v++               // the race condition
			counter = v       // the race condition
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
			v := counter      // the race condition
			runtime.Gosched() // the race condition
			v++               // the race condition
			counter = v       // the race condition
			wg.Done()
		}()
		fmt.Printf("Counter value: %d\t 2nd incrementer Goroutines : %d\n", counter, runtime.NumGoroutine())
	}
}
