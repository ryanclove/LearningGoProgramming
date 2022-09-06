package main

import (
	"fmt"
)

/*
Select

Select statements pull the value from whatever channel has a value ready to be pulled

go run -race main.go
no race condition
*/

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("About to exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i // put onto the even channel the value of i
		} else {
			o <- i // put onto the odd channel the value of i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel:\t", v)
		case v := <-o:
			fmt.Println("From the odd channel:\t", v)
		case v := <-q:
			fmt.Println("From the quit channel:\t", v)
			return
		}
	}
}
