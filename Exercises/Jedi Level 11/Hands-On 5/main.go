package main

import (
	"fmt"
)

/*
We are going to learn about testing next. For this exercise, take a moment and see how much
you can figure out about testing by reading the testing documentation & also Caleb Doxsey’s
article on testing. See if you can get a basic example of testing working.

	● testing documentation - https://pkg.go.dev/testing?utm_source=godoc
	● Caleb Doxsey's article on testing - https://www.golang-book.com/books/intro/12
*/

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := Average(xs)
	fmt.Println(avg)
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
