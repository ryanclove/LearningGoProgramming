package main

import "fmt"

/*
● Create a func which returns a func
● assign the returned func to a variable
● call the returned func
*/

func main() {
	f := foo()
	fmt.Println(f())
}

func foo() func() string {
	return func() string {
		return "You've returned a value of a func that returns a func()"
	}
}
