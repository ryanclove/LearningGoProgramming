package main

import (
	"fmt"
)

/*
Directional Channels

Read it from LEFT TO RIGHT
*/

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println("General channel buffer: ")
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("c:\t %T\n", c)

	// If we did a sim channel above,  chan <- int,2 and try to print <-c,
	// would return error: invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
	fmt.Println("----------------------------")

	fmt.Println("Send and Receive Channels: ")
	// bidrectional type
	c = make(chan int)

	// Read it from LEFT TO RIGHT
	cr := make(<-chan int) // receive from a channel of int
	cs := make(chan<- int) // send an int from a channel
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	fmt.Println("----------------------------")

	// specfific to general doesn't assign
	// c = cr  //error
	// c = cs  //error

	// specfific to specific doesn't assign
	// cs = cr // error

	// general to specific CAN assign
	fmt.Println("General to specific conversion:")
	cr = c
	cs = c
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}
