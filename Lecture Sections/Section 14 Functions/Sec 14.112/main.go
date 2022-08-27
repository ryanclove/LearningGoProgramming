package main

import "fmt"

// Defer

/*
A "defer" statement invokes a function whose execution is deferred to the moment
the surrounding function returns, either because the surrounding function executed a
return statement, reached the end of its function body, or because the corresponding
goroutine is panicking.
*/

func main() {
	// the defer causes bar to run first, not foo
	defer foo()
	bar()
	// after boo executes, calls foo

	// usually helpful when reading a file then later writing a file
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
