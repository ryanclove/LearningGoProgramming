package main

import (
	"fmt"
)

/*
Starting with this code, https://go.dev/play/p/NSaEoBTT6c9
pull the values off the channel using a select statement
*/

func main() {

	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("From the channel:\t", v)
		case v := <-q:
			fmt.Println("From the quit channel:\t", v)
			return
		}
	}
}
