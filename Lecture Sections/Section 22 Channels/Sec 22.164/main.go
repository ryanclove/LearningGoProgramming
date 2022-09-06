package main

import (
	"fmt"
)

/*
Using Channels

go run -race main.go
doesn't return a race
*/

func main() {
	// create channel
	c := make(chan int)

	// send
	go foo(c) // another goroutine

	// receive
	bar(c) // blocking until a value is sent

	fmt.Println("Succesfully sent and received from channel")
}

// send func
// making general c to specific send (cs)
func foo(c chan<- int) {
	// scope of this c is only for this function
	c <- 42
	// pass 42 onto c
}

// receive func
// making general c to specific receive (cr)
func bar(c <-chan int) {
	fmt.Println(<-c)
}
