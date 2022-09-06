package main

import (
	"fmt"
)

/*
get this code running:
	○ https://play.golang.org/p/oB-p3KMiH6
	○ https://play.golang.org/p/_DBRueImEq

*/

func main() {

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("cs\t%T\n", cs)

	fmt.Printf("------\n")

	cr := make(chan int)
	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)
	fmt.Printf("cr\t%T\n", cr)

}
