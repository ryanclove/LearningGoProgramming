package main

import (
	"errors"
	"fmt"
	"log"
)

/*
Errors with info

We can add information to our errors. We can do this with
	● errors.New()
		○ fmt.Errorf()
	● builtin.error
“Error values in Go aren’t special, they are just values like any other,
and so you have the entire language at your disposal.” - Rob Pike

*/

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
