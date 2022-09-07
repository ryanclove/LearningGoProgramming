package main

import (
	"fmt"
)

/*
write a program that
	○ launches 10 goroutines
		■ each goroutine adds 10 numbers to a channel
	○ pull the numbers off the channel and print them
*/

func main() {
	c := make(chan int)

	const goroutines = 10 // only launch 10 goroutines

	for j := 1; j <= goroutines; j++ {
		go func() {
			for i := 1; i <= 10; i++ {
				c <- i
			}
		}()
	}

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("About to exit")
}
