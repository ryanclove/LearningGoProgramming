package main

import (
	"fmt"
)

/*
Documentation

check README in this dir

*/

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
