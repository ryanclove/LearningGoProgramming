package main

import (
	"fmt"
)

/*
Starting with this code, https://go.dev/play/p/WHySMDT9iOP
pull the values off the channel using a for range loop
*/

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
