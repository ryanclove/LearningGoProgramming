package main

import (
	"fmt"
)

// Syntax

func main() {
	foo()
	bar("James") // Passing in the value James
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(paramaters) (return(s)) {...}

func foo() {
	fmt.Println("Hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

// Return
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

// Multiple returns
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ", says 'Hello'")
	b := true
	return a, b
}
