package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Fan in

Taking values from many channels, and putting those values onto one channel.

*/

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)           // taken values from 2 channels
	go receive(even, odd, fanin) // send em into 1

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Rob Pike code: ")
	robPike()

	fmt.Println("About to exit")
}

func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

// take vlaues from 2 channels and fan in to 1
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

// ============================================================================
//                               Rob Pike Code
// ============================================================================

func robPike() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
