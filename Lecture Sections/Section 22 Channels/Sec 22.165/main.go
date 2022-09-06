package main

import (
	"fmt"
)

/*
Range

Ranging over a channel and closing a channel
*/

func main() {
	c := make(chan int)

	// send
	go func() {
		// send values onto a channel
		for i := 1; i <= 10; i++ {
			c <- i
		}
		// close the channel -  if we don't close here, will deadlock
		close(c)
	}()

	// receive
	// range will loop over channel until it is closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Succesfully sent and received from channel")
}
