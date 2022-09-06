package main

import (
	"fmt"
)

/*
Comma ok idiom
*/

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool) // changed to bool

	go send(even, odd, quit)
	receive(even, odd, quit)

	// use quit as an int channel
	quitInt := make(chan int)
	go sendWithQuitInt(even, odd, quitInt)
	receiveWithQuitInt(even, odd, quitInt)

	// check when sending a value over
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Printf("Check when sending a value %v over\t ok: %v\n", v, ok) // 42 true
	v, ok = <-c
	fmt.Printf("After C closed, check sending a value %v over\t ok: %v\n", v, ok) // 0 false

	fmt.Println("About to exit")
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("Value from the even channel:\t", v)
		case v := <-o:
			fmt.Println("Value from the odd channel:\t", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok:\t", i, ok) // false false
				return
			} else {
				fmt.Println("From comma ok:\t", i)
				return
			}
		}
	}
}

func sendWithQuitInt(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receiveWithQuitInt(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Value from the even channel:\t", v)
		case v := <-o:
			fmt.Println("Value from the odd channel:\t", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok:\t", i, ok) // 0 false
				return
			} else {
				fmt.Println("From comma ok:\t", i)
				return
			}
		}
	}
}
