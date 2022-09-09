package main

import (
	"fmt"
)

/*
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
has a value of type error as a parameter. Create a value of type “customErr” and pass it into
“foo”. If you need a hint, here is one - https://go.dev/play/p/L5QWV8-p11
*/

type customErr struct {
	info string
}

// any value of type customErr will implicitely implement the Error() interface
func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func main() {
	c := customErr{
		info: "You aren't good at programming",
	}
	foo(c)
}

func foo(e error) {
	fmt.Println("Foo ran -", e)
	// use insertion (not conversion) to implement assertion of interface
	fmt.Println("Foo ran and used insertion -", e.(customErr).info)
}

/* conversion would be referencing a type to a var
type hotdog int
var x hotdog = 42
var y int
var y = int(x) // convert type hotdog to int
*/
