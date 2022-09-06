package main

import "fmt"

/*
Use the “defer” keyword to show that a deferred func
runs after the func containing it exits
*/

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("1st function, deferred, executes and prints second")
}

func bar() {
	fmt.Println("2nd function, not deferred, executes and prints first")
}
