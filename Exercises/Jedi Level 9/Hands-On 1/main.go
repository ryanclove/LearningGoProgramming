package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
● in addition to the main goroutine, launch two additional goroutines
	○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines at start:\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()

	wg.Wait()

	fmt.Println("Goroutines before system exit:\t", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("Goroutines after foo:\t", runtime.NumGoroutine())
	wg.Done()
}

func bar() {
	fmt.Println("Goroutines after bar:\t", runtime.NumGoroutine())
	wg.Done()
}
