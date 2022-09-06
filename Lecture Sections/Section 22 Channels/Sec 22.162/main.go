package main

import (
	"fmt"
)

/*
Understanding Channels

CHANNELS BLOCK

Send and receive must happen at the same time, or channel remains blocked

Reccomended to avoid buffer channels, but there's a time and place. Use unbuffer channels
*/

func main() {
	c := make(chan int)

	// ===========================================================================
	//                                Won't run
	// ===========================================================================

	// put something onto channel c
	// c <- 42

	// won't run: fmt.Println(<-c)
	// returns error: fatal error: all goroutines are asleep - deadlock!

	// unsuccessful Buffer (more put on than buffer allows)
	// e := unsuccesfulBuffer()
	// fmt.Println("Print by passing in unsuccesful buffer channel:", <-e)

	// ===========================================================================
	//                           2 ways to run this:
	// ===========================================================================

	// 1st way: channel launches off into 2nd goroutine
	successfulPassOfValue(c)
	// value blocks until it is ready to take it off
	fmt.Println("Print by passing in concurrent goroutine:", <-c)

	// 2nd way: channel buffer
	d := successfulBuffer()
	fmt.Println("Print by passing in buffer channel:", <-d)

	f := successfulMultipleBuffer()
	fmt.Println("Print by passing in multiple buffer channel (1st value put in channel):", <-f) // prints out 42
	fmt.Println("Print by passing in multiple buffer channel (2nd value put in channel):", <-f) // prints out 43

	g := successfulMultipleBuffer()
	fmt.Println("Print by passing in multiple buffer channel (both values put in channel):", <-g, <-g) // prints out both
}

func successfulPassOfValue(c chan int) {
	// code blocks into this goroutine
	go func() {
		// put something onto channel c
		c <- 42
	}()
}

func successfulBuffer() chan int {
	// buffer channel allows certain values to sit in that channel
	// regardless of if something is ready to take it off
	d := make(chan int, 1)
	d <- 42
	// 42 get put in the channel because it has a buffer of 1
	return d
}

func unsuccesfulBuffer() chan int {
	// unsuccesful buffer
	e := make(chan int, 1)
	e <- 42
	// we only have room for one (1), so deadlock error
	e <- 43
	return e
}

func successfulMultipleBuffer() chan int {
	f := make(chan int, 2)
	f <- 42
	f <- 43
	return f
}
