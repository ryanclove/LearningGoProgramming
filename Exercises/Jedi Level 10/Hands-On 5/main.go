package main

import (
	"fmt"
)

/*
Show the comma ok idiom starting with this code:
	○ https://go.dev/play/p/0nHcd1z3IUu
*/

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
