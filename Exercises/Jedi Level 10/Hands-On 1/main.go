package main

import (
	"fmt"
)

/*
get this code (https://go.dev/play/p/-DpZPo8o5JQ) working:
	○ using func literal, aka, anonymous self-executing func
	○ a buffered channel
*/

func main() {
	c := make(chan int)

	// using func literal
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	// using buffered channel
	c = make(chan int, 1)
	c <- 42
	fmt.Println(<-c)

}
