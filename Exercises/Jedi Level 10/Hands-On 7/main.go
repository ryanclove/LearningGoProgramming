package main

import (
	"fmt"
	"runtime"
)

/*
Create a program that prints out your OS and ARCH. Use the following commands to run it
	● go run
	● go build
	● go install
*/

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)

	// didn't ask for this but for fun:
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}
